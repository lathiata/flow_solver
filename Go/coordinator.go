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
	cond            sync.Cond
}

// TODO(tanay) should this check if initialState is solved?
func NewCoordinator(initialState state) *coordinator {
	cL := &sync.Mutex{}
	return &coordinator{
		transitionModel: &TransitionModelImplementation{},
		isSolved:        false,
		explored:        make([]state, 0),
		frontier:        []state{initialState},
		lock:            &sync.Mutex{},
		waitGroup:       &sync.WaitGroup{},
		cond:            sync.NewCond(&sync.Mutex{}),
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

// TODO(tanay) I should use a cond variable to sleep the thread
// if there is nothing in the frontier. Then, wake the thread up
// once something is placed back onto the frontier
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

		//check if s is nil, if it is, we should sleep this thread
		//TODO(tanay) not sure if this will cause deadlock/is inappopriate usage
		if s == nil {
			c.cond.L.Lock()
			c.cond.Wait()
			c.cond.L.Unlock()
			continue
		}

		// do work
		// check if states are solvable/solved here
		// if not solveable remove from nextStates
		isSolved := false
		var solvedState state
		nextStates := c.transitionModel.Explore(s)
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
		// if we solved, update isSolved
		// else
		//for ns := range nextStates {
		// check if ns in explored, else add to frontier
		//}
		if isSolved {
			c.solution = solvedState
			c.isSolved = true
		} else {
			for _, ns := range filteredNextStates {
				unique := true
				for s := range c.explored {
					if ns.Equals(s) {
						unique = false
						break
					}
				}
				if unique {
					c.frontier = append(c.frontier, ns)
					c.cond.Signal()
				}
			}
		}
		c.lock.Unlock()
	}

	c.waitGroup.Done()
}
