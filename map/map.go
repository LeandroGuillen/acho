package acho

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
			m.Data[i][j] = Tile{TileWater}
		}
	}

	m.Name = "Unnamed map"

	return &m
}

func (m *Map) PopulateWithCountries(conf MapConfiguration) {
	if conf.CountrySize == 1 {

	}
}

func (m *Map) SetTile(x, y int, t Tile) {
	m.Data[x][y] = t
}

func (m *Map) ToString() (s string) {
	s = ""
	for i := 0; i < m.Height; i++ {
		for j := 0; j < m.Width; j++ {
			s += m.Data[i][j].ToString() + " "
		}
		s += "\n"
	}
	return
}
