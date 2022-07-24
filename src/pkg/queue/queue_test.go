package queue_test

import (
	"fmt"
	"testing"

	"github.com/wada1001/algorithm/src/pkg/queue"
)

func TestQueue(t *testing.T) {
	
	type args struct {
		q *queue.Queue
		val int
	}
	tests := []struct {
		name string
		args args
	}{
		{"success", args{queue.MakeQueue(10), 20}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.args.q.Enqueue(test.args.val)
			val, err := test.args.q.Dequeue()
			if err != nil {
				t.Errorf("error occured")
			}
			if val != test.args.val {
				t.Errorf(fmt.Sprintf("worng value want = %d, got = %d" , test.args.val, val))
			}
		})
	}
}

func TestQueue2(t *testing.T) {
	
	type args struct {
		q *queue.QueueLinkList
		val int
	}
	tests := []struct {
		name string
		args args
	}{
		{"success", args{queue.MakeQueueLinkList(10), 20}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.args.q.Enqueue(test.args.val)
			val, err := test.args.q.Dequeue()
			if err != nil {
				t.Errorf("error occured")
			}
			if val != test.args.val {
				t.Errorf(fmt.Sprintf("worng value want = %d, got = %d" , test.args.val, val))
			}
		})
	}
}