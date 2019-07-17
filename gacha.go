package goroneko

import "math/rand"

type Gacha struct {
	Name       string
	Characters []*Character
}

func (g *Gacha) Roll() *Character {
	n := len(g.Characters)
	x := rand.Intn(n)

	return g.Characters[x]
}
