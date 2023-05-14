package test

import (
	"fmt"
	"testing"

	"github.com/edikrisnayana/data-structure/stack"
)

func TestStackPushSecondPushSuccess(t *testing.T) {
	stack := stack.Init()
	err := stack.Push(1)
	if err != nil {
		t.Fatalf(err.Error())
	}
	err = stack.Push(2)
	if err != nil {
		t.Fatalf(err.Error())
	}
}

func TestStackPushSecondPushMismatchType(t *testing.T) {
	stack := stack.Init()
	err := stack.Push(1)
	if err != nil {
		t.Fatalf(err.Error())
	}
	err = stack.Push("2")
	fmt.Println(err.Error())
	if err == nil {
		t.Fatalf("checking type failed. should give error mismatch type")
	}
}

func TestStackPopSuccess(t *testing.T) {
	stack := stack.Init()
	stack.Push(1)
	stack.Push(2)
	currSize := stack.Size()
	_, err := stack.Pop()
	if err != nil {
		t.Fatalf(err.Error())
	}

	if currSize == stack.Size() {
		t.Fatalf("stack pop failed")
	}
}

func TestStackPopEmptyStack(t *testing.T) {
	stack := stack.Init()
	_, err := stack.Pop()
	fmt.Println(err.Error())
	if err == nil {
		t.Fatalf("checking empty failed")
	}
}
