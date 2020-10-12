package main

import (
	genalgo2 "genalgo"
	"golang.org/x/exp/rand"
	"math"
)

func main() {

	sel := genalgo2.Panmixia{}
	ga := genalgo2.GenAlgo{
		MaxIteration: 4,
		Generator:    &genalgo2.Generator{20},
		Crossover: &genalgo2.UniformCrossover{
			Probability:     0.5,
			ProbabilityFunc: rand.Float64,
		},
		Mutate: &genalgo2.OneDotMutatation{
			Probability:     0.5,
			ProbabilityFunc: rand.Float64,
		},
		Schema: &genalgo2.Truncation{},
		Fitness: func(unit genalgo2.BaseUnit) float64 {
			cr := unit.GetCromosomes()
			fitness := 0.0
			x := 0
			y := 0

			for i := 0; i < len(cr)/2; i++ {
				x += cr[i]
			}

			for i := len(cr) / 2; i < len(cr); i++ {
				y += cr[i]
			}

			fitness = math.Cos(float64(x)) + math.Cos(float64(y))
			return fitness
		},
		Select: &sel,
	}
	ga.Init(20)
	ga.Simulation()

}
