package main

import "fmt"

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {

	oTruck := truck{
		vehicle: vehicle{
			doors: 2,
			color: `white`,
		},
		fourWheel: true,
	}

	oSedan := sedan{
		vehicle: vehicle{
			doors: 4,
			color: `Black`,
		},
		luxury: false,
	}

	fmt.Println(oTruck)
	fmt.Println(oSedan)
	fmt.Println(oTruck.doors)
	fmt.Println(oSedan.doors)
}
