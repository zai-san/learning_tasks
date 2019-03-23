package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
)

func GenerateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func GenerateRandomString(s int) (string, error) {
	b, err := GenerateRandomBytes(s)
	return base64.URLEncoding.EncodeToString(b), err
}

func main() {
	//	token, err := GenerateRandomString(32)
	//	var t int
	//	fmt.Println("Enter number: ")
	//	fmt.Scan(&t)
	//	token, err := GenerateRandomBytes(t)
	token, err := GenerateRandomBytes(8)
	if err != nil {
	}
	//	i, err := strconv.ParseInt(token, 10, 64)
	//	if err != nil {
	//		return
	//	}
	//	fmt.Println(i)
	//	data := binary.BigEndian.Uint64(token)
	//	fmt.Println(data)
	fmt.Println(token[1])
}
