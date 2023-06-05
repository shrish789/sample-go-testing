package testing

// MyDependency is a dependency with a simple function.
type MyDependency interface {
    GetMessage() (string, error)
}

// RealDependency is the real implementation of MyDependency.
type RealDependency struct{}

// GetMessage is the implementation of GetMessage for RealDependency.
func (r *RealDependency) GetMessage() (string, error) {
    return "Hello, World!", nil
}
