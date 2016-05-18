import util
from copy import deepcopy
import random

def is_solvable(flow_game_state):
    '''
    determines if it possible to solve the game from every color's perspective
    '''
    current_positions = flow_game_state.get_curr_positions()

    for i in range(len(current_positions)):
        state = deepcopy(flow_game_state)
        frontier = [current_positions[i]]

        while frontier:
            pos = frontier.pop(0)
            for new_pos in state.legal_moves(pos):
                state.move(i, new_pos)
                frontier.append(new_pos)

        if not state.solvable_helper(i):
            return False
            
    return True

#this astar seach is really just a greedy search because there is no backwards cost
def aStarSearch(problem, version, Print=False):
    '''
    regular astar search which expands nodes based on heuristic and backwards cost
    '''
    i = 0
    explored = []
    frontier = util.PriorityQueue()
    frontier.push(problem.get_start_state(), problem.heuristic(problem.get_start_state()))
    while not frontier.isEmpty():
        state = frontier.pop()
        if state.get_board() not in explored:
            i += 1
            if Print:
                print(i)
            if problem.goal_test(state):
                print("done " + str(i))
                return state
            explored.append(state.get_board())
            problem.explore(version, state, frontier)

def flow_game_search(problem, version, Print=False):
    '''
    greedy variation which does not expand unsolvable nodes
    a move here moves all the colors by one
    '''
    i = 0
    explored = []
    frontier = util.PriorityQueue()
    frontier.push(problem.get_start_state(), problem.heuristic(problem.get_start_state()))
    while not frontier.isEmpty():
        state = frontier.pop()
        if state.get_board() not in explored:
            i += 1
            if Print:
                print(i)
            if problem.goal_test(state):
                print("done " + str(i))
                return state
            explored.append(state.get_board)
            if is_solvable(state):
                problem.explore(version, state, frontier)

def breadthFirstSearch(problem, version):
    """
    Search the shallowest nodes in the search tree first.

    """
    explored = []
    frontier = [problem.get_start_state()]
    while len(frontier) != 0:
        state = frontier.pop(0)
        #print(state)
        if state.get_board() not in explored:
            if problem.goal_test(state):
                return state
            explored.append(state.get_board())
            problem.explore(version, state, frontier)

def cspSearch(problem):
    """
    Poses the flow game as a CSP

    Maintain Arc Consistency
    Forward check
    Backwards check

    Xi,j != -1 forall i,j 
    Xi,j == Xp,q for at least 1 adjacent 
    """
    def is_satisfied(state, gridsize):
        pass

    def inference(state, domains):
        pass

    def is_consistent(state, pos, val):
        pass
        
    def inference_check(domain):
        for val in domain.values():
            if len(val) <= 0:
                return False

    #recursively perform backtrack search (DFS w/ improvement)
    #foward checking
    #maintain arc consistency
    def backtrack_search(state, domains, gridsize):
        if is_satisfied(state, gridsize):
            return state
        else:
            next_pos = min(domains, len(key=domains.get)) #select the variable with the fewest things in domain
            for val in domains[next_pos]:
                if is_consistent(state, next_pos, val)
            return None #if we hit here, there is a failure (should not happen)

    state = problem.get_start_state()
    start_positions = state.get_start_positions()
    goal_positions = state.get_goal_positions()
    gridsize = state.get_gridsize()

    #create mapping (i,j) -> (possible values)
    root_domains = {(i,j) : range(len(state.get_start_positions())) \
               for i in range(gridsize) \
               for j in range(gridsize) \
               if (i,j) not in start_positions and (i,j) not in goal_positions}

    #create mapping (i,j) -> assigned value
    # root_assignment = {(i,j) : state.get_val(i,j) for (i,j) in list(set(start_positions).union(goal_positions))}

    #run backtrack algorithm
    return backtrack_search(state, root_domains, gridsize)

