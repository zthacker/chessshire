package factory

import (
	"ChessShire/internal/characters/basecharacter"
	"ChessShire/internal/characters/gamecharacter"
	"ChessShire/internal/common"
	"ChessShire/internal/components"
	"ChessShire/internal/components/attributes"
)

func NewDefaultQueen() *gamecharacter.GameCharacter {
	dc := &gamecharacter.GameCharacter{
		BaseCharacter: &basecharacter.BaseCharacter{
			Name: "default_queen",
		},
		PieceType:    common.PIECE_TYPE_QUEEN,
		MovementType: common.MOVEMENT_TYPE_QUEEN,
		Components:   make(map[string]components.Component),
	}

	dc.Components[common.ATTRIBUTES] = newDefaultQueenAttributesComponent()
	return dc
}

func newDefaultQueenAttributesComponent() components.Component {
	attrs := map[string]int{
		common.HEALTH:    1000,
		common.MANA:      600,
		common.STRENGTH:  30,
		common.AGILITY:   30,
		common.INTELLECT: 50,
	}

	return attributes.NewAttributeComponent(attrs)
}
