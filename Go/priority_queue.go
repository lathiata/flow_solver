package main

// This is taken from https://golang.org/pkg/container/heap/#Fix
// Once the interface is completed (Len, Less, Swap, Push, Pop), we can use
// container/heap methods

type StateWrapper struct {
	State state
	index int
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*StateWrapper

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the LOWEST, not HIGHEST, priority so we use < than here.
	return pq[i].State.Heuristic() < pq[j].State.Heuristic()
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*StateWrapper)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}
