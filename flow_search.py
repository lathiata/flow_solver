import util
from copy import deepcopy
import random
import flow_game_constants

#TODO  make this check faster by having things in the state constantly updating-especially for connectedness.
#      also this function should probably be in flow_game_problem because the abstraction is that the search
#      doesn't really know anything about how the problem works (can work for any puzzle)
def is_satisfied(state, gridsize):
    '''
    Determines if the board is in a solved state
    '''
    def has_adjacent(pos, board):
        val = state.get_pos_val(pos)
        for adj_pos in state.adjacent_positions(pos):
            if state.get_pos_val(adj_pos) == val:
                return True
        return False

    def connected(board, i):
        explored = []
        def connected_helper(curr, goal, val):
            if curr == goal:
                return True

            next_in_path = [pos for pos in state.adjacent_positions(curr) \
                            if pos not in explored and state.get_pos_val(pos) == val]
            dfs_values = []
            for pos in next_in_path:
                explored.append(pos)
                dfs_values.append(connected_helper(pos, goal, val))

            if not dfs_values:
                return False #reached end of path and nothing

            return reduce(lambda x, y: x or y, dfs_values)

        curr = state.get_start_positions()[i]
        goal = state.get_goal_positions()[i]

        return connected_helper(curr, goal, i)



    board = state.get_board()
    is_filled = reduce(lambda x, y: x and y, [state.get_val(i,j) != flow_game_constants.EMPTY for i in range(gridsize) for j in range(gridsize)])
    is_adjacent = reduce(lambda x, y: x and y, [has_adjacent((i,j), board) for i in range(gridsize) for j in range(gridsize)])
    is_connected = reduce(lambda x, y: x and y, [connected(board, i) for i in range(state.get_num_colors())])
    return is_filled and is_adjacent and is_connected

#TODO move to flow_game_problem
def is_solvable(flow_game_state):
    '''
    Determines if it possible to solve the game from every color's perspective.
    See the comment for flow_game_state.solvable_helper for implementation details
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

def flow_game_search(problem, version, Print=False):
    """
    Greedy search (Breadth First Search with a heuristic).
    Does not expand unsolvable states (brute force check using is_solvable).
    """
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
                print(state)
            if problem.goal_test(state):
                print("done " + str(i))
                return state
            explored.append(state.get_board)
            if is_solvable(state):
                problem.explore(version, state, frontier)

#TODO have an algorithm that checks if i can meet all constraints still
#     still running into the problem of making the constraints easy to check (connectedness)
def csp_search(problem):
    """
    Poses the flow game as a CSP

    Maintain Arc Consistency
    Forward check
    Backwards check

    forall (i,j) :
        Xi,j != flow_game_constants.EMPTY

    forall (i,j) not in starting positions, there exists adjacent (p,q) and (a,b) :
        Xi,j == Xp,q
        Xi,j == Xa,b

    forall (i,j) in starting positions, there exists adjacent (p,q) :
        Xi,j == Xp,q

    forall (i,j) in starting positions, there exists a path to the corresponding ending position (p,q)
    """
    def inferences(state, domains):
        pass

    def is_consistent(state, pos, val):
        pass

    #recursively perform backtrack search (DFS w/ improvement)
    #foward checking
    #maintain arc consistency
    def backtrack_search(state, domains, gridsize):
        if is_satisfied(state, gridsize):
            return state
        else:
            next_pos = min(domains, len(key=domains.get)) #select the variable with the fewest things in domain
            for val in domains[next_pos]:
                if is_consistent(state, next_pos, val):
                    state_copy = deepcopy(state)
                    state_copy.move(val, next_pos) #add the position to the state
                    inference_success, inference = inferences(state_copy, domains)

                    if inference_success:
                        #TODO add inference to {state, domains}
                        result = backtrack_search(state_copy, domains, gridsize)

                        if result:
                            return result

            return None #This path is a failure

    state = problem.get_start_state()
    start_positions = state.get_start_positions()
    goal_positions = state.get_goal_positions()
    gridsize = state.get_gridsize()

    #create mapping (i,j) -> (possible values)
    root_domains = {(i,j) : range(len(state.get_start_positions())) \
               if (i,j) not in start_positions and (i,j) not in goal_positions \
               else state.get_val(i,j) \
               for i in range(gridsize) \
               for j in range(gridsize) \
               }

    #run backtrack algorithm
    return backtrack_search(state, root_domains, gridsize)
