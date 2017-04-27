package main

func isPlayable(x int, y int, state[][]int) bool {
  if (state[x][y] != 0) {
    return false
  }
  var i = -AMP
  for i <= AMP {
    var j = -AMP
    for j <= AMP {
      if (x + i > 0 && x + i < SIZE && y + j > 0 && y + j < SIZE && state[x + i][y + j] != 0) {
        return true
      }
      j++
    }
    i++
  }
  return false
}

func findMoves(state[][]int) ([]move) {

  moves := make([]move, 0)
  var x = 0
  for x < SIZE {
    var y = 0
    for y < SIZE {
      if (isPlayable(x, y, state)) {
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
