import util
from copy import deepcopy
import random
def is_solvable(flow_game_state):
    '''
    determines if it possible to solve the game from every color's perspective
    '''
    current_positions = flow_game_state.get_curr_positions()
    goal_positions = flow_game_state.get_goal_positions()
    for i in range(len(current_positions)):
        state = deepcopy(flow_game_state)
        frontier = [current_positions[i]]
        while frontier:
            pos = frontier.pop(0)
            for new_pos in state.legal_moves(pos):
                state.move(i, new_pos)
                frontier.append(new_pos)
        if state.get_board()[goal_positions[i][0]][goal_positions[i][1]] != chr(i + ord('a')):
            return False
    return True


#this astar seach is really just a greedy search because there is no backwards cost
def aStarSearch(problem, Print=False):
    '''
    regular astar search which expands nodes based on heuristic and backwards cost
    '''
    i = 0
    explored = []
    frontier = util.PriorityQueue()
    frontier.push(problem.get_start_state(), problem.heuristic(problem.get_start_state()))
    while not frontier.isEmpty():
        state = frontier.pop()
        #print(state)
        if state.get_board() not in explored:
            i += 1
            if problem.goal_test(state):
                if Print:
                    print(i)
                return state
            explored.append(state.get_board)
            possible_actions = problem.get_actions(state)
            for action in possible_actions:
                new_state = problem.get_result(state, action)
                frontier.push(new_state, problem.heuristic(new_state))

def flow_game_search(problem, Print=False):
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
                possible_actions = problem.get_actions(state)
                for action in possible_actions:
                    new_state = problem.get_result(state, action)
                    frontier.push(new_state, problem.heuristic(new_state))

def flow_game_search_v1(problem, Print=False):
    '''
    greedy variation which does not expand unsolvable nodes
    v1 means it solves one color a time
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
                color, possible_actions = problem.get_alt_actions(state)
                for action in possible_actions:
                    new_state = problem.get_alt_result(state, action, color)
                    frontier.push(new_state, problem.heuristic(new_state))

def breadthFirstSearch(problem):
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
            possible_actions = problem.get_actions(state)
            for action in possible_actions:
                new_state = problem.get_result(state, action)
                frontier.append(new_state)
#############################################################################################
#iteration 1 will definitely have some flaws to it lol
#should these methods be in flow game search??
    #constraints for csp
        #1) each node must be connected a node of the same color next to it if its not a start node
        #2) all start nodes must be connected by an unbroken path to the goal node
    
    #these are for min conflicts algorithm
        #if this isnt good will try to implement a backtracking/MAC algorithm 

def check_constraints(self, flow_game_state):
    conflicted_variables = []
    num_constraints = 0 #number of violated constraint
    start_positions = flow_game_state.get_start_positions()
    goal_positions = flow_game_state.get_goal_positions()

    for i in range(flow_game_state.get_gridsize()):
        for j in range(flow_game_state.get_gridsize()):
            if (i, j) in start_positions:
                goal_pos = goal_positions[start_positions.index((i,j))]
                if not self.connected_goal(flow_game_state, (i, j), goal_pos):
                    num_constraints += 1
            elif (i, j) not in goal_positions:
                if flow_game_state.get_board()[i][j] and not self.is_connected(flow_game_state, (i,j)):
                    num_constraints += 1
                    conflicted_variables.append((i,j))

    if not num_constraints or not conflicted_variables:
        return num_constraints, None
    else:
        print(conflicted_variables)
        return num_constraints, random.choice(conflicted_variables) 

def is_connected(self, flow_game_state, pos):
    for adj_pos in flow_game_state.adjacent_positions(pos):
        if flow_game_state.get_board()[pos[0]][pos[1]] == flow_game_state.get_board()[adj_pos[0]][adj_pos[1]]:
            return True
    return False

def connected_goal(self, flow_game_state, start_pos, goal_pos):
    goal_val = flow_game_state.get_board()[start_pos[0]][start_pos[1]]
    explored = []

    def modified_dfs(current_position):

        adjacent_positions = flow_game_state.adjacent_positions(current_position)
        for pos in adjacent_positions:
            if flow_game_state.get_board()[pos[0]][pos[1]] == goal_val and pos not in explored:
                explored.append(pos)
                modified_dfs(pos)

    modified_dfs(start_pos)
    return goal_pos in explored

def random_assignment(flow_game_state):
    for i in range(flow_game_state.get_gridsize()):
        for j in range(flow_game_state.get_gridsize()):
            if (i, j) in flow_game_state.get_goal_positions():
                flow_game_state.move(flow_game_state.get_goal_positions().index((i,j)), (i,j))
            elif not flow_game_state.get_board()[i][j]:
                flow_game_state.move(random.choice(range(flow_game_state.get_gridsize())), (i,j))
    return flow_game_state


def min_conflicts_search(problem):
    flow_game_state = random_assignment(problem.get_start_state())
    print(flow_game_state.get_start_positions())
    print(flow_game_state)
    num_conflicts, next_var = problem.check_constraints(flow_game_state)
    while num_conflicts > 0:

        print(num_conflicts)
        print(next_var)
        print(flow_game_state)

        if not next_var:
            return flow_game_state

        fewest_violations = float('inf')
        curr_move = None
        next_next_var = None

        for i in range(flow_game_state.get_gridsize()):

            state = deepcopy(flow_game_state)
            state.move(i, next_var)

            violations, next = problem.check_constraints(state)
            if violations < fewest_violations:
                curr_move = i 
                next_next_var = next
                fewest_violations = violations

        num_conflicts = fewest_violations
        if num_conflicts > 0:
            flow_game_state.move(curr_move, next_var)
            next_var = next_next_var


    return flow_game_state
    #first do a random assignment of the board
