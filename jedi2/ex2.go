package jedi2

import (
	"fmt"
)

type person struct {
	first string
	last  string
	age   int
}

// Ex2 tbd
func Ex2() {
	fmt.Println("\nex2")

	p := person{
		first: "Jack",
		last:  "Mental Illness",
		age:   32,
	}

	fmt.Println(p)

	changeMe(&p)

	fmt.Println(p)
}

func changeMe(p *person) {
	(*p).first = "Tyler"
}
