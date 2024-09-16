package main

import "fmt"

func main() {
	var msg any

	msg = "hello"
	strMsg, ok := msg.(string) // mendapatkan nilai dengan type data yg asli
	if ok {
		cetak(strMsg)
	}

	msg = 123
	fmt.Println(msg)

	msg = []int{1, 2}
	sliceMsg, ok := msg.([]int) // mendapatkan nilai dengan type data yg asli
	if ok {
		fmt.Println(sliceMsg)
	}
}

func cetak(s string) {
	fmt.Println(s)
}
