package genalgo

import (
	"math/rand"
	"sort"
)

//N-Point Crossover
//N determines the number of break points
//If N = 1, operator is expressed in single point crossover
//Probability determines the probability that the gene of the first parent unit will be selected
//F. e. if x < Probability, unit A gene is selected, otherwise unit B gene is selected
//ProbabilityFunc sets the distribution of a random variable, so it can be changed to fit your needs
type NPointCrossover struct {
	N               int
	Probability     float64
	ProbabilityFunc func() float64
}

func (np *NPointCrossover) Cross(A, B IUnit) (C, D IUnit) {
	probability := np.ProbabilityFunc()

	if probability < np.Probability {

		return A, B
	}

	parentA := (A).GetCromosomes()
	parentB := (B).GetCromosomes()

	//TODO: select size properly
	size := len(parentA)

	var dots []int
	for i := 0; i <= np.N; i++ {
		dots = append(dots, rand.Intn(size))
	}

	sort.Ints(dots)

	cdot := 0

	PA := []int{}
	PB := []int{}

	swA := parentA
	swB := parentB

	for i := range dots {
		PA = append(PA, swA[cdot:i]...)
		PB = append(PB, swB[cdot:i]...)
		cdot = i
		swA, swB = swB, swA
	}

	PA = append(PA, swA[cdot:]...)
	PB = append(PB, swB[cdot:]...)

	childA := BaseUnit{}
	childB := BaseUnit{}

	childA.SetCromosomes(PA)
	childB.SetCromosomes(PB)

	return &childA, &childB
}

func (np *NPointCrossover) GetSpeed() float64 {
	return np.Probability
}
