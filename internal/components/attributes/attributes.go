package attributes

import (
	"fmt"
	"log"
)

type AttributeComponent struct {
	attributes map[string]int
}

func NewAttributeComponent(initialValues map[string]int) *AttributeComponent {
	return &AttributeComponent{attributes: initialValues}
}

// Implements the Component interface
func (a *AttributeComponent) Name() string {
	return "AttributesComponent"
}

// Get retrieves the value of an attribute
func (a *AttributeComponent) Get(attribute string) int {
	if value, ok := a.attributes[attribute]; ok {
		return value
	}
	return 0 // Default if attribute doesn't exist
}

// Set sets the value of an attribute
func (a *AttributeComponent) Set(attribute string, value int) {
	a.attributes[attribute] = value
}

// Modify adjusts the value of an attribute by a delta
func (a *AttributeComponent) Modify(attribute string, delta int) {
	previous := a.attributes[attribute]
	log.Println("pre modify -- %s: %s", attribute, previous)
	a.attributes[attribute] += delta
	if a.attributes[attribute] < 0 {
		a.attributes[attribute] = 0 // Prevent negative values
	}
	post := a.attributes[attribute]
	log.Println("post modify -- %s: %s", attribute, post)
}

// Apply modifies an attribute, respecting resistances if applicable
func (a *AttributeComponent) Apply(attribute string, amount int, resistanceAttr string) {
	//if resistanceAttr != "" {
	//	resistance := a.Get(resistanceAttr)
	//	amount -= resistance
	//	if amount < 0 {
	//		amount = 0
	//	}
	//}
	log.Println(fmt.Sprintf("applying %s damage to %s", amount, attribute))
	a.Modify(attribute, amount)
}

// IsZero checks if a specific attribute has reached zero (e.g., "health")
func (a *AttributeComponent) IsZero(attribute string) bool {
	return a.Get(attribute) == 0
}
