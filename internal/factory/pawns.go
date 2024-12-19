package factory

import (
	"ChessShire/internal/characters/basecharacter"
	"ChessShire/internal/characters/gamecharacter"
	"ChessShire/internal/common"
	"ChessShire/internal/components"
	"ChessShire/internal/components/attributes"
)

func NewDefaultPawn() *gamecharacter.GameCharacter {
	dc := &gamecharacter.GameCharacter{
		BaseCharacter: &basecharacter.BaseCharacter{
			Name:     "default_pawn",
			Position: 0,
		},
		PieceType:    common.PIECE_TYPE_PAWN,
		MovementType: common.MOVEMENT_TYPE_PAWN,
		Components:   make(map[string]components.Component),
	}

	dc.Components[common.ATTRIBUTES] = defaultPawnAttributeComponent()

	return dc
}

func defaultPawnAttributeComponent() components.Component {
	attrs := map[string]int{
		common.HEALTH:    100,
		common.MANA:      100,
		common.STRENGTH:  20,
		common.AGILITY:   15,
		common.INTELLECT: 12,
	}
	return attributes.NewAttributeComponent(attrs) //new component
}
