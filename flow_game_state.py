import sys
from termcolor import colored, cprint
class FlowGameState:
	def __init__(self, gridsize, start_positions):
		self.colors = ['grey', 'red', 'green', 'yellow', 'blue', 'magenta', 'cyan']
		self.on_colors = ['on_grey', 'on_red', 'on_green', 'on_yellow', 'on_blue', 'on_magenta', 'on_cyan']
		#initialize board
		self.domains = {} #maps position to domains
		self.gridsize = gridsize
		self.unoccupied_space = self.gridsize ** 2
		self.board = []
		self.start_positions = []
		self.curr_positions = []
		self.goal_positions = []
		self.val_colors = {}
		j = 0
		while j < gridsize:
			self.board.append([])
			for i in range(self.gridsize):
				self.board[j].append(None)
			j += 1
		#add start/goal positions
		for i in range(len(start_positions)):
			#SWAPPED I AND J
			self.board[start_positions[i][0][1]][start_positions[i][0][0]] = chr(i + ord('a'))
			self.curr_positions.append((start_positions[i][0][1], start_positions[i][0][0]))

			self.start_positions.append((start_positions[i][0][1], start_positions[i][0][0]))

			# self.board[start_positions[i][1][1]][start_positions[i][1][0]] = chr(i + ord('a'))

			self.goal_positions.append((start_positions[i][1][1], start_positions[i][1][0]))

			self.unoccupied_space -= 1

		#add domains
		for i in range(gridsize):
			for j in range(gridsize):
				if (i, j) not in self.start_positions:
					self.domains[(i, j)] = range(gridsize)


	def adjacent_positions(self, pos):
		#returns a list of valid adjacent positions
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

	def get_gridsize(self):
		return self.gridsize

	def get_start_positions(self):
		return self.start_positions

	def is_adjacent(self, pos1, pos2):
		return abs(pos1[0]-pos2[0]) + abs(pos1[1] - pos2[1]) == 1

	def is_goal(self, pos, i):
		return pos == self.goal_positions[i]

	def get_board(self):
		return self.board

	def full_grid(self):
		return self.unoccupied_space == 0

	def get_curr_positions(self):
		return self.curr_positions

	def get_goal_positions(self):
		return self.goal_positions

	def move(self, i, position):
		self.board[position[0]][position[1]] = chr(i + ord('a'))
		self.curr_positions[i] = position
		self.unoccupied_space -= 1

	def border(self, pos):
		return pos[0] == 0 or pos[0] == self.gridsize-1 or pos[1] == 0 or pos[1] == self.gridsize-1

	#maybe move to flow game problem?
	def legal_moves(self, current_position):
		moves = []
		i = current_position[0]
		j = current_position[1]

		if i - 1 >= 0 and not self.board[i-1][j]:
			moves.append((i-1, j))
		if i + 1 < self.gridsize and not self.board[i+1][j]:
			moves.append((i+1, j))
		if j - 1 >= 0 and not self.board[i][j-1]:
			moves.append((i, j-1))
		if j + 1 < self.gridsize and not self.board[i][j+1]:
			moves.append((i, j+1))

		return moves


	def __repr__(self):
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
				val = self.board[i][j] if self.board[i][j] else " "
				if (i, j) in self.goal_positions:
					val = chr(self.goal_positions.index((i, j)) + ord("a"))
				if val not in self.val_colors.keys():
					if ind < len(self.colors):
						self.val_colors[val] = colored(val, self.colors[ind])
						ind += 1
					else:
						first = ind/len(self.colors)
						second = ind%len(self.colors)
						self.val_colors[val] = colored(val, self.colors[first], self.on_colors[second]) 
						ind += 1
				val_colored = self.val_colors[val]
				representation += val_colored + "|"
			representation += "\n"
		return representation


