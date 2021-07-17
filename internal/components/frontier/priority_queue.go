package frontier

import "container/heap"

type PriorityQueue interface {
	Len() int
	Insert(*pqItem)
	Get() *pqItem
}

type pqItem struct {
	value    interface{}
	priority int
	index    int
}

type priorityQueue []*pqItem

func NewPriorityQueue() *priorityQueue {
	pq := make(priorityQueue, 0)
	heap.Init(&pq)
	return &pq
}

func (pq *priorityQueue) Insert(pqItem *pqItem) {
	heap.Push(pq, pqItem)
}

func (pq *priorityQueue) Get() *pqItem {
	if pq.Len() == 0 {
		return nil
	}
	if item, ok := heap.Pop(pq).(*pqItem); ok {
		return item
	}
	return nil
}

func (pq priorityQueue) Len() int {
	return len(pq)
}

func (pq priorityQueue) Less(i, j int) bool {
	return pq[i].priority > pq[j].priority
}

func (pq priorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *priorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*pqItem)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *priorityQueue) Pop() interface{} {
	if pq.Len() == 0 {
		return nil
	}
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[0 : n-1]
	return item
}
