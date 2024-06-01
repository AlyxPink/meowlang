package object

type ObjectType string

type Object interface {
	Type() ObjectType
	Inspect() string
}

// For variable storage
type Environment struct {
	store map[string]Object
	outer *Environment
}

func NewEnvironment() *Environment {
	return &Environment{store: make(map[string]Object)}
}

// Creates a new enclosed environment with an outer environment
func NewEnclosedEnvironment(outer *Environment) *Environment {
	env := NewEnvironment()
	env.outer = outer
	return env
}

// Retrieves a variable's value
func (e *Environment) Get(name string) (Object, bool) {
	obj, ok := e.store[name]
	if !ok && e.outer != nil {
		obj, ok = e.outer.Get(name)
	}
	return obj, ok
}

// Sets a variable's value
func (e *Environment) Set(name string, val Object) Object {
	e.store[name] = val
	return val
}
