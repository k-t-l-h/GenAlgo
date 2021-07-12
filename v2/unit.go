package genalgo

type IUnit interface {
	GetCromosomes() []int
	SetCromosomes([]int)

	GetFitness() float64
	SetFitness(fitness float64)
}

