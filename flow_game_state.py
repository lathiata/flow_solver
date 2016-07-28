import sys
import flow_game_constants
from termcolor import colored, cprint

# TODO: have each cell remember the direction the previous cell was going
#       create readme for the file
#       implement min conflicts
#       implement k-beam search utilizing multiprocessing
#		use pep8 style docstring comments

class FlowGameState:
	def __init__(self, gridsize, start_positions):
		"""
		Initialize a Flow gameboard given gridsize and start and goal positions.
		"""
		self.colors = ['red', 'green', 'blue', 'magenta', 'cyan', 'white', 'yellow']
		self.on_colors = ['on_red', 'on_green', 'on_blue', 'on_magenta', 'on_cyan']

		#initialize board, empty spaces are flow_game_constants.EMPTY
		self.gridsize = gridsize
		self.unoccupied_space = self.gridsize ** 2
		self.board = []
		self.start_positions = []
		self.curr_positions = []
		self.goal_positions = []
		self.val_colors = {}
		self.num_colors = len(start_positions)

		j = 0
		while j < gridsize:
			self.board.append([])
			for i in range(self.gridsize):
				self.board[j].append(flow_game_constants.EMPTY)
			j += 1

		#add start/goal positions
		for i in range(len(start_positions)):
			#SWAPPED I AND J

			#add starting positions
			self.board[start_positions[i][0][1]][start_positions[i][0][0]] = i
			self.curr_positions.append((start_positions[i][0][1], start_positions[i][0][0]))
			self.start_positions.append((start_positions[i][0][1], start_positions[i][0][0]))

			#add goal positions
			self.board[start_positions[i][1][1]][start_positions[i][1][0]] = i
			self.goal_positions.append((start_positions[i][1][1], start_positions[i][1][0]))

			#minus two because we add both the 'starting' and the 'goal' positions
			self.unoccupied_space -= 2

		#TODO: come up with a better way to decide whether or not to reorder
		if flow_game_constants.REORDER:
			start_tup, goal_tup = zip(*sorted(zip(self.start_positions, self.goal_positions),\
							   	  			  key= lambda x: abs(x[0][0] - x[1][0]) + abs(x[0][1] - x[1][1])))
			self.start_positions = list(start_tup)
			self.curr_positions = list(self.start_positions)
			self.goal_positions = list(goal_tup)

	def get_num_colors(self):
		return self.num_colors

	def adjacent_positions(self, pos):
		"""
		Returns a list of valid adjacent positions.
		"""
		adj_pos = []
		i, j = pos

		if i > 0:
			adj_pos.append((i-1, j))

		if i < self.gridsize-1:
			adj_pos.append((i+1, j))

		if j > 0:
			adj_pos.append((i, j-1))

		if j < self.gridsize-1:
			adj_pos.append((i, j+1))

		return adj_pos

	def get_val(self, i, j):
		return self.board[i][j]

	def get_pos_val(self, pos):
		return self.board[pos[0]][pos[1]]

	def get_gridsize(self):
		return self.gridsize

	def get_start_positions(self):
		return self.start_positions

	def is_adjacent(self, pos1, pos2):
		return abs(pos1[0]-pos2[0]) + abs(pos1[1] - pos2[1]) == 1

	def get_board(self):
		return self.board

	def full_grid(self):
		return self.unoccupied_space == 0

	def get_curr_positions(self):
		return self.curr_positions

	def get_goal_positions(self):
		return self.goal_positions

	def move(self, i, position):
		self.board[position[0]][position[1]] = i
		self.curr_positions[i] = position
		self.unoccupied_space -= 1

	def reached_goal(self, i):
		"""
		True if the current position of the i_th color is adjacent to the i_th goal position.
		We check this because the goal position is filled into the game board on initialization.
		"""
		curr_pos = self.curr_positions[i]
		goal_pos = self.goal_positions[i]
		return self.is_adjacent(curr_pos, goal_pos)

	def border(self, pos):
		return pos[0] == 0 or pos[0] == self.gridsize-1 or pos[1] == 0 or pos[1] == self.gridsize-1

	def solvable_helper(self, i):
		"""
		Helper function for determining whether or not the board is currently in a solvable config.
		This works by performing every legal move for a color and then checking if the color has
		filled a square adjacent to the goal.
		"""
		goal_pos = self.goal_positions[i]
		for pos in self.adjacent_positions(goal_pos):
			if self.board[pos[0]][pos[1]] == i:
				return True
		return False


	def legal_moves(self, curr_pos):
		"""
		Returns a list of (x,y) tuples that correspond to open positions adjacent to
		curr_pos.
		"""
		return [pos for pos in self.adjacent_positions(curr_pos) if self.board[pos[0]][pos[1]] == flow_game_constants.EMPTY]


	def __repr__(self):
		"""
		Writes a grid representation of the flow game with nice colors.
		"""
		ind = 0
		representation = "  "

		#add column headers
		for i in range(self.gridsize):
			pipe = "" if i == self.gridsize - 1 else "|"
			representation += str(i) + pipe
		representation += "\n"

		#fill in rest of grid
		for i in range(self.gridsize):
			representation += str(i) + "|"
			for j in range(self.gridsize):
				val = self.board[i][j] if self.board[i][j] >= 0 else " "

				if val not in self.val_colors.keys():
					if ind < len(self.colors):
						self.val_colors[val] = colored(val, self.colors[ind])
					else:
						first = ind/len(self.colors)
						second = ind%len(self.colors)
						self.val_colors[val] = colored(val, self.colors[first], self.on_colors[second])
					ind += 1
				val_colored = self.val_colors[val]

				representation += val_colored + "|"
			representation += "\n"
		return representation
