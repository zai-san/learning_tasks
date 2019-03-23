package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var mas [10]int

	i := 0
	var x, y, z int
	rand.Seed(time.Now().UTC().UnixNano())
	for i < 10 {
		x = rand.Intn(100)
		y = rand.Intn(100)
		if x == y {
			z = x
		} else if x != z {
			mas[i] = x
			i++
		} else if y != z {
			mas[i] = y
			i++
		} else {
			mas[i] = x
			i++
		}
	}
	jsonmas, err := json.Marshal(mas)
	if err != nil {
		return
	}
	fmt.Println(string(jsonmas))
}
