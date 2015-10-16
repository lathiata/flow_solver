from flow_game_state import FlowGameState
from flow_search import aStarSearch, breadthFirstSearch, flow_game_search
from copy import copy, deepcopy

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

	def heuristic(self, flow_game_state):
		curr_positions = flow_game_state.get_curr_positions()
		goal_positions = flow_game_state.get_goal_positions()
		max_distance = -float('inf')
		for i in range(len(curr_positions)):
			dist = abs(curr_positions[i][0] - goal_positions[i][0]) + abs(curr_positions[i][1] - goal_positions[i][1])
			max_distance = max(max_distance, dist)
		return max_distance

#light blue, yellow, red, blue, green, pink, purp, orange, brown
# fg = FlowGameProblem(5, [[(0,2),(3,0)], [(0,3),(4,3)], [(1,3),(2,1)], [(3,3),(4,4)], [(3,1),(4,0)]])
# fg = FlowGameProblem(8, [[(0,0),(3,0)], [(4,0),(7,6)], [(2,1),(3,5)], [(2,2),(1,6)], [(2,3),(3,6)], [(4,1),(6,1)],
# 						[(3,1), (4,4)], [(4,3),(6,6)], [(5,3),(7,7)]])
fg = FlowGameProblem(8, [[(1,6),(2,1)], [(2,2),(7,3)], [(3,1),(6,6)], [(4,2),(7,7)], [(5,4),(6,5)], [(6,1),(7,6)]])
# fg = FlowGameProblem(8, [[(0,0),(5,2)], [(1,1),(3,0)], [(1,2),(2,6)], [(1,6),(5,4)], [(4,0),(5,3)]])
print(fg.get_start_state())
print('flow game search')
result = flow_game_search(fg)
print(result)
print('astar')
result2 = aStarSearch(fg, True)
print(result2)
# print('bfs')
# result = breadthFirstSearch(fg)
# print(result)

