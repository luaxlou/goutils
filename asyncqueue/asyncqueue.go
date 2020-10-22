package asyncqueue

import (
	"container/list"
	"sync"
)

type Queue struct {
	*sync.Mutex
	*sync.Cond
	*list.List
}

func New() *Queue {
	mut := &sync.Mutex{}
	return &Queue{
		Mutex: mut,
		Cond:  sync.NewCond(mut),
		List:  list.New(),
	}
}

func (q *Queue) Push(item interface{}) {
	q.Lock()
	q.PushBack(item)
	q.Unlock()

	q.Signal()
}

func (q *Queue) BatchPush(items []interface{}) {
	q.Lock()

	for _, item := range items {
		q.PushBack(item)

	}
	q.Unlock()

	q.Signal()
}


func (q *Queue) BatchPushs(items ...interface{}) {
	q.Lock()

	for _, item := range items {
		q.PushBack(item)

	}
	q.Unlock()

	q.Signal()
}

func (q *Queue) Pop() interface{} {
	q.Lock()
	defer q.Unlock()

	for q.Len() == 0 {
		q.Wait()
	}

	return q.Remove(q.Front())
}

func (q *Queue) BatchPop(n int) []interface{} {
	q.Lock()
	defer q.Unlock()
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

func (q *Queue) PopAll() []interface{} {
	q.Lock()
	defer q.Unlock()
	l := q.Len()

	return q.BatchPop(l)
}
