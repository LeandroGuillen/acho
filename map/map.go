package acho

type Tile int

const (
	TileWater = iota
	TileGround
)

type Map struct {
	//IslandStyle string
	Width  int
	Height int
	Data   [][]Tile
	Name   string
}

type MapConfiguration struct {
	CountrySize               int
	CountrySizeVariance       int
	CountryBorderSoftness     float64
	NeutralTroopCount         int
	NeutralTroopCountVariance int
	NeutralTroopPercentage    float64
	WaterSize                 int
	WaterSizeVariance         int
}

func NewMap(width, height int) *Map {
	m := Map{
		Width:  width,
		Height: height,
	}

	m.Data = make([][]Tile, height)
	for i := 0; i < height; i++ {
		m.Data[i] = make([]Tile, width)
		for j := 0; j < width; j++ {
			m.Data[i][j] = TileWater
		}
	}

	m.Name = "Unnamed map"

	return &m
}

func (m *Map) PopulateWithCountries(conf MapConfiguration) {
	if conf.CountrySize == 1 {

	}
}
