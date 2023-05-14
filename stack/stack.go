package stack

import (
	"fmt"
	"reflect"
)

type Stack interface {
	Size() int
	IsEmpty() bool
	Push(item T) error
	Pop() (T, error)
}

type stackItems []T

func NewStack() Stack {
	var stack Stack = new(stackItems)
	return stack
}

func (s *stackItems) Size() int {
	return len(*s)
}

func (s *stackItems) IsEmpty() bool {
	return s.Size() == 0
}

func (s *stackItems) Push(item T) error {
	if !s.IsEmpty() {
		firstItemType := reflect.ValueOf((*s)[0]).Type()
		if firstItemType != reflect.ValueOf(item).Type() {
			return fmt.Errorf("item type is not consistent. Should use %v", firstItemType)
		}
	}
	*s = append(*s, item)
	return nil
}

func (s *stackItems) Pop() (T, error) {
	if s.IsEmpty() {
		return nil, fmt.Errorf("stack is empty")
	}

	index := s.Size() - 1
	item := (*s)[index]
	(*s)[index] = nil
	*s = (*s)[:index]
	return item, nil
}
