package main

import (
	"encoding/json"
	"fmt"

	xds "github.com/cncf/xds/go/xds/type/v3"
	"google.golang.org/protobuf/types/known/anypb"

	"github.com/envoyproxy/envoy/contrib/golang/common/go/api"
	"github.com/envoyproxy/envoy/contrib/golang/filters/http/source/go/pkg/http"
)

const Name = "alicloud-clickhouse-autopause-proxy-http"

func init() {
	http.RegisterHttpFilterConfigFactoryAndParser(Name, ConfigFactory, &Parser{})
}

type HttpConfig struct {
	AutopauseService string `json:"autopause_service"`
	RegionID         string `json:"region_id"`
	InstanceID       string `json:"instance_id"`
}

type Parser struct {
}

// Parse the filter configuration. We can call the ConfigCallbackHandler to control the filter's
// behavior
func (p *Parser) Parse(any *anypb.Any, callbacks api.ConfigCallbackHandler) (interface{}, error) {
	configStruct := &xds.TypedStruct{}
	if err := any.UnmarshalTo(configStruct); err != nil {
		return nil, err
	}
	v := configStruct.Value
	conf := &HttpConfig{}

	jsonbody, err := json.Marshal(v.AsMap())

	if err != nil {
		fmt.Println("error marshal config")
		return nil, err
	}

	if err := json.Unmarshal(jsonbody, conf); err != nil {
		fmt.Println("error unmarshal config")
		return nil, err
	}

	return conf, nil
}

// Merge configuration from the inherited parent configuration
func (p *Parser) Merge(parent interface{}, child interface{}) interface{} {
	parentConfig := parent.(*HttpConfig)
	childConfig := child.(*HttpConfig)

	newConfig := *parentConfig
	if childConfig.AutopauseService != "" {
		newConfig.AutopauseService = childConfig.AutopauseService
	}

	if childConfig.InstanceID != "" {
		newConfig.InstanceID = childConfig.InstanceID
	}

	return &(*parent.(*HttpConfig))
}

func ConfigFactory(c interface{}) api.StreamFilterFactory {
	conf, ok := c.(*HttpConfig)
	if !ok {
		panic("unexpected config type")
	}

	return func(callbacks api.FilterCallbackHandler) api.StreamFilter {
		return &AutoPauseStreamFilter{
			callbacks: callbacks,
			config:    conf,
		}
	}
}
