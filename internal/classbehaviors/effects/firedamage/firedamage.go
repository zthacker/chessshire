package firedamage

import (
	"ChessShire/internal/classbehaviors"
	"ChessShire/internal/common"
	"fmt"
)

type FireDamage struct {
	BaseDamage int
}

func (f *FireDamage) CalculateDamage() int {
	if f.BaseDamage < 0 {
		f.BaseDamage = 0
	}

	return f.BaseDamage
}

func (f *FireDamage) ApplyEffect(targets []classbehaviors.Attributes) string {
	result := ""
	for _, target := range targets {
		dmg := f.CalculateDamage()
		target.Apply(common.HEALTH, -dmg, common.RES_FIRE)
		result += fmt.Sprintf("Target takes %d fire damage!\n", dmg)
	}
	return result
}
