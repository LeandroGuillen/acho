package acho

import (
	"log"
	"math"
	"math/rand"
	"strconv"
	"time"
)

type Country struct {
	Id        int
	Adjacency []int
}

type Map struct {
	Name      string
	Countries []Country
	*MapConfiguration
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

func NewMap(conf *MapConfiguration) *Map {
	seed := rand.NewSource(time.Now().UnixNano())

	// Basic Values
	m := Map{Name: "Unnamed map"}
	m.MapConfiguration = conf
	m.Random = rand.New(seed)

	// Create all countries assigning an id
	m.Countries = make([]Country, m.CountryCount)
	for i := 0; i < m.CountryCount; i++ {
		m.Countries[i] = Country{Id: i}
	}

	for i := 0; i < (m.CountryCount/2)+1; i++ {
		nConnections := m.Random.Intn(int(math.Max(float64(m.CountryCount)+2, 7))) + 2

		m.Countries[i].Adjacency = make([]int, nConnections)
		for j := 0; j < nConnections; j++ {
			newC := m.Random.Intn(m.CountryCount)
			m.Countries[i].Adjacency[j] = newC
		}
	}

	m.cleanup()
	// log.Println(m.Countries)

	for i := 0; i < (m.CountryCount/2)+1; i++ {
		for j := 0; j < len(m.Countries[i].Adjacency); j++ {
			c := m.Countries[i].Adjacency[j]
			m.Countries[c].Adjacency = append(m.Countries[c].Adjacency, i)
		}
	}

	m.cleanup()

	log.Println(m.Countries)

	return &m
}

// Remove duplicated connections or connections to self
func (m *Map) cleanup() {
	for i := 0; i < len(m.Countries); i++ {
		newAdj := make([]int, 0)
		present := make([]bool, m.CountryCount)
		for j, _ := range present {
			present[j] = false
		}

		present[i] = true
		for _, a := range m.Countries[i].Adjacency {
			if !present[a] {
				newAdj = append(newAdj, a)
				present[a] = true
			}
		}

		m.Countries[i].Adjacency = make([]int, len(newAdj))
		for j := 0; j < len(newAdj); j++ {
			m.Countries[i].Adjacency[j] = newAdj[j]
		}
	}
}

func (m *Map) ToString() string {
	return m.Name + " (" + strconv.Itoa(m.CountryCount) + " countries)"
}
