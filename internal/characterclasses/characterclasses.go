package characterclasses

import (
	"ChessShire/internal/components/ability"
	"ChessShire/internal/components/attributes"
)

type CharacterClasses interface {
	SetAttributeComponent(ac *attributes.AttributeComponent)
	GetAttributeComponent() *attributes.AttributeComponent
	SetAbilityComponent(ab *ability.AbilityComponent)
	GetAbilityComponent() *ability.AbilityComponent
}
