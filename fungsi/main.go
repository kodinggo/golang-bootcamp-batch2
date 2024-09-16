package main

import "fmt"

func main() {
	fmt.Println(sayHello("Agus"))

	fmt.Println(sum(1, 2, 3))

	sum, diff := cal(5, 2)
	fmt.Println(sum, diff)

	greeting := func() {
		fmt.Println("sum", sum)
		fmt.Println("hello")
	}
	greeting()

	greet("Joko", func(name string) {
		fmt.Println(name) // Joko
	})
}

func sayHello(name string) (s string) {
	s = "Hi " + name
	return
}

func sum(n ...int) int {
	var total int

	for _, val := range n {
		total += val
	}

	return total
}

func cal(a, b int) (int, int) {
	return a + b, a - b
}

func greet(name string, callback func(string)) {
	callback(name)
}
