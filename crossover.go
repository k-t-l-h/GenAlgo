package genalgo

//Crossover Operator Interface
//The Cross() function takes two parent units and returns two descendant units
//Speed determines how many descendant units will be created
type ICrossover interface {
	Cross(A, B IUnit) (C, D IUnit)
	GetSpeed() float64
}



