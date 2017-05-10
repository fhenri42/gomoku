package main
import (
  "math"
)

// Evaluate a grid
func evaluate(game *Game) Move {
  var res int = 0
  var i int = 0
  var j int = 0

  for i < game.score[game.friend - 1] {
    res += i//int(math.Pow(float64(i), 2))
    i++
  }
  i = 0
  for i < game.score[game.friend % 2] {
    res -= i//int(math.Pow(float64(i), 2))
    i++
  }
  i = 0
  for i < SIZE {
    j = 0
    for j < SIZE {
      if (game.board[i][j] == game.friend) {
        res += evaluateEach(&game.board, i, j)
      } else if (game.board[i][j] == game.friend % 2 + 1) {
        res -= evaluateEach(&game.board, i, j)
      }
      j++
    }
    i++
  }
  return (newMove(0, 0, res))
}

// Evaluate a Pion
func evaluateEach(board *[SIZE][SIZE]int, x int, y int) int {
  var i int = -1
  var j int = -1
  var k int = 1
  var player = board[x][y]
  var res int = 0
  var countFree int = 0
  var countFriend int = 0

  if (isEatable(x, y, board)) {
    res -= 10
  }

  for i <= 1 {
    j = -1
    for j <= 1 {

      countFree = 0
      countFriend = 0

      k = -1
      for (k >= -4 && x + i * k >= 0 && x + i * k < SIZE && y + j * k >= 0 && y + j * k < SIZE && (*board)[x + i * k][y + j * k] != player % 2 + 1) {
        if ((*board)[x + i * k][y + j * k] == player) {
          countFriend++
        } else {
          countFree++
        }
        k--
      }

      k = 1
      for (k <= 4 && x + i * k >= 0 && x + i * k < SIZE && y + j * k >= 0 && y + j * k < SIZE && (*board)[x + i * k][y + j * k] != player % 2 + 1) {
        if ((*board)[x + i * k][y + j * k] == player) {
          countFriend++
        } else {
          countFree++
        }
        k++
      }

      if (countFree + countFriend >= 5) {
        k = 0
        for k < countFree {
          res += i
          k++
        }

        k = 0
        for k < countFree {
          res += int(math.Pow(float64(k + 1), 2))
          k++
        }
      }


      j++
    }
    i++
  }

  return res
}
