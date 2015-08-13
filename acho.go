package main

import (
	"fmt"
	"github.com/leandroguillen/acho/map"
)

func main() {

	conf := acho.MapConfiguration{
		CountryCount:              10,
		CountrySize:               10,
		CountrySizeVariance:       3,
		CountryBorderSoftness:     0.1,
		NeutralTroopCount:         5,
		NeutralTroopCountVariance: 3,
		NeutralTroopPercentage:    0.75,
		WaterSize:                 10,
		WaterSizeVariance:         2,
	}

	m := acho.NewMap(&conf)
	m.Name = "My First Map"

	fmt.Println(m.ToString())
}
