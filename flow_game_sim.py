from flow_game_state import FlowGameState
from flow_game_problem import FlowGameProblem
from flow_search import *
from flow_game_constants import TransitionModels

# TODO: make a command line argument parser to specify search model and problem difficulty
# TODO: specify problems by difficuly 
# TODO: change VERSION to specify the search type (min conflicts, greedy, k-beams, etc)

PROBLEM_NUM = 0 #indexes into problems list below
PRINT = False
VERSION = TransitionModels.VERSION_1

problems = [FlowGameProblem(5, [[(0,2),(3,0)], [(0,3),(4,3)], [(1,3),(2,1)], [(3,3),(4,4)], [(3,1),(4,0)]]),
			FlowGameProblem(8, [[(0,0),(3,0)], [(4,0),(7,6)], [(2,1),(3,5)], [(2,2),(1,6)], [(2,3),(3,6)], [(4,1),(6,1)],
						[(3,1), (4,4)], [(4,3),(6,6)], [(5,3),(7,7)]]),
			FlowGameProblem(8, [[(1,6),(2,1)], [(2,2),(7,3)], [(3,1),(6,6)], [(4,2),(7,7)], [(5,4),(6,5)], [(6,1),(7,6)]]),
			FlowGameProblem(8, [[(0,0),(5,2)], [(1,1),(3,0)], [(1,2),(2,6)], [(1,6),(5,4)], [(4,0),(5,3)]])]

def test(version, problem):
	print(problem.get_start_state())
	print(version)
	result = flow_game_search(problem, version, PRINT)
	print(result)

test(VERSION, problems[PROBLEM_NUM])
