package main

import (
	"errors"
	"fmt"
)

func main() {
	if n := 1; n == 1 {
		fmt.Println("n is 1")
	}

	hari := 3
	switch hari {
	case 1, 2, 3, 4, 5:
		fmt.Println("yes")
	}

	day := 2
	switch day {
	case 1:
		fmt.Println("sunday")
	case 2:
		fmt.Println("monday")
		fallthrough
	case 3:
		fmt.Println("tuesday")
	}

	fruits := []string{"manggo", "orange", "banana"}
	for index := range fruits {
		fmt.Println(index, fruits[index])
	}

	// Forever loop
	i := 0
	for {
		if i >= 20 {
			break
		}
		fmt.Println(i)
		i++
	}

	if err := save(); err != nil {
		panic(err)
	}
}

func save() error {
	return errors.New("error")
}
