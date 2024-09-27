package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	resp, _ := http.Get("https://dummyjson.com/products")
	byteBody, _ := io.ReadAll(resp.Body)
	fmt.Println(string(byteBody))
}
