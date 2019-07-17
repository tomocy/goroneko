package goroneko

type Player struct {
	Name       string
	Characters []*Character
}

func (p *Player) Roll(rollable rollable) *Character {
	return rollable.Roll()
}

type rollable interface {
	Roll() *Character
}
