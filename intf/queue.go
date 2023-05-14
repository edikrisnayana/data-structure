package intf

type Queue interface {
	IsEmpty() bool
	Offer(item T) error
	Peek() (T, error)
	Poll() (T, error)
	Size() int
}
