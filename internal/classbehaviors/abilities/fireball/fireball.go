package fireball

import (
	"ChessShire/internal/classbehaviors"
	"ChessShire/internal/common"
	"fmt"
)

// Fireball implements Ability interface
type Fireball struct {
	ManaCost int
	Effects  []classbehaviors.Effect
}

func (f *Fireball) AddEffect(eff classbehaviors.Effect) {
	f.Effects = append(f.Effects, eff)
}

func (f *Fireball) Use(user classbehaviors.Attributes, targets []classbehaviors.Attributes) string {
	mana := user.Get(common.MANA)
	if mana < f.ManaCost {
		return "not enough mana"
	}

	user.Modify(common.MANA, -f.ManaCost)

	result := fmt.Sprintf("Casting Fireball (Mana Cost: %d):\n", f.ManaCost)
	for _, eff := range f.Effects {
		result += eff.ApplyEffect(targets)
	}

	return result
}
