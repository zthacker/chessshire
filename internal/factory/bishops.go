package factory

import (
	"ChessShire/internal/characters/basecharacter"
	"ChessShire/internal/characters/gamecharacter"
	"ChessShire/internal/classbehaviors/abilities/fireball"
	"ChessShire/internal/classbehaviors/effects/firedamage"
	"ChessShire/internal/common"
	"ChessShire/internal/components"
	"ChessShire/internal/components/ability"
	"ChessShire/internal/components/attributes"
)

func NewDefaultBishop() *gamecharacter.GameCharacter {
	dc := &gamecharacter.GameCharacter{
		BaseCharacter: &basecharacter.BaseCharacter{
			Name: "default_bishop",
		},
		PieceType:    common.PIECE_TYPE_BISHOP,
		MovementType: common.MOVEMENT_TYPE_BISHOP,
		Components:   make(map[string]components.Component),
	}

	dc.Components[common.ATTRIBUTES] = newDefaultBishopAttributesComponent()
	return dc
}

func NewAdeptBishop() *gamecharacter.GameCharacter {
	ab := &gamecharacter.GameCharacter{
		BaseCharacter: &basecharacter.BaseCharacter{
			Name: "adept_bishop",
		},
		PieceType:    common.PIECE_TYPE_BISHOP,
		MovementType: common.MOVEMENT_TYPE_BISHOP,
		Components:   make(map[string]components.Component),
	}

	//ability component and ability
	abilityComponent := ability.NewAbilityComponent()
	fireBall := fireball.Fireball{}
	fireBall.ManaCost = 10
	fireBall.AddEffect(&firedamage.FireDamage{BaseDamage: 10})
	abilityComponent.AddAbility("fireball", &fireBall)
	ab.Components[common.ABILITIES] = abilityComponent

	//attrs component
	ab.Components[common.ATTRIBUTES] = newDefaultBishopAttributesComponent()

	return ab
}

func newDefaultBishopAttributesComponent() components.Component {
	attrs := map[string]int{
		common.HEALTH:    200,
		common.MANA:      600,
		common.STRENGTH:  5,
		common.AGILITY:   5,
		common.INTELLECT: 40,
	}

	return attributes.NewAttributeComponent(attrs)
}
