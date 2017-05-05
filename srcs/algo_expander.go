package main

func isClose(x int, y int, board[SIZE][SIZE]int, movesOfone*[]move, movesOfTwo *[]move) {
  if (board[x][y] != 0) {
    return
  }
  var i int = -AMP
  for i <= AMP {
    var j int = -AMP
    for j <= AMP {
        if ((x + i) + i >= 0 && (x + i) + i < SIZE && (y + j) + j >= 0 && (y + j) + j < SIZE && board[x + i][y + j] != 0 && board[(x + i) + i][( y + j) + j] != 0) {
          *movesOfTwo = append(*movesOfTwo, newMove(x, y, 0))
          return
        }
        if (x + i >= 0 && x + i < SIZE && y + j >= 0 && y + j < SIZE && board[x + i][y + j] != 0) {
          *movesOfone = append(*movesOfone, newMove(x, y, 0))
        return
      }
      j++
    }
    i++
  }
  return
}

func findMoves(tools *sdlTools) ([]move) {

  movesOfone := make([]move, 0)
  movesOfTwo := make([]move, 0)

  var x int = 0
  for x < SIZE {
    var y int = 0
    for y < SIZE {
      if isPlayable(tools,x,y) {
        isClose(x,y, tools.board, &movesOfone, &movesOfTwo)
      }
      y++
    }
    x++
  }

if len(movesOfTwo) > 0  {
  return movesOfTwo
}
  return movesOfone
}
