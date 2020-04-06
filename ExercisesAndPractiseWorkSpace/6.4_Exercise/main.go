package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Printf("%s %s is of age %d", p.first, p.last, p.age)

}

func main() {
	p1 := person{
		first: `Vishal`,
		last:  `Sahare`,
		age:   40,
	}

	p1.speak()
}
