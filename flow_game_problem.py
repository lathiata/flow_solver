from flow_game_state import FlowGameState
from copy import copy, deepcopy
import random
from flow_game_constants import TransitionModels

class FlowGameProblem:
	def __init__(self, gridsize, start_positions):
		self.start_state = FlowGameState(gridsize, start_positions)

	def get_start_state(self):
		return self.start_state 

	def goal_test(self, flow_game_state):
		'''
		Determine if the game is in a solved state
		Only works if you are exploring a tree in a way you play the game 
		Does not work with CSP search which can have break rules
		'''
		for i in range(len(flow_game_state.get_curr_positions())):
			if not flow_game_state.reached_goal(i):
				return False
		return flow_game_state.full_grid()

	def explore_v0(self, state, frontier):
		'''
		v0 has all colors moving towards their goal states at the same time
		so a move is a collection of moves for all colors 
		'''
		def get_actions_v0(flow_game_state):
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

		def get_result_v0(flow_game_state, action):
			next_fgs = deepcopy(flow_game_state)
			for i in action.keys():
				if action[i]:
					next_fgs.move(i, action[i])
			return next_fgs

		possible_actions = get_actions_v0(state)
		for action in possible_actions:
			new_state = get_result_v0(state, action)
			frontier.push(new_state, self.heuristic(new_state))

	def explore_v1(self, state, frontier):
		'''
		v1 will solve one color at a time
		'''
		def get_actions_v1(flow_game_state):
			curr_positions = flow_game_state.get_curr_positions()
			for i in range(len(curr_positions)):
				if not flow_game_state.reached_goal(i):
					return i, flow_game_state.legal_moves(curr_positions[i])

		def get_result_v1(flow_game_state, action, i):
			next_fgs = deepcopy(flow_game_state)
			next_fgs.move(i, action)
			return next_fgs

		color, possible_actions = get_actions_v1(state)
		for action in possible_actions:
			new_state = get_result_v1(state, action, color)
			frontier.push(new_state, self.heuristic(new_state))

	#TODO: create a mapping TransitionModel.Version -> explore function
	def explore(self, version, state, frontier):
		'''
		controller function for the different transition models
		'''
		if version == TransitionModels.VERSION_0:
			self.explore_v0(state, frontier)
		elif version == TransitionModels.VERSION_1:
			self.explore_v1(state, frontier)

	#TODO: implement better heuristic
	def heuristic(self, flow_game_state):
		curr_positions = flow_game_state.get_curr_positions()
		goal_positions = flow_game_state.get_goal_positions()
		sum_distance = 0
		# max_distance = -float('inf')
		for i in range(len(curr_positions)):
			dist = abs(curr_positions[i][0] - goal_positions[i][0]) + abs(curr_positions[i][1] - goal_positions[i][1])
			# max_distance = max(max_distance, dist)
			sum_distance += dist

		#try giving preference to the edges
		if flow_game_state.border(curr_positions[i]):
			# max_distance /= 4
			sum_distance /= 4

		return sum_distance
		# return max_distance

# beam search?