package frontier

import "crawler/internal/components"

type frontier struct {
	priorityQueue PriorityQueue
}

func NewFrontier() *frontier {
	return &frontier{
		priorityQueue: NewPriorityQueue(),
	}
}

func (f *frontier) Push(item *components.Item) {
	f.priorityQueue.Insert(&pqItem{
		value:    item.Value,
		priority: item.Priority,
	})
}

func (f *frontier) Pop() *components.Item {
	pqItem := f.priorityQueue.Get()
	if pqItem != nil {
		return &components.Item{
			Value:    pqItem.value,
			Priority: pqItem.priority,
		}
	}
	return nil
}

func (f *frontier) Len() int {
	return f.priorityQueue.Len()
}
