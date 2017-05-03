package main

func isPlayable(tools *sdlTools, x int, y int) bool {
  if (x >= SIZE || y >= SIZE || x < 0 || y < 0 || tools.board[x][y] != 0) {
    return false
  }
  if (!isForbidenMove(tools, x, y)) {
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
      if (x + i * 3 > 0 && x + i * 3 < SIZE && y + j * 3 > 0 && y + j * 3 < SIZE && tools.board[x + i * -1][y + j * -1] != pion2 && tools.board[x + i][y + j] == pion && tools.board[x + i * 2][y + j * 2] == pion && tools.board[x + i * 3][y + j * 3] != pion2) {
        if (tools.board[x + i * -2][y + j * -2] == pion && tools.board[x + i * -1][y + j * -1] == pion) {
          count--
        }
        count++
      }
      if (x + i * 4 > 0 && x + i * 4 < SIZE && y + j * 4 > 0 && y + j * 4 < SIZE && tools.board[x + i * -1][y + j * -1] != pion2 && tools.board[x + i][y + j] == 0 && tools.board[x + i * 2][y + j * 2] == pion && tools.board[x + i * 3][y + j * 3] == pion && tools.board[x + i * 4][y + j * 4] != pion2) {
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
