package classbehaviors

/*
Calculation Happens in Effects:

    Effects (e.g., FireDamage, AoEAttack) calculate damage based
	on their logic. For example, AoEAttack might reduce damage
	per jump in a chain.
*/

type Effect interface {
	CalculateDamage() int                    // Calculates damage for a specific target
	ApplyEffect(targets []Attributes) string // Applies the effect to one or more targets
}
