package main

import (
	"fmt"
	"github.com/leandroguillen/acho/map"
)

func main() {
	m := acho.NewMap(40, 10)

	for i := 0; i < 1; i++ {
		m.CreateCountry(6)
	}

	fmt.Println(m.ToString())
}
