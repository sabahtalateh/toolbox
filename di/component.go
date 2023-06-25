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

// With
//
// First parameter is a component constructor function parameter name
// Following parameters are functions returning components, structures or interfaces
func (d *сomponent) With(_ string, _ ...any) *сomponent {
	return nil
}
