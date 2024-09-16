package main

import "fmt"

type User interface {
	CetakName()
	SetName(name string)
	X()
}

type Person struct {
	Name string
	Age  int
}

// Method milik si Person
func (p Person) CetakName() {
	fmt.Println(p.Name)
}

func (p *Person) SetName(name string) {
	p.Name = name
}

func (p Person) X() {

}

func main() {
	// Proses intance
	var person User = &Person{
		Name: "Bambang",
		Age:  23,
	}

	person.CetakName()

	person.SetName("Joko")

	person.CetakName()
}
