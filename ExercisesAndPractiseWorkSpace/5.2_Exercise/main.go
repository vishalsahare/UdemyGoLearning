package main

import (
	"fmt"
)

type people struct {
	first   string
	last    string
	flavors []string
}

func main() {
	p1 := people{
		first:   `Vishal`,
		last:    `Sahare`,
		flavors: []string{`Vanila`, `strawberry`},
	}

	p2 := people{
		first:   `Amit`,
		last:    `Sakhare`,
		flavors: []string{`Chocolate`, `pista`},
	}

	myMap := map[string]people{
		p1.last: p1,
		p2.last: p2,
	}

	for _, v := range myMap {
		fmt.Printf("%s %s: ", v.first, v.last)

		for _, flavor := range v.flavors {
			fmt.Print(flavor)
			fmt.Print(" ")
		}
		fmt.Println()
	}
}
