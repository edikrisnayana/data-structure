package test

import (
	"fmt"
	"testing"

	"github.com/edikrisnayana/data-structure/queue"
)

func TestQueueOfferSecondOfferSuccess(t *testing.T) {
	queue := queue.Init()
	err := queue.Offer(1)
	if err != nil {
		t.Fatalf(err.Error())
	}
	err = queue.Offer(2)
	if err != nil {
		t.Fatalf(err.Error())
	}
}

func TestQueueOfferSecondOfferMismatchType(t *testing.T) {
	queue := queue.Init()
	err := queue.Offer(1)
	if err != nil {
		t.Fatalf(err.Error())
	}
	err = queue.Offer("2")
	fmt.Println(err.Error())
	if err == nil {
		t.Fatalf("checking type failed. should give error mismatch type")
	}
}

func TestQueuePollSuccess(t *testing.T) {
	queue := queue.Init()
	queue.Offer(1)
	queue.Offer(2)
	currSize := queue.Size()
	_, err := queue.Poll()
	if err != nil {
		t.Fatalf(err.Error())
	}

	if currSize == queue.Size() {
		t.Fatalf("stack pop failed")
	}
}

func TestQueuePollEmptyStack(t *testing.T) {
	queue := queue.Init()
	_, err := queue.Poll()
	fmt.Println(err.Error())
	if err == nil {
		t.Fatalf("checking empty failed")
	}
}
