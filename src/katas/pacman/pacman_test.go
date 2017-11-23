package pacman

import "test"
import "testing"

var pacmanLookUp = "V"
var pacmanLookLeft = ">"
var pacmanLookDown = "^"

func TestPacmanStartAtCenter(t *testing.T) {

  var board = NewBoard()
  test.AssertEquals(pacmanLookUp, board.Board[1][1], t)

}

func TestPacmanMoveUp(t *testing.T) {

  var board = NewBoard()

  board.Tick()

  test.AssertEquals(pacmanLookUp, board.Board[0][1], t)
  test.AssertEquals(" ", board.Board[1][1], t)

}

func TestPacmanWrapAroundWhenAtTop(t *testing.T) {

  var board = NewBoard()

  board.Tick()
  board.Tick()

  test.AssertEquals(pacmanLookUp, board.Board[2][1], t)
  test.AssertEquals(" ", board.Board[0][1], t)

}

func TestPacmanGoLeft(t *testing.T) {

  var board = NewBoard()

  board.Tick()
  board.Left()
  board.Tick()

  test.AssertEquals(pacmanLookLeft, board.Board[0][0], t)
  test.AssertEquals(" ", board.Board[1][1], t)
  test.AssertEquals(" ", board.Board[0][1], t)

}

func TestPacmanGoDown(t *testing.T) {

  var board = NewBoard()

  board.Down()
  board.Tick()

  test.AssertEquals(pacmanLookDown, board.Board[2][1], t)
  test.AssertEquals(" ", board.Board[1][1], t)

}


