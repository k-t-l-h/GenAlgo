package genalgo

import (
	"math/rand"
)

type Panmixia struct {
}

func (s *Panmixia) Mater(population []BaseUnit) (A, B BaseUnit) {
	size := len(population)

	i := rand.Intn(size)
	j := rand.Intn(size)

	return population[i], population[j]
}

func (s *Panmixia) Mutator(population []BaseUnit) (A BaseUnit) {
	size := len(population)

	i := rand.Intn(size)

	return population[i]
}
