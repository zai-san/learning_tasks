package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/ghodss/yaml"
)

type Array []struct {
	NullValue interface{} `json:"null_value,omitempty"`
	Boolean   bool        `json:"boolean,omitempty"`
	Integer   int         `json:"integer,omitempty"`
}

type Object struct {
	Key   string `json:"key"`
	Array `json:"array"`
}

type HTTPReferrer struct {
	JSON      []string        `json:"json"`
	Yaml      []string        `json:"yaml"`
	Object    json.RawMessage `json:"object"`
	Paragraph string          `json:"paragraph"`
	Content   string          `json:"content"`
}

type JFile struct {
	RemoteAddr          string          `json:"remote_addr"`
	BodyBytesSent       int             `json:"body_bytes_sent"`
	RequestTime         float64         `json:"request_time"`
	ResponseStatus      int             `json:"response_status"`
	Request             string          `json:"request"`
	RequestMethod       string          `json:"request_method"`
	Host                string          `json:"host"`
	UpstreamCacheStatus string          `json:"upstream_cache_status"`
	UpstreamAddr        string          `json:"upstream_addr"`
	HTTPXForwardedFor   string          `json:"http_x_forwarded_for"`
	HTTPReferrer        json.RawMessage `json:"http_referrer"`
	HTTPUserAgent       string          `json:"http_user_agent"`
	ClusterName         string          `json:"cluster_name"`
	Type                string          `json:"type"`
}

func main() {
	file, err := ioutil.ReadFile("7th_task.json")
	if err != nil {
		fmt.Println("Failed read file: %s\n", err)
	}
	ydat, err := yaml.JSONToYAML(file)
	if err != nil {
		fmt.Println("Failed convert file: %s\n", err)
	}
	ioutil.WriteFile("8th_task.yaml", ydat, 0777)
	fmt.Println("File writed!")
}
