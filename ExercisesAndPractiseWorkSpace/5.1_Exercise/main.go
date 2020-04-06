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

	var lpeople = []people{
		people{
			first:   `Vishal`,
			last:    `Sahare`,
			flavors: []string{`Vanila`, `strawberry`},
		},

		people{
			first:   `Amit`,
			last:    `Sahare`,
			flavors: []string{`Chocolate`, `pista`},
		},
	}

	for _, indi := range lpeople {
		fmt.Printf("%s %s: ", indi.first, indi.last)

		for _, flavor := range indi.flavors {
			fmt.Print(flavor)
			fmt.Print(" ")
		}
		fmt.Println()
	}
}
