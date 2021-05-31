package genalgo

import (
	"sort"
)


type Elite struct {

}

func (t *Elite) Create (parents, child []BaseUnit) (generation []BaseUnit)  {

	sort.Sort(UnitSort(parents))
	sort.Sort(UnitSort(child))

	elite := parents[0]

	size := len(parents)


	var all []BaseUnit

	all = append(all, elite)
	all = append(all, child[size-1])

	return all
}
