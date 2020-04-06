package main

import "fmt"

func main() {

	oTruck := struct {
		doors     int
		color     string
		fourWheel bool
	}{
		doors:     2,
		color:     `white`,
		fourWheel: true,
	}

	oSedan := struct {
		doors  int
		color  string
		luxury bool
	}{
		doors:  4,
		color:  `Black`,
		luxury: false,
	}

	fmt.Println(oTruck)
	fmt.Println(oSedan)
	fmt.Println(oTruck.doors)
	fmt.Println(oSedan.doors)
}
