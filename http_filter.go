package main

import (
	"context"
	"log"
	"time"

	"github.com/envoyproxy/envoy/contrib/golang/common/go/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "alicloud-clickhouse-autopause-proxy/clickhouse"
)

// 最后一次请求时间戳(秒)
var lastRequestTime time.Time = time.Now()

type AutoPauseStreamFilter struct {
	api.PassThroughStreamFilter

	callbacks api.FilterCallbackHandler
	// The configuration for this filter.
	config *HttpConfig
}

func (f *AutoPauseStreamFilter) ResponseError(statusCode int) api.StatusType {
	f.callbacks.SendLocalReply(statusCode, "", nil, 0, "")
	return api.LocalReply
}

func (f *AutoPauseStreamFilter) DecodeHeaders(header api.RequestHeaderMap, endStream bool) api.StatusType {
	// 如果距离上次请求没超过5秒，直接返回
	if time.Since(lastRequestTime).Seconds() < 5 {
		return api.Continue
	}

	log.Default().Println("Request KeepAlive: RegionID: ", f.config.RegionID, ", DBInstanceID: ", f.config.InstanceID)

	conn, err := grpc.Dial(f.config.AutopauseService, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Default().Println(err)
		return api.Continue
	}

	defer conn.Close()

	client := pb.NewAliYunClickhouseClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)

	defer cancel()

	r, err := client.KeepAlive(ctx, &pb.KeepAliveRequest{RegionID: f.config.RegionID, DBInstanceID: f.config.InstanceID})

	if err != nil {
		log.Default().Println(err)
		return f.ResponseError(499)
	}

	if !r.Success {
		return f.ResponseError(503)
	}

	lastRequestTime = time.Now()

	return api.Continue
}

func main() {
}
