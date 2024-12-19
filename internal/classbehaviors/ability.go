package classbehaviors

/*
Abilities orchestrate multiple effects and apply
them to the appropriate targets. For example, Fireball
coordinates FireDamage, ChainAttack, and AoEAttack
*/

type Ability interface {
	AddEffect(effect Effect)
	Use(user Attributes, targets []Attributes) string
}
