package main

import (
	"encoding/json"
	"fmt"
)

type myStruct struct{ Name string }

func main() {
	var dat myStruct
	json.Unmarshal([]byte(`{"type":"fruit","name":"apple","color":"red"}`), &dat)
	fmt.Print(dat.Name, "\n")
	i := 1
	switch i {
	case 2:
		fmt.Println("inside switch")
	}
	fmt.Println("outside switch")
}
