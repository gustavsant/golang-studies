package main

import "fmt"

type car struct {
	name  string
	doors int
	color string
}

type sedan struct {
	car         car
	luxuryModel bool
}

type pickup struct {
	car    car
	fourWD bool
}

func main() {
	corolla := sedan{
		car{
			"Toyota Corolla",
			4,
			"black",
		},
		false,
	}

	rampage := pickup{
		car{
			"RAM Rampage",
			4,
			"grey",
		},
		true,
	}

	fmt.Println(corolla, rampage)
}
