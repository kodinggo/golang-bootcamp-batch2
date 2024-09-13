package main

import "fmt"

func main() {
	fmt.Println("Hello")
	fmt.Println("World")

	fmt.Print("hello\n")
	fmt.Print("world\n")

	// Cetak outout dg formatting
	fmt.Printf("%.0f\n", 10.0)
	fmt.Printf("My name is %s, I'm %d years old\n", "John", 20)

	// Mendapatkan input dari keyboard
	var input string
	fmt.Scan(&input)
	fmt.Printf("Welcome %s", input)

	// Formatting string (tidak mencatak)
	grade := fmt.Sprintf("Your grade is %.0f\n", 10.0)
	fmt.Println(grade)
}
