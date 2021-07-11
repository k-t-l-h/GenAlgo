package main

import (
	genalg "github.com/k-t-l-h/GenAlgo"
	"golang.org/x/exp/rand"
	"math"
)

func main() {

	sel := genalg.Panmixia{}
	ga := genalg.GenAlgo{
		MaxIteration: 4,
		Generator:    &genalg.Generator{20},
		Crossover: &genalg.UniformCrossover{
			Probability:     0.5,
			ProbabilityFunc: rand.Float64,
		},
		Mutate: &genalg.OneDotMutatation{
			Probability:     0.5,
			ProbabilityFunc: rand.Float64,
		},
		Schema: &genalg.Truncation{},
		Fitness: func(unit genalg.BaseUnit) float64 {
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
	ga.OnBegin = func() {}
	ga.OnEnd = func() {}
	ga.Init(20)
	ga.Simulation()

}
