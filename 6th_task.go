package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
)

func main() {
	bs, err := ioutil.ReadFile("6th_task.txt")
	if err != nil {
		return
	}
	str := base64.URLEncoding.EncodeToString(bs)
	fmt.Println(str)
}
