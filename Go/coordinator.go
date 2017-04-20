package main

import (
	"container/heap"
	"log"
	"sync"
)

type coordinator struct {
	numThreads int
	isSolved   bool
	solution   state
	explored   []string
	frontier   PriorityQueue
	lock       sync.Locker
	waitGroup  *sync.WaitGroup
	//cond            *sync.Cond
}

// TODO(tanay) should this check if initialState is solved?
func NewCoordinator(initialState state) *coordinator {
	frontier := make(PriorityQueue, 0)
	frontier = append(frontier, &StateWrapper{
		State: initialState,
		index: 0,
	})
	heap.Init(&frontier)
	return &coordinator{
		isSolved:  false,
		explored:  make([]string, 0),
		frontier:  frontier,
		lock:      &sync.Mutex{},
		waitGroup: &sync.WaitGroup{},
		//cond:            sync.NewCond(&sync.Mutex{}),
		numThreads: 10, //TODO(tanay) this could be configurable
	}
}

// the solve function will spawn threads that will work together to solve
// the flow instance based on the chosen transition model and will return the
// final, solution state
func (c *coordinator) Solve() state {
	for i := 0; i < c.numThreads; i++ {
		c.waitGroup.Add(1)
		go c.helper()
	}

	c.waitGroup.Wait()
	return c.solution
}

// TODO(tanay) do i need to deal with the case when all colors are "solved" but
// not all spaces are filled (dont think so explicitly, will just keep searching tree
// TODO(tanay) behavior when you have a non-solvable game board
// if there is nothing in the frontier. Then, wake the thread up
// once something is placed back onto the frontier
func (c *coordinator) helper() {
	for !c.isSolved {
		var s state

		// first critical section
		// pop the first state from the frontier
		c.lock.Lock()
		if len(c.frontier) > 0 {
			stateWrapper := heap.Pop(&c.frontier)
			s = stateWrapper.(*StateWrapper).State
			c.explored = append(c.explored, s.Serialize())
		}
		c.lock.Unlock()

		//check if s is nil, if it is, we should sleep this thread
		//TODO(tanay) not sure if this will cause deadlock/is inappopriate usage
		if s == nil {
			//c.cond.L.Lock()
			//c.cond.Wait()
			//c.cond.L.Unlock()
			continue
		}

		// do work
		// check if states are solvable/solved here
		// if not solveable remove from nextStates
		log.Print(s)
		isSolved := false
		var solvedState state
		nextStates := s.NextStates()
		filteredNextStates := nextStates[:0]
		for _, ns := range nextStates {
			if ns.IsSatisfied() {
				isSolved = true
				solvedState = ns
				break
			}
			if ns.IsSatisfiable() {
				filteredNextStates = append(filteredNextStates, ns)
			}
		}

		// last critical section
		c.lock.Lock()
		if isSolved {
			c.solution = solvedState
			c.isSolved = true
		} else {
			for _, ns := range filteredNextStates {
				serializedState := ns.Serialize()
				unique := true
				for _, exploredSerializedState := range c.explored {
					if exploredSerializedState == serializedState {
						unique = false
						break
					}
				}
				if unique {
					heap.Push(&c.frontier, &StateWrapper{State: ns})
					//c.cond.Signal()
				}
			}
		}
		c.lock.Unlock()
	}

	c.waitGroup.Done()
}
