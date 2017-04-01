package main

import (
	"sync"
)

// TODO(tanay) implement threaded search algorithms

// TODO(tanay) determine the best to pass in
// transition model to coordinator
type coordinator struct {
	transitionModel TransitionModel
	numThreads      int
	isSolved        bool
	solution        state
	explored        []state
	frontier        []state
	lock            sync.Locker
	waitGroup       sync.WaitGroup
}

// TODO(tanay) should this check if initialState is solved?
func NewCoordinator(initialState state) *coordinator {
	return &coordinator{
		transitionModel: &TransitionModelImplementation{},
		isSolved:        false,
		explored:        make([]state, 0),
		frontier:        []state{initialState},
		lock:            &sync.Mutex{},
		waitGroup:       &sync.WaitGroup{},
		numThreads:      10, //TODO(tanay) this could be configurable
	}
}

// the solve function will spawn
// threads that will work together to solve
// the flow instance based on the chosen
// transition model and will return the
// final, solution state
func (c *coordinator) solve() state {
	for i := 0; i < c.numThreads; i++ {
		c.waitGroup.Add(1)
		go c.helper()
	}

	c.waitGroup.Wait()
	return c.solution
}

func (c *coordinator) helper() {
	for !c.isSolved {
		var s state
		// first critical section
		c.lock.Lock()
		if len(c.frontier) > 0 {
			s = c.frontier[0]
			c.frontier = c.frontier[1:]
		}
		c.lock.Unlock()

		// do work
		nextStates := c.transitionModel.Explore(s)
		// check if states are solvable/solved here
		// if not solveable remove from nextStates
		//

		// last critical section
		c.lock.Lock()
		// if we solved, update isSolved
		// else
		//for ns := range nextStates {
		// check if ns in explored, else add to frontier
		//}
		c.lock.Unlock()

	}
	c.Done()
}
