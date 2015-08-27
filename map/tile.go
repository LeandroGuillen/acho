package acho

import (
	"strconv"
)

type TileType int

const (
	TileWater = iota
	TileGround
)

type Tile struct {
	Type        TileType
	CountryCode int
}

func (t *Tile) ToString() (s string) {

	switch t.Type {
	case TileWater:
		s = "~"
		break
	case TileGround:
		s = strconv.Itoa(t.CountryCode)
		break
	default:
		s = "?"
	}

	return
}
