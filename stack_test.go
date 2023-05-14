package stack

import (
	"fmt"
	"testing"

	"github.com/edikrisnayana/data-structure/stack"
)

func TestPushSecondPushSuccess(t *testing.T) {
	stack := stack.NewStack()
	err := stack.Push(1)
	if err != nil {
		t.Fatalf(err.Error())
	}
	err = stack.Push(2)
	if err != nil {
		t.Fatalf(err.Error())
	}
}

func TestPushSecondPushMismacthType(t *testing.T) {
	stack := stack.NewStack()
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

func TestPopSuccess(t *testing.T) {
	stack := stack.NewStack()
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

func TestPopEmptyStack(t *testing.T) {
	stack := stack.NewStack()
	_, err := stack.Pop()
	fmt.Println(err.Error())
	if err == nil {
		t.Fatalf("checking empty failed")
	}
}
