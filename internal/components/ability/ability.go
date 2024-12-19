package ability

import (
	"ChessShire/internal/classbehaviors"
	"fmt"
)

// AbilityComponent acts as a Container for Ability behaviors
type AbilityComponent struct {
	Abilities map[string]classbehaviors.Ability
}

func NewAbilityComponent() *AbilityComponent {
	return &AbilityComponent{Abilities: map[string]classbehaviors.Ability{}}
}

// implements Component interface
func (a *AbilityComponent) Name() string {
	return "Abilities"
}

func (a *AbilityComponent) AddAbility(name string, ability classbehaviors.Ability) {
	a.Abilities[name] = ability
}

func (a *AbilityComponent) UseAbility(user classbehaviors.Attributes, name string, targets []classbehaviors.Attributes) string {
	if user == nil {
		return "ability user Attributes is nil"
	}
	if theAbility, ok := a.Abilities[name]; ok {
		return theAbility.Use(user, targets)
	}
	return fmt.Sprintf("Ability %s not found!", name)
}
