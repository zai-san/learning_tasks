package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {

	file, err := os.OpenFile("4th_task.txt", os.O_WRONLY, 0666)
	if err != nil {
		return
	}
	defer file.Close()
	i := 1
	var x, y, z int
	rand.Seed(time.Now().UTC().UnixNano())
	for i <= 10 {
		x = rand.Intn(100)
		y = rand.Intn(100)
		if x == y {
			z = x
		} else if x != z {
			str := strconv.Itoa(x)
			file.WriteString(str)
			file.WriteString(" \r\n")
			i++
		} else if y != z {
			str := strconv.Itoa(y)
			file.WriteString(str)
			file.WriteString(" \r\n")
			i++
		} else {
			str := strconv.Itoa(x)
			file.WriteString(str)
			file.WriteString(" \r\n")
			i++
		}

	}
	fmt.Println("Write in file DONE")
}
