package asyncqueue

import (
	"container/list"
	"sync"
)

type Queue interface {
	Push(interface{})
	BatchPush([]interface{})
	BatchPop(n int) []interface{}
	Pop() interface{}
}

type queue struct {
	*sync.Mutex
	*sync.Cond
	*list.List
}

func NewQueue() Queue {
	mut := &sync.Mutex{}
	return &queue{
		Mutex: mut,
		Cond:  sync.NewCond(mut),
		List:  list.New(),
	}
}

func (q *queue) Push(item interface{}) {
	q.Lock()
	q.PushBack(item)
	q.Unlock()

	q.Signal()
}

func (q *queue) BatchPush(items []interface{}) {
	q.Lock()

	for _, item := range items {
		q.PushBack(item)

	}
	q.Unlock()

	q.Signal()
}

func (q *queue) Pop() interface{} {
	q.Lock()
	defer q.Unlock()

	for q.Len() == 0 {
		q.Wait()
	}

	return q.Remove(q.Front())
}

func (q *queue) BatchPop(n int) []interface{} {
	q.Lock()
	defer q.Unlock()

	for q.Len() == 0 {
		q.Wait()
	}

	l := q.Len()

	if l < n {
		n = l
	}

	items := make([]interface{}, n)

	for i := 0; i < n; i++ {

		items[i] = q.Remove(q.Front())
	}

	return items
}
