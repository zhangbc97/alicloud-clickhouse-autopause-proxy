package main

import (
	"log"
	"net"
	"sync"
	"time"

	xds "github.com/cncf/xds/go/xds/type/v3"
	"google.golang.org/protobuf/types/known/anypb"

	"github.com/envoyproxy/envoy/contrib/golang/common/go/api"
	"github.com/envoyproxy/envoy/contrib/golang/filters/network/source/go/pkg/network"
)

func init() {
	network.RegisterNetworkFilterConfigFactory("autopause_tcp", cf)
}

var globalLock = sync.Mutex{}

var lastKeepAliveTime int64 = 0

var cf = &configFactory{}

type configFactory struct{}

func (f *configFactory) CreateFactoryFromConfig(config interface{}) network.FilterFactory {
	a := config.(*anypb.Any)
	configStruct := &xds.TypedStruct{}
	_ = a.UnmarshalTo(configStruct)

	serverAddr := configStruct.Value.AsMap()["clickhouse_server_addr"]
	serverPort := configStruct.Value.AsMap()["clickhouse_server_port"]

	addr, err := net.LookupHost(serverAddr.(string))
	if err != nil {
		log.Default().Printf("fail to resolve: %v, err: %v\n", serverAddr.(string), err)
		return nil
	}
	upAddr := addr[0] + ":" + serverPort.(string)

	return &filterFactory{
		upAddr: upAddr,
	}
}

type filterFactory struct {
	upAddr string
}

func (f *filterFactory) CreateFilter(cb api.ConnectionCallback) api.DownstreamFilter {
	return &downFilter{
		upAddr: f.upAddr,
		cb:     cb,
	}
}

type downFilter struct {
	api.EmptyDownstreamFilter

	cb       api.ConnectionCallback
	upAddr   string
	upFilter *upFilter
}

// 调用的时候判断，如果间隔小于5秒那就不调用，否则调用Clickhouse的Keepalive，注意并发加锁
func KeepAlive() {
	globalLock.Lock()
	defer globalLock.Unlock()
	if time.Now().Unix()-lastKeepAliveTime < 5 {
		return
	}
	lastKeepAliveTime = time.Now().Unix()
	// 调用Clickhouse的Keepalive

}

func (f *downFilter) OnNewConnection() api.FilterStatus {

	f.upFilter = &upFilter{
		downFilter: f,
		ch:         make(chan []byte, 1),
	}

	network.CreateUpstreamConn(f.upAddr, f.upFilter)
	return api.NetworkFilterContinue
}

func (f *downFilter) OnData(buffer []byte, endOfStream bool) api.FilterStatus {
	f.upFilter.ch <- buffer
	return api.NetworkFilterContinue
}

func (f *downFilter) OnEvent(event api.ConnectionEvent) {
}

func (f *downFilter) OnWrite(buffer []byte, endOfStream bool) api.FilterStatus {
	return api.NetworkFilterContinue
}

type upFilter struct {
	api.EmptyUpstreamFilter

	cb         api.ConnectionCallback
	downFilter *downFilter
	ch         chan []byte
}

func (f *upFilter) OnPoolReady(cb api.ConnectionCallback) {
	f.cb = cb
	go func() {
		for {
			buf, ok := <-f.ch
			if !ok {
				return
			}
			f.cb.Write(buf, false)
		}
	}()
}

func (f *upFilter) OnPoolFailure(poolFailureReason api.PoolFailureReason, transportFailureReason string) {
}

func (f *upFilter) OnData(buffer []byte, endOfStream bool) {
	f.downFilter.cb.Write(buffer, endOfStream)
}

func (f *upFilter) OnEvent(event api.ConnectionEvent) {
	if event == api.LocalClose || event == api.RemoteClose {
		close(f.ch)
	}
}

func main() {}
