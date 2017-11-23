package pacman

type Position struct {
  x int
  y int
}

type Board struct {
  currentDirection string
  currentPosition  *Position
  Board            [][]string
}

func (board *Board) Tick() {

  var position = board.currentPosition

  board.Board[position.x][position.y] = " "

  switch  board.currentDirection  {
  case "V" :
	goUp(position)
  case ">" :
	goLeft(position)
  case "^" :
	goDown(position)
  }

  board.Board[position.x][position.y] = board.currentDirection
}

func (board *Board) Left() {
  board.currentDirection = ">"
}

func (board *Board) Down() {
  board.currentDirection = "^"
}

func goUp(position *Position) {
  if isOnTop(*position) {
	position.x = 2
  } else {
	position.x = position.x - 1
  }
}

func goLeft(position *Position) {
  position.y = position.y - 1
}

func goDown(position *Position) {
  position.x = position.x + 1
}

func isOnTop(position Position) bool {
  return position.x == 0
}

func NewBoard() Board {

  return Board{"V", &Position{1, 1}, [][]string{
	{".", ".", "."},
	{".", "V", "."},
	{".", ".", "."},
  },
  };
}
