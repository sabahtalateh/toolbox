package di

type сomponent struct {
}

// Component
//
// Parameter is a function returning structure or interface
func Component(_ any) *сomponent {
	return nil
}

// NamedComponent
//
// First parameter is a component name
// Second parameter is a function returning structure or interface
func NamedComponent(_ string, _ any) *сomponent {
	return nil
}

// Value
//
// First parameter is a parameter name
// Second parameter is a function providing corresponding value
func (d *сomponent) Value(_ string, _ any) {
}
