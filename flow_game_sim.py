import argparse
import random
from flow_game_state import FlowGameState
from flow_game_problem import FlowGameProblem
from flow_search import *
from flow_game_constants import TransitionModels

# TODO: change VERSION to specify the search type (min conflicts, flow_search, k-beams, etc)
#		will be interesting bc VERSION refers to transition model but these are search algs
PRINT = False

EASY_PUZZLES = [FlowGameProblem(5, [[(0,2),(3,0)], [(0,3),(4,3)], [(1,3),(2,1)], [(3,3),(4,4)], [(3,1),(4,0)]])]
MEDIUM_PUZZLES = [FlowGameProblem(8, [[(0,0),(3,0)], [(4,0),(7,6)], [(2,1),(3,5)], [(2,2),(1,6)], [(2,3),(3,6)], [(4,1),(6,1)],
							 		 [(3,1), (4,4)], [(4,3),(6,6)], [(5,3),(7,7)]])]
HARD_PUZZLES = [FlowGameProblem(8, [[(1,6),(2,1)], [(2,2),(7,3)], [(3,1),(6,6)], [(4,2),(7,7)], [(5,4),(6,5)], [(6,1),(7,6)]]),
				FlowGameProblem(8, [[(0,0),(5,2)], [(1,1),(3,0)], [(1,2),(2,6)], [(1,6),(5,4)], [(4,0),(5,3)]])]

DIFFICULTY = {"easy": EASY_PUZZLES, "medium": MEDIUM_PUZZLES, "hard": HARD_PUZZLES}
SEARCH_MODEL = {"single": TransitionModels.VERSION_1, "multi": TransitionModels.VERSION_0}

def test(version, problem):
	print(problem.get_start_state())
	print(version)
	result = flow_game_search(problem, version, PRINT)
	print(result)

def validate_args(args):
	'''
	returns bool, problem, search_model
	'''
	if args.difficulty not in DIFFICULTY.keys() or args.search_model not in SEARCH_MODEL.keys():
		return False, None, None
	return True, random.choice(DIFFICULTY[args.difficulty]), SEARCH_MODEL[args.search_model]


parser = argparse.ArgumentParser()
parser.add_argument("difficulty", nargs='?', help="complexity of puzzle: easy, medium, hard", type=str, default="hard")
parser.add_argument("search_model", nargs='?', help="name of search model: single, multi", type=str, default="single")
args = parser.parse_args()
success, problem, version = validate_args(args)

if success:
	test(version, problem)
else:
	print("use the -h flag to see help for valid inputs")
