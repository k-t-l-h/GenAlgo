package genalgo

//Crossover Operator Interface
//The Cross() function takes two parent units and returns two descendant units
//Speed determines how many descendant units will be created
type ICrossover interface {
	Cross(A, B BaseUnit) (C, D *BaseUnit)
	GetSpeed() float64
}



