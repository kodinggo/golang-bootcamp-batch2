package main

import "fmt"

func main() {
	defer func() {
		fmt.Println(recover())
	}()

	panic("contoh")
}
