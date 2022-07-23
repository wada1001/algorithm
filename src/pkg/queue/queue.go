package queue

import (
	"errors"

	"github.com/wada1001/algorithm/src/pkg/list"
)

type Queue struct {
	arr []int
	first int
	last int
}

func MakeQueue(length int) *Queue {
	return &Queue{
		arr: make([]int, length),
		first: 0,
		last: 0,
	}
}

func (q *Queue) Enqueue(val int) error {
	q.arr[q.last] = val
	if (q.last + 1) % len(q.arr) == q.first  {
		return errors.New("queue is full")
	}

	q.last = (q.last + 1) % len(q.arr)
	return nil
}

func (q *Queue) Dequeue() (int, error) {
	if q.first == q.last {
		return 0, errors.New("queue is empty")
	}

	v := q.arr[q.first]
	q.first = (q.first + 1) % len(q.arr)
	return v, nil
}


type QueueLinkList struct {
	first *list.LinkListNode
	last *list.LinkListNode
}

func MakeQueueLinkList(length int) *QueueLinkList {
	return &QueueLinkList{
		first: nil,
		last: nil,
	}
}

func (q *QueueLinkList) Enqueue(val int) {
	node := list.MakeLinkListNode(nil, val)
	if q.first == nil && q.last == nil {
		q.first = node
		q.last = node
	} else {
		q.last.Next = node
		q.last = node
	}
}

func (q *QueueLinkList) Dequeue() (int, error) {
	if q.first == nil && q.last == nil {
		return 0, errors.New("queue is empty")
	}

	if q.first == q.last {
		// same instance
		v := q.first.Val
		q.first = nil
		q.last = nil
		return v, nil
	} 

	v := q.first.Val
	q.first = q.first.Next
	return v, nil
}

