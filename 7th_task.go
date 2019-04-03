package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"reflect"
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
	var dat = &JFile{}
	if err := json.Unmarshal(file, dat); err != nil {
		fmt.Println("Failed parse file1: %s\n", err)
	}

	tx := reflect.ValueOf(dat).Elem()
	dattx := tx.Type()

	var dat1 = &HTTPReferrer{}
	if err := json.Unmarshal(dat.HTTPReferrer, dat1); err != nil {
		fmt.Println("Failed parse file2: %s\n", err)
	}

	tx1 := reflect.ValueOf(dat1).Elem()
	dattx1 := tx1.Type()

	var dat2 = &Object{}
	if err := json.Unmarshal(dat1.Object, dat2); err != nil {
		fmt.Println("Failed parse file3: %s\n", err)
	}

	tx2 := reflect.ValueOf(dat2).Elem()
	dattx2 := tx2.Type()

	for i := 0; i < tx.NumField(); i++ {
		f := tx.Field(i)
		if dattx.Field(i).Name == "HTTPReferrer" {
			fmt.Printf("%d: %s : \n", i, dattx.Field(i).Name)
			for i1 := 0; i1 < tx1.NumField(); i1++ {
				f1 := tx1.Field(i1)
				if dattx1.Field(i1).Name == "Object" {
					fmt.Printf("%d.%d: %s : \n", i, i1, dattx1.Field(i1).Name)
					for i2 := 0; i2 < tx2.NumField(); i2++ {
						f2 := tx2.Field(i2)
						fmt.Printf("%d.%d.%d: %s = %v\n", i, i1, i2, dattx2.Field(i2).Name, f2.Interface())
					}
				} else {
					fmt.Printf("%d.%d: %s = %v\n", i, i1, dattx1.Field(i1).Name, f1.Interface())
				}
			}
		} else {
			fmt.Printf("%d: %s = %v\n", i, dattx.Field(i).Name, f.Interface())
		}
	}
}
