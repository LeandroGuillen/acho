package acho

import (
	"testing"
)

func TestMapCreation(t *testing.T) {
	// Create NxM map
	m := NewMap(10, 10)

	// Any random position should be a water tile
	if m.Data[3][5].Type != TileWater {
		t.Error("m.Data[3][5] should equal to", TileWater, ", not", m.Data[3][5].Type)
	}
}
