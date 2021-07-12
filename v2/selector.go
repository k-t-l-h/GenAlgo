package genalgo

type ISelector interface {
	Mater([]IUnit) (A, B IUnit)
	Mutator([]IUnit) (A IUnit)
}

//гены похожи
type Inbreeding struct {
}

func (s *Inbreeding) Select() (A, B IUnit) {
	return *(new(IUnit)), *(new(IUnit))
}

//гены не похожи
type Outbreeding struct {
}

func (s *Outbreeding) Select() (A, B IUnit) {
	return *(new(IUnit)), *(new(IUnit))
}

//функции приспособленности похожи
type Associative struct {
}

func (s *Associative) Select() (A, B IUnit) {
	return *(new(IUnit)), *(new(IUnit))
}
