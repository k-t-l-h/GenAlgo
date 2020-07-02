package genalgo

func NewUnit() *BaseUnit {
	return &BaseUnit{}
}

type BaseUnit struct {
	cromo   []int
	fitness float64
}

func (bi *BaseUnit) GetCromosomes() []int {
	return bi.cromo
}

func (bi *BaseUnit) SetCromosomes(cromo []int) {
	bi.cromo = cromo
}

func (bi *BaseUnit) GetFitness() float64 {
	return bi.fitness
}

func (bi *BaseUnit) SetFitness(fitness float64) {
	bi.fitness = fitness
}

