package classbehaviors

type Attributes interface {
	Get(attribute string) int           // Retrieves the value of an attribute
	Set(attribute string, value int)    // Sets the value of an attribute
	Modify(attribute string, delta int) // Modifies the attribute by a delta
	Apply(attribute string, amount int, resistanceAttr string)
	IsZero(attribute string) bool
}
