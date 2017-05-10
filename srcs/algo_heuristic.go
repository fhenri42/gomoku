package main
import (
  //"math"
)

const EATABLE_MALUS = 10
const SCORE_BONUS = 1
const SPACE_5_BONUS = 10
const SPACE_9_BONUS = 10
const SUP_SPACE_BONUS = 2
const ALIGNED_BONUS = 5

// Evaluate a grid
func evaluate(game *Game) Move {
  var res int = 0
  var i int = 0
  var j int = 0

  for i < game.score[game.friend - 1] {
    res += SCORE_BONUS
    i++
  }
  i = 0
  for i < game.score[game.friend % 2] {
    res -= SCORE_BONUS
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
    res -= EATABLE_MALUS
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

      if (countFree + countFriend >= 4) {
        res += SPACE_5_BONUS
        if (countFree + countFriend >= 8) {
          res += SPACE_9_BONUS
        }
        k = 0
        for k + 4 < countFree + countFriend {
          res += SUP_SPACE_BONUS
          k++
        }

        k = 0
        for k < countFriend {
          res += ALIGNED_BONUS
          k++
        }
      }


      j++
    }
    i++
  }

  return res
}
