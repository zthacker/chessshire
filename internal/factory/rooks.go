package factory

import (
	"ChessShire/internal/characters/basecharacter"
	"ChessShire/internal/characters/gamecharacter"
	"ChessShire/internal/common"
	"ChessShire/internal/components"
	"ChessShire/internal/components/attributes"
)

func NewDefaultRook() *gamecharacter.GameCharacter {
	dc := &gamecharacter.GameCharacter{
		BaseCharacter: &basecharacter.BaseCharacter{
			Name:     "default_rook",
			Position: 0,
		},
		PieceType:    common.PIECE_TYPE_ROOK,
		MovementType: common.MOVEMENT_TYPE_ROOK,
		Components:   make(map[string]components.Component),
	}

	dc.Components[common.ATTRIBUTES] = defaultRookAttributeComponent()

	return dc
}

func defaultRookAttributeComponent() components.Component {
	attrs := map[string]int{
		common.HEALTH:    400,
		common.MANA:      50,
		common.STRENGTH:  30,
		common.AGILITY:   20,
		common.INTELLECT: 5,
	}
	return attributes.NewAttributeComponent(attrs) //new component
}
