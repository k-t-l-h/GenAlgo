package genalgo

//Uniform Crossover
//Probability determines the probability that the gene of the first parent unit will be selected
//F. e. if x < Probability, unit A gene is selected, otherwise unit B gene is selected
//ProbabilityFunc sets the distribution of a random variable, so it can be changed to fit your needs

type UniformCrossover struct {
	Probability     float64
	ProbabilityFunc func() float64
}

func (uc *UniformCrossover) GetSpeed() float64 {
	return uc.Probability
}

func (uc *UniformCrossover) Cross(A, B BaseUnit) (C, D *BaseUnit) {

	parentA := A.GetCromosomes()
	parentB := B.GetCromosomes()

	var CA, CB []int

	for i := 0; i < len(parentA); i++ {
		pA := uc.ProbabilityFunc()
		pB := uc.ProbabilityFunc()

		if pA < uc.Probability {
			CA = append(CA, parentA[i])
		} else {
			CA = append(CA, parentB[i])
		}

		if pB < uc.Probability {
			CB = append(CB, parentB[i])
		} else {
			CB = append(CB, parentA[i])
		}
	}

	childA := BaseUnit{}
	childB := BaseUnit{}

	childA.SetCromosomes(CA)
	childB.SetCromosomes(CB)

	return &childA, &childB
}

