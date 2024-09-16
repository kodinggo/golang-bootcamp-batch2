package main

import "fmt"

func main() {
	defer fmt.Println("hello 1")
	defer fmt.Println("hello 2")
	defer fmt.Println("hello 3")

	fmt.Println("hi")
	fmt.Println("whatsapp")
	fmt.Println("halo")
	fmt.Println("apa kabar")
}
