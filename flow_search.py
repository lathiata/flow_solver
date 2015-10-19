import util
from copy import deepcopy

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
    an astar variation which does not expand unsolvable nodes
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
