package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
)

func main() {
	bs, err := ioutil.ReadFile("3rd_task.txt")
	if err != nil {
		return
	}
	str := string(bs)

	i, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return
	}

	fmt.Println(i + 2)
}
