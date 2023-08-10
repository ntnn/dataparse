package dataparse

type Value struct {
	Data any
}

func NewValue(data any) Value {
	return Value{Data: data}
}

//go:generate go run ./value_numbers.go
