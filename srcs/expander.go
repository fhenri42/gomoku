package main

func isClose(x int, y int, board[SIZE][SIZE]int) bool {
  if (board[x][y] != 0) {
    return false
  }
  var i = -AMP
  for i <= AMP {
    var j = -AMP
    for j <= AMP {
      if (x + i > 0 && x + i < SIZE - 1 && y + j > 0 && y + j < SIZE - 1 && board[x + i][y + j] != 0) {
        return true
      }
      j++
    }
    i++
  }
  return false
}

func findMoves(board[SIZE][SIZE]int) ([]move) {

  moves := make([]move, 0)
  var x = 0
  for x < SIZE - 1 {
    var y = 0
    for y < SIZE - 1 {
      if (isClose(x, y, board)) {
        var tmp move
        tmp.x = x
        tmp.y = y
        moves = append(moves, tmp)
      }
      y++
    }
    x++
  }
  return moves
}
