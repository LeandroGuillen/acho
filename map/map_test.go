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

func TestMapToString(t *testing.T) {
	// Create NxM map
	m := NewMap(3, 3)

	m.SetTile(0, 1, Tile{TileGround})
	m.SetTile(0, 2, Tile{TileGround})
	m.SetTile(1, 1, Tile{TileGround})

	expected := "~ O O \n~ O ~ \n~ ~ ~ \n"
	output := m.ToString()

	if output != expected {
		t.Errorf("expected:\n%s\n\ngot:\n%s", expected, output)
	}
}
