package acho

import (
	"log"
	//	"math"
)

func (m *Map) PopulateWithCountries(conf MapConfiguration) {

	isPossible := true
	for i := 0; i < m.CountryCount && isPossible; i++ {
		isPossible = m.CreateCountry(m.CountrySize)
	}
}

func (m *Map) RoomForCountry() bool {
	return true
}

func (m *Map) CreateCountry(size int) bool {
	found := false
	coordx := m.Random.Intn(m.Width)
	coordy := m.Random.Intn(m.Height)

	// log.Printf("coordx=%d, coordy=%d", coordx, coordy)

	//countryWidth := int(math.Sqrt(float64(size)))
	//countryHeight := int(math.Sqrt(float64(size)))

	// For now just traverse and look for water tiles
	for i := 0; i < m.Height && !found; i++ {
		for j := 0; j < m.Width && !found; j++ {
			if m.Tiles[i][j].Type == TileWater {
				found = true
				coordx = i
				coordy = j
			}
		}
	}

	// If not suitable place was found we can go
	if !found {
		return false
	}

	// We found water, so create a country
	actualSize := size
	// TODO Generate random country code

	countryCode := 123
	m.growCountry(countryCode, coordx, coordy, &actualSize)

	// for i := coordx; i < m.Height && actualSize < size; i++ {
	log.Println("Outer loop")
	// 	for j := coordy; j < m.Width && actualSize < size; i++ {
	// 		log.Println("Inner loop")
	// 		m.Tiles[i][j] = Tile{TileGround}
	// 		actualSize++
	// 	}
	// }

	return true
}

func (m *Map) growCountry(code, x, y int, size *int) {
	var currentTile Tile

	m.Tiles[x][y] = Tile{Type: TileGround, CountryCode: code}
	*size--
	log.Println("size:", *size)
	if *size > 0 {
		// TODO The direction on which the country grows should be random

		if x-1 > 0 {
			currentTile = m.Tiles[x-1][y]
			if currentTile.Type == TileWater || currentTile.CountryCode == 0 {
				m.growCountry(code, x-1, y, size)
			}
		}
		if y-1 > 0 {
			currentTile = m.Tiles[x][y-1]
			if currentTile.Type == TileWater || currentTile.CountryCode == 0 {
				m.growCountry(code, x, y-1, size)
			}
		}
		if x+1 < m.Width {
			currentTile = m.Tiles[x+1][y]
			if currentTile.Type == TileWater || currentTile.CountryCode == 0 {
				m.growCountry(code, x+1, y, size)
			}
		}
		if y+1 > m.Height {
			currentTile = m.Tiles[x][y+1]
			if currentTile.Type == TileWater || currentTile.CountryCode == 0 {
				m.growCountry(code, x, y+1, size)
			}
		}
	}
}
