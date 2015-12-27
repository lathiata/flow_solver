from flow_game_state import FlowGameState
from copy import copy, deepcopy
import random

class FlowGameProblem:
	def __init__(self, gridsize, start_positions):
		self.start_state = FlowGameState(gridsize, start_positions)
	
	def get_start_state(self):
		return self.start_state 

	def goal_test(self, flow_game_state):
		curr_positions = flow_game_state.get_curr_positions()
		goal_positions = flow_game_state.get_goal_positions()
		for i in range(len(curr_positions)):
			if not flow_game_state.is_goal(curr_positions[i], i):
				return False
		return flow_game_state.full_grid() and True

	def get_actions(self, flow_game_state):
		all_actions = []
		def action_helper(actions):
			if len(actions.keys()) == len(flow_game_state.get_curr_positions()):
				all_actions.append(actions)
			else:
				i = len(actions.keys())
				curr_pos = flow_game_state.get_curr_positions()[i]
				if not flow_game_state.is_goal(curr_pos, i):
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

	def get_result(self, flow_game_state, action):
		next_fgs = deepcopy(flow_game_state)
		for i in action.keys():
			if action[i]:
				next_fgs.move(i, action[i])
		return next_fgs

	def get_alt_actions(self, flow_game_state):
		curr_positions = flow_game_state.get_curr_positions()
		goal_positions = flow_game_state.get_goal_positions
		for i in range(len(curr_positions)):
			if not flow_game_state.is_goal(curr_positions[i], i):
				return i, flow_game_state.legal_moves(curr_positions[i])

	def get_alt_result(self, flow_game_state, action, i):
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
			max_distance /= 2
		
		return max_distance


