package toolbox

type С struct {
}

// Component
//
// Parameter is a function returning structure or interface
func Component(_ any) *С {
	return nil
}

// Name
//
// Specify component name
func (d *С) Name(_ string) *С {
	return nil
}

// With
//
// get parameter is a component constructor function parameter name
// Following parameters are functions returning components, structures or interfaces
func (d *С) With(_ string, _ any, _ ...any) *С {
	return nil
}
