package genalgo

//Interface for new Generation selection Schema
type ISchema interface {
	Create (parents, child []IUnit) (generation []IUnit)
}