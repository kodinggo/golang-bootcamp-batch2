package main

import "fmt"

func main() {
	var fruits [3]string
	fruits = [3]string{"banana", "mango", "melon"}
	fmt.Println(fruits)
	fmt.Println(fruits[1]) // mango

	var numbers []int
	numbers = []int{1, 2, 3}
	fmt.Println(numbers)
	numbers = append(numbers, 4)
	fmt.Println(numbers)
	numbers = append([]int{5}, numbers...) // prepend
	fmt.Println(numbers)

	// Operasi slicing
	angka := []int{1, 2, 3, 4, 5}
	subAngka := angka[2:4]
	fmt.Println(subAngka)

	angka = append(angka[0:2], angka[3:]...)
	fmt.Println(angka)
}
