package gamecharacter

import (
	"ChessShire/internal/characters/basecharacter"
	"ChessShire/internal/components"
)

type GameCharacter struct {
	*basecharacter.BaseCharacter
	PieceType    uint8
	MovementType uint8
	Components   map[string]components.Component
}

func (gc *GameCharacter) GetComponent(name string) components.Component {
	return gc.Components[name]
}

func (gc *GameCharacter) SetComponent(component components.Component) {
	gc.Components[component.Name()] = component
}
