package main

import "fmt"

type Person struct {
	name string
	age  int
}

func (p *Person) SetName(name string) {
	p.name = name
}

func (p Person) SetAge(age int) {
	p.age = age
	fmt.Println(p.age)
}

func main() {
	person := Person{}
	person.SetName("Miguel")
	person.SetAge(3)
	fmt.Printf("Name: %s - Age: %d", person.name, person.age)
}
