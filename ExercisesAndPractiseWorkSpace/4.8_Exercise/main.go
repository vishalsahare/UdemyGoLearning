package main

import "fmt"

func main() {
	mmap := map[string][]string{
		`bond_james`:      []string{`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: []string{`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           []string{`Being evil`, `Ice cream`, `Sunsets`},
	}

	for k, v := range mmap {
		fmt.Println("record: ", k)

		for i, vv := range v {
			fmt.Println("\t", i, vv)
		}
	}
}
