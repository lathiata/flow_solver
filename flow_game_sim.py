import argparse
from random import choice

from flow_game_problem import FlowGameProblem
from flow_search import flow_game_search
from flow_game_constants import TransitionModels


# TODO	change VERSION to specify the search type (min conflicts, flow_search, k-beams, etc)

# TODO move to constants ?
EASY_PUZZLES = [
    FlowGameProblem(5, [[(0, 2), (3, 0)], [(0, 3), (4, 3)], [(1, 3), (2, 1)], [(3, 3), (4, 4)], [(3, 1), (4, 0)]])]
MEDIUM_PUZZLES = [FlowGameProblem(8, [[(0, 0), (3, 0)], [(4, 0), (7, 6)], [(2, 1), (3, 5)], [(2, 2), (1, 6)],
                                      [(2, 3), (3, 6)], [(4, 1), (6, 1)],
                                      [(3, 1), (4, 4)], [(4, 3), (6, 6)], [(5, 3), (7, 7)]])]
HARD_PUZZLES = [FlowGameProblem(8, [[(1, 6), (2, 1)], [(2, 2), (7, 3)], [(3, 1), (6, 6)], [(4, 2), (7, 7)],
                                    [(5, 4), (6, 5)], [(6, 1), (7, 6)]]),
                FlowGameProblem(8, [[(0, 0), (5, 2)], [(1, 1), (3, 0)], [(1, 2), (2, 6)], [(1, 6), (5, 4)],
                                    [(4, 0), (5, 3)]])]

DIFFICULTY = {"easy": EASY_PUZZLES, "medium": MEDIUM_PUZZLES, "hard": HARD_PUZZLES}
SEARCH_MODEL = {"single": TransitionModels.VERSION_1,
                "multi": TransitionModels.VERSION_0}
SEARCH_MODEL_NAMES = {TransitionModels.VERSION_0: "multi",
                      TransitionModels.VERSION_1: "single"}


def run_simulation(model, problem, verbose):
    print("Solving problem using the {0} transition model".format(SEARCH_MODEL_NAMES[model]))
    print(problem.get_start_state())
    print(flow_game_search(problem, model, verbose))


def validate_args(args):
    """
    validates command line arguments
    returns is_valid, FlowGameProblem, TransitionModel, Verbose
    """
    if args.difficulty not in DIFFICULTY.keys() or args.search_model not in SEARCH_MODEL.keys():
        return False, None, None, None
    return True, choice(DIFFICULTY[args.difficulty]), SEARCH_MODEL[args.search_model], args.verbose


def main():
    parser = argparse.ArgumentParser()
    parser.add_argument("--difficulty", nargs='?', help="complexity of puzzle: easy, medium, hard", type=str,
                        default="easy")
    parser.add_argument("--search_model", nargs='?', help="name of search model: single, multi", type=str,
                        default="single")
    parser.add_argument("--verbose", nargs='?', help="verbose output to std out", type=bool,
                        default=False)
    args = parser.parse_args()
    success, problem, model, verbose = validate_args(args)
    if success:
        run_simulation(model, problem, verbose)
    else:
        print("use the -h flag to see help for valid inputs")


if __name__ == "__main__":
    main()
