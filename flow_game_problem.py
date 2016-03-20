from flow_game_state import FlowGameState
from copy import copy, deepcopy
import random

class FlowGameProblem:
	def __init__(self, gridsize, start_positions):
		self.start_state = FlowGameState(gridsize, start_positions)

	def get_start_state(self):
		return self.start_state 

	# change how this is designed: have the ending positions initialized from the start
	# in the get_actions functions, instead of checking if the curr position is at the goal position,
	# simply check if the goal position would be a legal next move (use reached_goal(color/index) from
	# the flow_game_state)
	# ^ DONE ^

	def goal_test(self, flow_game_state):
		for i in range(len(flow_game_state.get_curr_positions())):
			if not flow_game_state.reached_goal(i):
				return False
		return flow_game_state.full_grid()

	#a list of dictionaries which represent all of the sets of all possible moves 
	def get_actions_v0(self, flow_game_state):
		all_actions = []

		def action_helper(actions):
			if len(actions.keys()) == len(flow_game_state.get_curr_positions()):
				all_actions.append(actions)
			else:
				i = len(actions.keys())
				curr_pos = flow_game_state.get_curr_positions()[i]
				if not flow_game_state.reached_goal(i):
					for move in flow_game_state.legal_moves(curr_pos):
						if move not in actions.values():
							new_actions = copy(actions)
							new_actions[i] = move
							action_helper(new_actions)
				else:
					new_actions = copy(actions)
					new_actions[i] = None
					action_helper(new_actions)

		action_helper({})
		return [dict(t) for t in set([tuple(d.items()) for d in all_actions])]

	def get_result_v0(self, flow_game_state, action):
		next_fgs = deepcopy(flow_game_state)
		for i in action.keys():
			if action[i]:
				next_fgs.move(i, action[i])
		return next_fgs

	#move 1 color 1 square at a time
	def get_actions_v1(self, flow_game_state):
		curr_positions = flow_game_state.get_curr_positions()
		for i in range(len(curr_positions)):
			if not flow_game_state.reached_goal(i):
				return i, flow_game_state.legal_moves(curr_positions[i])

	def get_result_v1(self, flow_game_state, action, i):
		next_fgs = deepcopy(flow_game_state)
		next_fgs.move(i, action)
		return next_fgs

	#need to update this ish
	def heuristic(self, flow_game_state):
		curr_positions = flow_game_state.get_curr_positions()
		goal_positions = flow_game_state.get_goal_positions()
		max_distance = -float('inf')
		for i in range(len(curr_positions)):
			dist = abs(curr_positions[i][0] - goal_positions[i][0]) + abs(curr_positions[i][1] - goal_positions[i][1])
			max_distance = max(max_distance, dist)

		#try giving preference to the edges
		if flow_game_state.border(curr_positions[i]):
			max_distance /= 4
		
		return max_distance








# flow game as a constraint satisfaction problem

# all positions must be colored
# for all non start/end positions on the board, there must be an adjacent position of the same color 
# for start positions there must be an uninterrupted connection to the goal position
# each position in the grid will have a domain 



# beam search?


