package main

import "fmt"

var defaultNumberPerPage = 10

func main() {
	// Variable dg keyword "var"
	var angka int = 10
	fmt.Println(angka) // 10

	var name string = "John"
	fmt.Println(name) // name

	// Variable "var" tanpa tipe data
	var number = 10
	fmt.Println(number) // 10

	// Variable tanpa value
	var email string
	fmt.Println(email)

	var age int
	fmt.Println(age)

	// Varible shorthand
	city := "Jakarta"
	isActive := true
	gender := int64(1)
	fmt.Println(city)
	fmt.Println(isActive)
	fmt.Println(gender)

	// Deklarasi multiple variable dg "var"
	// var (
	// 	abc = 1
	// 	efg = 2
	// )

}

func sum() int {
	// total := 0
	var total int
	for i := 0; i <= 100; i++ {
		total += i
	}
	return total
}
