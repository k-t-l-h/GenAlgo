package genalgo

type UnitSort []IUnit

func (us UnitSort) Len() int {
	return len(us)
}

func (us UnitSort) Less(i, j int) bool {
	return (us)[i].GetFitness() > (us)[j].GetFitness()
}

func (us UnitSort) Swap(i, j int) {
	(us)[i], (us)[j] = (us)[j], (us)[i]
}
