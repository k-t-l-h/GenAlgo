package genalgo

type UnitSort []BaseUnit

func (us UnitSort) Len() int {
	return len(us)
}

func (us UnitSort) Less(i, j int) bool {
	return (us)[i].fitness > (us)[j].fitness
}

func (us UnitSort) Swap(i, j int) {
	(us)[i], (us)[j] = (us)[j], (us)[i]
}
