package genalgo

type IMutate interface {
	Mutate(individuals BaseUnit) (mutant BaseUnit)
	GetSpeed() float64
}
