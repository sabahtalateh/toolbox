package di

// Component
//
// Parameter is a function returning structure or interface
func Component(_ any) {
}

// NamedComponent
//
// First parameter is a component name
// Second parameter is a function returning structure or interface
func NamedComponent(_ string, _ any) {
}
