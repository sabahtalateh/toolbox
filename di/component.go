package di

type сomponent struct {
}

// Component
//
// Parameter is a function returning structure or interface
func Component(_ any) *сomponent {
	return nil
}

// Name
//
// Specify component name
func (d *сomponent) Name(_ string) *сomponent {
	return nil
}

// Value
//
// First parameter is a parameter name
// Second parameter is a function providing corresponding value
func (d *сomponent) Value(_ string, _ any) *сomponent {
	return nil
}
