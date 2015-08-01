package acho

import (
	"testing"
)

func TestTileToString(t *testing.T) {
	var water, ground TileType
	water = TileWater
	ground = TileGround

	var tile Tile

	tile.Type = water
	if tile.ToString() != "~" {
		t.Error("Water tile should be printed as '~', got", tile.ToString())
	}

	tile.Type = ground
	if tile.ToString() != "O" {
		t.Error("Ground tile should be printed as 'O', got", tile.ToString())
	}
}
