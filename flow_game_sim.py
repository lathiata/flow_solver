from flow_game_state import FlowGameState
from flow_game_problem import FlowGameProblem
from flow_search import *
from flow_game_constants import TransitionModels

def csp_test():
	fg = FlowGameProblem(5, [[(0,2),(3,0)], [(0,3),(4,3)], [(1,3),(2,1)], [(3,3),(4,4)], [(3,1),(4,0)]])
	print(fg.get_start_state())
	print('Min Conflicts')
	cspSearch(fg)


def test1():
	fg = FlowGameProblem(5, [[(0,2),(3,0)], [(0,3),(4,3)], [(1,3),(2,1)], [(3,3),(4,4)], [(3,1),(4,0)]])
	print(fg.get_start_state())
	print('aStarSearch v0')
	result = aStarSearch(fg, TransitionModels.VERSION_0)
	print(result)
	print('aStarSearch v1')
	result = aStarSearch(fg, TransitionModels.VERSION_1)
	print(result)

def test2():
	fg = FlowGameProblem(8, [[(0,0),(3,0)], [(4,0),(7,6)], [(2,1),(3,5)], [(2,2),(1,6)], [(2,3),(3,6)], [(4,1),(6,1)],
						[(3,1), (4,4)], [(4,3),(6,6)], [(5,3),(7,7)]])
	print(fg.get_start_state())
	print('flow game search')
	result = flow_game_search(fg, TransitionModels.VERSION_0)
	print(result)
	print('flow game search v1')
	result1 = flow_game_search(fg, TransitionModels.VERSION_1)
	print(result1)

def test3():
	fg = FlowGameProblem(8, [[(1,6),(2,1)], [(2,2),(7,3)], [(3,1),(6,6)], [(4,2),(7,7)], [(5,4),(6,5)], [(6,1),(7,6)]])
	print(fg.get_start_state())
	print('flow game search')
	result = flow_game_search(fg, TransitionModels.VERSION_0)
	print(result)
	print('flow game search v1')
	result1 = flow_game_search(fg, TransitionModels.VERSION_1)
	print(result1)

def test4():
	fg = FlowGameProblem(8, [[(0,0),(5,2)], [(1,1),(3,0)], [(1,2),(2,6)], [(1,6),(5,4)], [(4,0),(5,3)]])
	print(fg.get_start_state())
	print('flow game search')
	result = flow_game_search(fg, TransitionModels.VERSION_0)
	print(result)
	print('flow game search v1')
	result1 = flow_game_search(fg, TransitionModels.VERSION_1)
	print(result1)


csp_test()
# test1()
# test2()
# test3()
# test4()