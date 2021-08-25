package gogo

func genMove(x, y int, player byte) Move {
  return Move{Player: player, Position: Coordinate{X: x, Y: y}}
}

func (gameboard GameBoard) copy() GameBoard {
  copiedBoard := newBoard(cap(gameboard.Positions))
  copy(copiedBoard.Positions, gameboard.Positions)
  return copiedBoard
}

func newBoard(size int) GameBoard {
  board := GameBoard{}
  a := make([][]byte, size)
  for i := range a {
    a[i] = make([]byte, size)
  }
  board.Positions = a
  return board
}
