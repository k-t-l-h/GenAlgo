package genalgo

import "sort"

type Truncation struct {

}

func (t *Truncation) Create (parents, child []IUnit) (generation []IUnit)  {
	size := len(parents)

	var all []IUnit

	all = append(all, parents...)
	all = append(all, child...)

	sort.Sort(UnitSort(all))

	all = all[:size]

	return all
}