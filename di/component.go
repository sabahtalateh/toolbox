package di

// Component
//
// Parameter is a function returning structure or interface
// Following parameters are "value-functions"
func Component(_ any, _ ...any) {
}

// NamedComponent
//
// First parameter is a component name
// Second parameter is a function returning structure or interface
// Following parameters are "value-functions"
func NamedComponent(_ string, _ any, _ ...any) {
}
