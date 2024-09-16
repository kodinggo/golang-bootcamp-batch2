package main

import "fmt"

func main() {
	fmt.Println(sayHello("Agus"))

	fmt.Println(sum(1, 2, 3))

	sum, diff := cal(5, 2)
	fmt.Println(sum, diff)
}

func sayHello(name string) string {
	return "Hi " + name
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
