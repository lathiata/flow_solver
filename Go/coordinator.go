package main

import (
	"container/heap"
	//"log"
	"sync"
)

type coordinator struct {
	numThreads     int
	isSolved       bool
	solution       state
	explored       []string
	frontier       PriorityQueue
	cond           *sync.Cond
	waitGroup      *sync.WaitGroup
}

// TODO(tanay) push logging into custom handler to easily enable/disable logging
// TODO(tanay) should this check if initialState is solved?
func NewCoordinator(initialState state, numThreads int) *coordinator {
	frontier := make(PriorityQueue, 0)
	frontier = append(frontier, &StateWrapper{
		State: initialState,
		index: 0,
	})
	heap.Init(&frontier)
	return &coordinator{
		isSolved:       false,
		explored:       make([]string, 0),
		frontier:       frontier,
		waitGroup:      &sync.WaitGroup{},
		cond:           sync.NewCond(&sync.Mutex{}),
		numThreads:     numThreads,
	}
}

// the solve function will spawn threads that will work together to solve
// the flow instance based on the chosen transition model and will return the
// final, solution state
func (c *coordinator) Solve() state {
	for i := 0; i < c.numThreads; i++ {
		c.waitGroup.Add(1)
		go c.helper(i)
	}

	c.waitGroup.Wait()
	return c.solution
}

// TODO(tanay) behavior when you have a non-solvable game board
// if there is nothing in the frontier. Then, wake the thread up
// once something is placed back onto the frontier
func (c *coordinator) helper(id int) {
	defer c.waitGroup.Done()
	for {
		// first critical section
		// pop the first state from the frontier
		c.cond.L.Lock()
		for c.frontier.Len() == 0 && !c.isSolved {
			//log.Printf("[%d] sleeping until frontier has something", id)
			// Sleep until there is something on the frontier
			c.cond.Wait()
		}

		// When we wake up, it is because there is a new state to explore OR
		// another goroutine solved the board
		if c.isSolved {
			//log.Printf("[%d] Another thread solved board, return", id)
			c.cond.L.Unlock()
			return
		}

		// This goroutine will have the lock when it resumes
		stateWrapper := heap.Pop(&c.frontier)
		// TODO(tanay): Could check the val here and panic if its negative?
		s := stateWrapper.(*StateWrapper).State
		//log.Printf("[%d] exploring %v", id, s)
		c.explored = append(c.explored, s.Serialize())
		c.cond.L.Unlock()

		// do work:
		// 	check if states are solvable/solved here
		// 	if not solveable remove from nextStates
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
		c.cond.L.Lock()
		if isSolved {
			c.solution = solvedState
			c.isSolved = true
			//log.Printf("[%d] solved the board", id)
			c.cond.Broadcast()
		} else {
			pushed := false
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
					pushed = true
				}
			}

			if pushed {
				c.cond.Broadcast()
			}
		}
		c.cond.L.Unlock()
	}
}
