package factory

import (
	"ChessShire/internal/characters/basecharacter"
	"ChessShire/internal/characters/gamecharacter"
	"ChessShire/internal/common"
	"ChessShire/internal/components"
	"ChessShire/internal/components/attributes"
)

func NewDefaultKnight() *gamecharacter.GameCharacter {
	dc := &gamecharacter.GameCharacter{
		BaseCharacter: &basecharacter.BaseCharacter{
			Name:     "default_knight",
			Position: 0,
		},
		PieceType:    common.PIECE_TYPE_KNIGHT,
		MovementType: common.MOVEMENT_TYPE_KNIGHT,
		Components:   make(map[string]components.Component),
	}

	dc.Components[common.ATTRIBUTES] = defaultKnightAttributeComponent()

	return dc
}

func defaultKnightAttributeComponent() components.Component {
	attrs := map[string]int{
		common.HEALTH:    300,
		common.MANA:      50,
		common.STRENGTH:  15,
		common.AGILITY:   30,
		common.INTELLECT: 15,
	}
	return attributes.NewAttributeComponent(attrs) //new component
}
