package stack

import (
	"fmt"
	"reflect"

	"github.com/edikrisnayana/data-structure/intf"
)

type stack struct {
	itemType reflect.Type
	items    []intf.T
}

func Init() intf.Stack {
	var stack intf.Stack = new(stack)
	return stack
}

func (s *stack) IsEmpty() bool {
	return s.Size() == 0
}

func (s *stack) Peek() (intf.T, error) {
	if s.IsEmpty() {
		return nil, fmt.Errorf("stack is empty")
	}

	return s.items[s.Size()-1], nil
}

func (s *stack) Pop() (intf.T, error) {
	if s.IsEmpty() {
		return nil, fmt.Errorf("stack is empty")
	}

	index := s.Size() - 1
	item := (s.items)[index]
	(s.items)[index] = nil
	s.items = s.items[:index]
	return item, nil
}

func (s *stack) Push(item intf.T) error {
	if s.itemType == nil {
		s.itemType = reflect.ValueOf(item).Type()
	}

	if s.itemType != nil && s.itemType != reflect.ValueOf(item).Type() {
		return fmt.Errorf("item type is not consistent. Should use %v", s.itemType)
	}

	s.items = append(s.items, item)
	return nil
}

func (s *stack) Size() int {
	return len(s.items)
}
