from enum import Enum

EMPTY = -1
REORDER = False


# TODO Eventually will change to be different types of searches completely (somehow)
class TransitionModels(Enum):
    VERSION_0 = 0
    VERSION_1 = 1
