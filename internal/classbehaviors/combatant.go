package classbehaviors

type Combatant interface {
	Attack() string                 // Melee attack, uses a weapon if equipped
	Cast(abilityName string) string // Use an ability, may or may not require a target
	Defend() string                 // Defensive action
	Run() string                    // Flee from combat
}
