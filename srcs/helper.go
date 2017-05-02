package main

func isPlayable(board [SIZE][SIZE]int, x int, y int) bool {
  if (x >= SIZE || y >= SIZE || board[x][y] != 0) {
    return false
  }
  return true
}

func moveAndEat(board [SIZE][SIZE]int, x int, y int, player int, aiScore *int, playerScore *int) [SIZE][SIZE]int {
  otherPlayer := player % 2 + 1
  var i = -1
  var j = -1

  for i <= 1 {
    j = -1
    for j <= 1 {
      if (x + i * 3 > 0 && x + i * 3 < SIZE && y + j * 3 > 0 && y + j * 3 < SIZE && board[x + i][y + j] == otherPlayer && board[x + i * 2][y + j * 2] == otherPlayer && board[x + i * 3][y + j * 3] == player) {
        board[x + i][y + j] = 0
        board[x + i * 2][y + j * 2] = 0
        if (player == PLAYER1) {
          *playerScore += 2
        } else {
          *aiScore += 2
        }
      }
      j++
    }
    i++
  }
  board[x][y] = player
  return board
}

func endGame(board [SIZE][SIZE]int) (bool, int) {
  var x = 0
  for x < SIZE {
    var y = 0
    for y < SIZE {
      if (isUnbreakableLine(x, y, board)) {
        return true, board[x][y]
      }
      y++
    }
    x++
  }
  return false, 0
}

func isUnbreakableLine(i int, j int, board [SIZE][SIZE]int) bool {
  var pion = board[i][j]

  if (pion == 0) {
    return false
  }

  if i + 4 <= SIZE && board[i][j] == pion && board[i + 1][j] == pion && board[i + 2][j] == pion && board[i + 3][j] == pion && board[i + 4][j] == pion {
    return true

  } else if j + 4 <= SIZE && board[i][j] == pion && board[i][j + 1] == pion && board[i][j + 2] == pion && board[i][j + 3] == pion && board[i][j + 4] == pion {
    return true

  } else if j + 4 <= SIZE && i + 4 <= SIZE && board[i][j] == pion && board[i + 1][j + 1] == pion && board[i + 2][j + 2] == pion && board[i + 3][j + 3] == pion && board[i + 4][j + 4] == pion {
    return true

  } else if j - 4 >= 0 && i + 4 <= SIZE && board[i][j] == pion && board[i + 1][j - 1] == pion && board[i + 2][j - 2] == pion && board[i + 3][j - 3] == pion && board[i + 4][j - 4] == pion {
    return true
  }
  return false
}

// func eatable() bool {
//
// }
