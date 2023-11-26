FROM envoyproxy/envoy:contrib-v1.28.0 AS build-network

RUN sed -i s@/archive.ubuntu.com/@/mirrors.aliyun.com/@g /etc/apt/sources.list && \
    sed -i s@/security.ubuntu.com/@/mirrors.aliyun.com/@g /etc/apt/sources.list && \
    apt-get clean && apt update && apt install wget gcc g++ make -y

RUN wget https://go.dev/dl/go1.21.4.linux-amd64.tar.gz && \
    tar -C /usr/local -xzf go1.21.4.linux-amd64.tar.gz && \
    rm go1.21.4.linux-amd64.tar.gz && \
    ln -s /usr/local/go/bin/go /usr/local/bin/go

WORKDIR /go/src/app

COPY . .

# RUN go env -w GOPROXY=https://goproxy.cn,direct

RUN go build -buildmode=c-shared -v -o autopause_tcp.so

FROM envoyproxy/envoy:contrib-v1.28.0

COPY --from=build-network /go/src/app/autopause_tcp.so /lib/autopause_tcp.so

COPY envoy.yaml /etc/envoy/envoy.yaml

ENV GODEBUG cgocheck=0