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
        if state.get_board() not in explored:
            i += 1
            if Print:
                print(i)
            if problem.goal_test(state):
                print("done " + str(i))
                return state
            explored.append(state.get_board())
            possible_actions = problem.get_actions_v0(state)
            for action in possible_actions:
                new_state = problem.get_result_v0(state, action)
                frontier.push(new_state, problem.heuristic(new_state))

def aStarSearchV1(problem, Print=False):
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
            color, possible_actions = problem.get_actions_v1(state)
            for action in possible_actions:
                new_state = problem.get_result_v1(state, action, color)
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
                possible_actions = problem.get_actions_v0(state)
                for action in possible_actions:
                    new_state = problem.get_result_v0(state, action)
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
                color, possible_actions = problem.get_actions_v1(state)
                for action in possible_actions:
                    new_state = problem.get_result_v1(state, action, color)
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
            possible_actions = problem.get_actions_v0(state)
            for action in possible_actions:
                new_state = problem.get_result_v0(state, action)
                frontier.append(new_state)

def cspSearch(problem):
    """
    Poses the flow game as a CSP
    """
    print problem.get_start_state()

