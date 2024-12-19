package fighter

import (
	"ChessShire/internal/characters/gamecharacter"
	"ChessShire/internal/components/ability"
	"ChessShire/internal/components/attributes"
)

/*
Needs to be able to Attack, Cast, Defend, Run
Take damage
Level Up
Move
*/

type Fighter struct {
	*gamecharacter.GameCharacter
	attibutesComponent *attributes.AttributeComponent
	abilitiesComponent *ability.AbilityComponent
}

func (f *Fighter) SetAttributeComponent(ac *attributes.AttributeComponent) {
	f.attibutesComponent = ac
}

func (f *Fighter) GetAttributeComponent() *attributes.AttributeComponent {
	return f.attibutesComponent
}

func (f *Fighter) SetAbilityComponent(ab *ability.AbilityComponent) {
	f.abilitiesComponent = ab
}

func (f *Fighter) GetAbilityComponent() *ability.AbilityComponent {
	return f.abilitiesComponent
}
