from Tkinter import *

class Number(Canvas):
    def __init__(self, master=None, number=1):
        Canvas.__init__(self, master, width=100, height=100, bd=-2, highlightthickness=0)
        self.create_text(50, 50, text=str(number), anchor=CENTER)
        self.create_rectangle(4, 4, 98, 98, width=2)


class Application(Frame):
	def callback():
		print("okay")

	def create_widgets(self):
		self.board = Frame(self)
		self.board.grid_propagate(0)
		self.board.grid()
		self.board.numbers = []
		size = 3
		for i in range(size):
			for j in range(size):
				b = Button(self.board, text=str(i) + ", " + str(j), command=self.callback)
				self.board.numbers.append(b)
				b.grid(row=i, column=j)
				print("yo")

	def __init__(self, master=None):
		Frame.__init__(self, master)
		self.grid()
		self.create_widgets()


def main():
    app = Application()
    app.master.title("Flow Game")
    app.mainloop()

if __name__ == '__main__':
    main()