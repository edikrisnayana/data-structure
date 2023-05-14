package queue

import (
	"fmt"
	"reflect"

	"github.com/edikrisnayana/data-structure/intf"
)

type queue struct {
	itemType reflect.Type
	items    []intf.T
}

func Init() intf.Queue {
	var queue intf.Queue = new(queue)
	return queue
}

func (q *queue) IsEmpty() bool {
	return q.Size() == 0
}

func (q *queue) Offer(item intf.T) error {
	if q.itemType == nil {
		q.itemType = reflect.ValueOf(item).Type()
	}

	if q.itemType != nil && q.itemType != reflect.ValueOf(item).Type() {
		return fmt.Errorf("item type is not consistent. Should use %v", q.itemType)
	}
	q.items = append(q.items, item)
	return nil
}

func (q *queue) Peek() (intf.T, error) {
	if q.IsEmpty() {
		return nil, fmt.Errorf("queue is empty")
	}

	return q.items[0], nil
}

func (q *queue) Poll() (intf.T, error) {
	if q.IsEmpty() {
		return nil, fmt.Errorf("queue is empty")
	}

	item := q.items[0]
	q.items[0] = nil
	q.items = q.items[1:]
	return item, nil
}

func (q *queue) Size() int {
	return len(q.items)
}
