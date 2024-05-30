package object

type ObjectType string

const (
	INTEGER_OBJ = "INTEGER"
)

type Object interface {
	Type() ObjectType
	Inspect() string
}

type Environment struct {
	store map[string]Object
}

func NewEnvironment() *Environment {
	return &Environment{store: make(map[string]Object)}
}

func (e *Environment) Get(name string) (Object, bool) {
	obj, ok := e.store[name]
	return obj, ok
}

func (e *Environment) Set(name string, val Object) Object {
	e.store[name] = val
	return val
}
