package genalgo

type ISelector interface {
	Mater([]BaseUnit) (A, B BaseUnit)
	Mutator([]BaseUnit) (A BaseUnit)
}

//гены похожи
type Inbreeding struct {
}

func (s *Inbreeding) Select() (A, B BaseUnit) {
	return *(new(BaseUnit)), *(new(BaseUnit))
}

//гены не похожи
type Outbreeding struct {
}

func (s *Outbreeding) Select() (A, B BaseUnit) {
	return *(new(BaseUnit)), *(new(BaseUnit))
}

//функции приспособленности похожи
type Associative struct {
}

func (s *Associative) Select() (A, B BaseUnit) {
	return *(new(BaseUnit)), *(new(BaseUnit))
}
