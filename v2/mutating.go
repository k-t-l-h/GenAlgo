package genalgo

import (
	"math/rand"
)

type IMutate interface {
	Mutate(individuals IUnit) (mutant IUnit)
	GetSpeed() float64
}

type OneDotMutatation struct {
	Probability     float64
	ProbabilityFunc func() float64
	//TODO: добавить модуль
}

func (odm *OneDotMutatation) Mutate(parent IUnit) (child IUnit) {

	/*check := odm.ProbabilityFunc()
	if check < odm.Probability {
		return parent
	}

	*/

	cromo := parent.GetCromosomes()

	r := rand.Intn(len(cromo))

	n := cromo[r]

	if n == 0 {
		n++
		cromo[r] = n
	} else {
		cromo[r] = 0
	}

	child = &BaseUnit{}

	child.SetCromosomes(cromo)

	return child
}

func (odm *OneDotMutatation) GetSpeed() float64 {
	return odm.Probability
}
