package main

import (
	"fmt"
	"time"
)

func main() {
	go cetak("Hai")
	go cetak("bye")
	go cetak("Halo")

	cetak("Apa kabar?")

	nums := []int{20, 50, 30, 40}
	for _, num := range nums {
		go getSequare(num)
	}

	time.Sleep(1 * time.Millisecond)
}

func cetak(s string) {
	fmt.Println(s)
}

func getSequare(n int) {
	sq := n * n
	fmt.Println(sq)
}
