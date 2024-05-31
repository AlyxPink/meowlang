package object

import "fmt"

const STRING_OBJ = "STRING"

type String struct {
	Value string
}

func (s *String) Type() ObjectType {
	return STRING_OBJ
}

func (s *String) Inspect() string {
	return fmt.Sprint(s.Value)
}
