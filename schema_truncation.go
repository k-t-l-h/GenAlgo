package genalgo

import "sort"

type Truncation struct {

}

func (t *Truncation) Create (parents, child []BaseUnit) (generation []BaseUnit)  {
	size := len(parents)

	var all []BaseUnit

	all = append(all, parents...)
	all = append(all, child...)

	sort.Sort(UnitSort(all))

	all = all[:size]

	return all
}