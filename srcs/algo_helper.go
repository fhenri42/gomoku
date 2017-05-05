package main

func moveAndEat(board [SIZE][SIZE]int, x int, y int, player int, aiScore *int, playerScore *int) ([SIZE][SIZE]int, bool, bool) {
  var res bool = false
  var breakable bool = false
  otherPlayer := player % 2 + 1
  var i = -1
  var j = -1
  var k = -4
  var count = 0
  var countMax = 0
  var posi = -1
  var posj = -1

  for i <= 1 {
    j = -1
    for j <= 1 {
      k = -4
      count = 0
      countMax = 0
      for k <= 4 {
        if (k == 0 || (x + i * k >= 0 && x + i * k < SIZE && y + j * k >= 0 && y + j * k < SIZE && board[x + i * k][y + j * k] == player)) {
          count++
          if (count == 5) {
            posi = x + i * k
            posj = y + j * k
          }
          countMax = count
        } else {
          count = 0
        }
        k++
      }
      if (countMax >= 5) {
        if (isUnbreakableLine(posi, posj, i, j, board)) {
          res = true
        } else {
          breakable = true
        }
      }
      if (x + i * 3 >= 0 && x + i * 3 < SIZE && y + j * 3 >= 0 && y + j * 3 < SIZE && board[x + i][y + j] == otherPlayer && board[x + i * 2][y + j * 2] == otherPlayer && board[x + i * 3][y + j * 3] == player) {
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
  if (*aiScore == 10 || *playerScore == 10) {
    res = true
  }

  return board, res, breakable
}

func isUnbreakableLine(i int, j int, x int, y int, board [SIZE][SIZE]int) bool {
  var count = 0

  for count <= 5 {
    if (isEatable(x + count * i, y + count * j, board)) {
      return false
    }
    count++
  }
  return true
}

func isEatable(x int, y int, board [SIZE][SIZE]int) bool {
  var pion = board[x][y]
  var pion2 = 0

  if (pion == 0) {
    return false
  } else if (pion == 1) {
    pion2 = 2
  } else {
    pion2 = 1
  }

  var i = -1
  var j = -1

  for i <= 1 {
    j = -1
    for j <= 1 {
      if (x + i * -1 >= 0 && x + i * -1 < SIZE && y + j * -1 < SIZE && y + j * -1 >= 0 && x + i * 2 >= 0 && x + i * 2 < SIZE && y + j * 2 >= 0 && y + j * 2 < SIZE && board[x + i * -1][y + j * -1] == 0 && board[x + i][y + j] == pion && board[x + i * 2][y + j * 2] == pion2) {
        return true
      }
      if (x + i * -2 >= 0 && x + i * -2 < SIZE && y + j * -2 < SIZE && y + j * -2 >= 0 && x + i * 1 >= 0 && x + i * 1 < SIZE && y + j * 1 >= 0 && y + j * 1 < SIZE && board[x + i * -1][y + j * -1] == pion && board[x + i][y + j] == pion2 && board[x + i * -2][y + j * -2] == 0) {
        return true
      }
      j++
    }
    i++
  }
  return false
}

func isPlayable(tools *sdlTools, x int, y int) bool {
  if (x >= SIZE || y >= SIZE || x < 0 || y < 0 || tools.board[x][y] != 0) {
    return false
  }
  if (isForbidenMove(tools, x, y)) {
    return false
  }
  return true
}

func isForbidenMove(tools *sdlTools, x int, y int) bool {
  var pion2 int
  var pion int

  if (tools.player == PLAYER1) {
    pion = 1
    pion2 = 2
  } else {
    pion = 2
    pion2 = 1
  }

  var count = 0
  var count2 = 0
  var i = -1
  var j = -1

  for i <= 1 {
    j = -1
    for j <= 1 {
      if (x + i * -1 >= 0 && x + i * -1 < SIZE && y + j * -1 < SIZE && y + j * -1 >= 0 && x + i * 3 >= 0 && x + i * 3 < SIZE && y + j * 3 >= 0 && y + j * 3 < SIZE && tools.board[x + i * -1][y + j * -1] != pion2 && tools.board[x + i][y + j] == pion && tools.board[x + i * 2][y + j * 2] == pion && tools.board[x + i * 3][y + j * 3] != pion2) {
        if (x + i * -2 >= 0 && x + i * -2 < SIZE && y + j * -2 < SIZE && y + j * -2 >= 0 && tools.board[x + i * -2][y + j * -2] == pion && tools.board[x + i * -1][y + j * -1] == pion) {
          count--
        }
        count++
      }
      if (x + i * -1 >= 0 && x + i * -1 < SIZE && y + j * -1 < SIZE && y + j * -1 >= 0 && x + i * 4 >= 0 && x + i * 4 < SIZE && y + j * 4 >= 0 && y + j * 4 < SIZE && tools.board[x + i * -1][y + j * -1] != pion2 && tools.board[x + i][y + j] == 0 && tools.board[x + i * 2][y + j * 2] == pion && tools.board[x + i * 3][y + j * 3] == pion && tools.board[x + i * 4][y + j * 4] != pion2) {
        count2++
      }
      j++
    }
    i++
  }
  if (count >= 2) {
    return true
  } else if (count >= 1 && count2 >= 1) {
    return true
  }
  return false
}
