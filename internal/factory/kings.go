package factory

import (
	"ChessShire/internal/characters/basecharacter"
	"ChessShire/internal/characters/gamecharacter"
	"ChessShire/internal/common"
	"ChessShire/internal/components"
	"ChessShire/internal/components/attributes"
)

func NewDefaultKing() *gamecharacter.GameCharacter {
	dc := &gamecharacter.GameCharacter{
		BaseCharacter: &basecharacter.BaseCharacter{
			Name: "default_king",
		},
		PieceType:    common.PIECE_TYPE_KING,
		MovementType: common.MOVEMENT_TYPE_KING,
		Components:   make(map[string]components.Component),
	}

	dc.Components[common.ATTRIBUTES] = newDefaultKingAttributesComponent()
	return dc
}

func newDefaultKingAttributesComponent() components.Component {
	attrs := map[string]int{
		common.HEALTH:    500,
		common.MANA:      600,
		common.STRENGTH:  10,
		common.AGILITY:   5,
		common.INTELLECT: 50,
	}

	return attributes.NewAttributeComponent(attrs)
}
