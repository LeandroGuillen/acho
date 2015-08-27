package main

import (
	"fmt"
	"github.com/leandroguillen/acho/map"
)

func main() {
	m, err := acho.NewMapFromFile("files/map.csv")

	if err != nil {
		fmt.Println(err)
		panic(1)
	}

	fmt.Println(m.ToString())

	//generate()
}

// Generate a new map with a few countries
func generate() {
	m := acho.NewMap(40, 10)

	for i := 0; i < 1; i++ {
		m.CreateCountry(6)
	}

	fmt.Println(m.ToString())

}
