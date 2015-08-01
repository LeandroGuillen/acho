package main

import (
	"fmt"
	"github.com/leandroguillen/acho/map"
)

func main() {
	m := acho.NewMap(3, 3)
	fmt.Println(m.Data[1][1])
}
