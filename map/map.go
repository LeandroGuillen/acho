package acho

import (
	"bytes"
	"encoding/csv"
	"io/ioutil"
	"math/rand"
	"strconv"
	"time"
)

type Map struct {
	//IslandStyle string
	Width  int
	Height int
	Tiles  [][]Tile
	Name   string
	MapConfiguration
}

type MapConfiguration struct {
	CountryCount              int
	CountrySize               int
	CountrySizeVariance       int
	CountryBorderSoftness     float64
	NeutralTroopCount         int
	NeutralTroopCountVariance int
	NeutralTroopPercentage    float64
	WaterSize                 int
	WaterSizeVariance         int
	Random                    *rand.Rand
}

func (m *Map) SetTile(x, y int, t Tile) {
	m.Tiles[x][y] = t
}

func (m *Map) ToString() (s string) {
	s = ""
	for i := 0; i < m.Height; i++ {
		for j := 0; j < m.Width; j++ {
			s += m.Tiles[i][j].ToString() + " "
		}
		s += "\n"
	}
	return
}

// Create a map in memory with the specified width and height.
// The tiles are all set to empty
func NewMap(width, height int) *Map {
	m := Map{
		Width:  width,
		Height: height,
	}

	m.Tiles = make([][]Tile, height)
	for i := 0; i < height; i++ {
		m.Tiles[i] = make([]Tile, width)
		for j := 0; j < width; j++ {
			m.Tiles[i][j] = Tile{TileWater, 0}
		}
	}

	m.Name = "Unnamed map"

	s1 := rand.NewSource(time.Now().UnixNano())
	m.Random = rand.New(s1)

	return &m
}

func (m *Map) GetRandomTile() (x, y int, t *Tile) {
	x = m.Random.Intn(m.Height)
	y = m.Random.Intn(m.Width)
	t = &m.Tiles[x][y]
	return
}

// Read a map from a CSV file
func NewMapFromFile(filename string) (*Map, error) {
	var content []byte
	var err error

	if content, err = ioutil.ReadFile(filename); err != nil {
		return nil, err
	}
	r := csv.NewReader(bytes.NewReader(content))

	// Read the first line, it contains the height and width
	record, err := r.Read()
	if err != nil {
		return nil, err
	}
	height, err := strconv.Atoi(record[0])
	if err != nil {
		return nil, err
	}
	width, err := strconv.Atoi(record[1])
	if err != nil {
		return nil, err
	}

	// Reserve memory for our map
	theMap := NewMap(width, height)

	records, err := r.ReadAll()
	if err != nil {
		return nil, err
	}

	// Create individual tiles, assigns countryID to tiles
	for i, row := range records {
		for j, value := range row {
			tileValue, err := strconv.Atoi(value)
			if err != nil {
				return nil, err
			}
			t := Tile{TileGround, tileValue}
			theMap.SetTile(i, j, t)
		}
	}

	return theMap, nil
}
