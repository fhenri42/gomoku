package main

func isClose(x int, y int, board[SIZE][SIZE]int) bool {
  if (board[x][y] != 0) {
    return false
  }
  var i = -AMP
  for i <= AMP {
    var j = -AMP
    for j <= AMP {
      if (x + i > 0 && x + i < SIZE && y + j > 0 && y + j < SIZE && board[x + i][y + j] != 0) {
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
  for x < SIZE {
    var y = 0
    for y < SIZE {
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
