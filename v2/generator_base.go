package genalgo

import "math/rand"

func NewGenerator(len int) *Generator {
	return &Generator{Len: len}
}

//Base Generator
//Generator is used to form null generation
//Creates units with fixed-length chromosomes, each gene is either one or zero
//Length depends on Len
type Generator struct {
	Len int
}

func (g *Generator) Generate() []int {
	var cromo []int

	for i := 0; i < g.Len; i++ {
		j := rand.Intn(2)
		cromo = append(cromo, j)
	}

	return cromo
}

