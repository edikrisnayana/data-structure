package intf

type Stack interface {
	IsEmpty() bool
	Peek() (T, error)
	Pop() (T, error)
	Push(item T) error
	Size() int
}
