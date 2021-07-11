package genalgo

import (
	"math/rand"
)

type Panmixia struct {
}

func (s *Panmixia) Mater(population []IUnit) (A, B IUnit) {
	size := len(population)

	i := rand.Intn(size)
	j := rand.Intn(size)

	return population[i], population[j]
}

func (s *Panmixia) Mutator(population []IUnit) (A IUnit) {
	size := len(population)

	i := rand.Intn(size)

	return population[i]
}
