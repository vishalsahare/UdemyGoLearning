package main

import "fmt"

type person struct {
	first string
	age   int
}

type human interface {
	speak()
}

func (p *person) speak() {
	fmt.Printf("Person %s having age %d speaking...\n",
		p.first, p.age)
}

func saySomething(h human) {
	h.speak()
}

func main() {
	p1 := person{
		first: "Vishal",
		age:   40,
	}

	saySomething(&p1)
	p1.speak()
}
