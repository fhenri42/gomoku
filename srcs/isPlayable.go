package main

func isPlayable(tools *sdlTools, x int, y int) bool {
  if (x >= SIZE || y >= SIZE || x < 0 || y < 0 || tools.board[x][y] != 0) {
    return false
  }
  // if (!forbidenMove(tools, x, y)) {
  //   return false
  // }
  return true
}

func forbidenMove(tools *sdlTools, x int, y int) bool {
  var pion2 int
  var pion int

  if (tools.player == PLAYER1) {
    pion = 1
    pion2 = 2
  } else {
    pion = 2
    pion2 = 1
  }

  if (y - 3 >= 0 && x + 3 <= 19 && x - 3 >= 0 && y + 1 <= 19 && tools.board[x + 1][y] == pion &&  tools.board[x + 2][y] == pion && tools.board[x - 1][y - 1] == pion && tools.board[x - 2][y - 2] == pion) {
      if (tools.board[x - 3][y - 3] == pion2 || tools.board[x - 1][y] == pion2 || tools.board[x + 1][y + 1] == pion2 || tools.board[x + 3][y] == pion2) {
        return true
      } else {
        return false
      }
  }

  if (y - 3 >= 0 && x + 3 <= 19 && x - 3  >= 0 && y + 1 <= 19 && tools.board[x - 1][y] == pion &&  tools.board[x - 2][y] == pion && tools.board[x + 1][y - 1] == pion && tools.board[x + 2][y - 2] == pion) {
      if (tools.board[x + 1][y] == pion2 || tools.board[x - 3][y] == pion2 || tools.board[x - 1][y + 1] == pion2 || tools.board[x + 3][y - 3] == pion2) {
        return true
      } else {
        return false
      }
  }

  if ( x + 3 <= 19 && y + 3 <= 19 && y - 2  >= 0 && x - 1 >= 0 && tools.board[x][y + 1] == pion &&  tools.board[x][y + 2] == pion && tools.board[x + 1][y - 1] == pion && tools.board[x + 2][y - 2] == pion) {
    if (tools.board[x][y - 1] == pion2 || tools.board[x + 3][y - 3] == pion2 || tools.board[x - 1][y + 1] == pion2 || tools.board[x][y + 3] == pion2) {
      return true
    } else {
      return false
    }
  }

  if ( y + 3 <= 19 && x + 3 <= 19 && x - 3 >= 0 && y - 1 >= 0 && tools.board[x - 1][y] == pion &&  tools.board[x - 2][y] == pion && tools.board[x + 1][y + 1] == pion && tools.board[x + 2][y + 2] == pion) {
    if (tools.board[x - 3][y] == pion2 || tools.board[x + 1][y] == pion2 || tools.board[x + 1][y - 1] == pion2 || tools.board[x + 3][y + 3] == pion2) {
      return true
    } else {
      return false
    }
  }

  if (y + 3 <= 19 && y - 3 >= 0 && x - 3 >= 0 && x + 1 <= 19 && tools.board[x][y - 1] == pion &&  tools.board[x][y - 2] == pion && tools.board[x - 1][y + 1] == pion && tools.board[x - 2][y + 2] == pion) {
    if (tools.board[x][y + 1] == pion2 || tools.board[x + 1][y - 1] == pion2 || tools.board[x][y - 3] == pion2 || tools.board[x - 3][y + 3] == pion2) {
      return true
    } else {
      return false
    }
  }

  if (y - 3 >= 0 && y + 3 <= 19 && x + 3 <= 19 && x - 1 >= 0 && tools.board[x][y - 1] == pion &&  tools.board[x][y - 2] == pion && tools.board[x + 1][y + 1] == pion && tools.board[x + 2][y + 2] == pion) {
    if (tools.board[x - 1][y - 1] == pion2 || tools.board[x][y - 3] == pion2 || tools.board[x][y + 1] == pion2 || tools.board[x + 3][y + 3] == pion2) {
      return true
    } else {
      return false
    }
  }

  if (y + 3 <= 19 && y - 3 >= 0 && x + 3 <= 19 && x - 1 >= 0 && tools.board[x][y + 1] == pion &&  tools.board[x][y + 2] == pion && tools.board[x + 1][y - 1] == pion && tools.board[x + 2][y - 2] == pion) {
    if (tools.board[x][y - 1] == pion2 || tools.board[x - 1][y + 1] == pion2 || tools.board[x][y + 3] == pion2 || tools.board[x + 3][y - 3] == pion2) {
      return true
    } else {
      return false
    }
  }

  if (y + 3 <= 19 && y - 3 >= 0 && x - 3 >= 0 && x + 1 <= 19 && tools.board[x][y + 1] == pion &&  tools.board[x][y + 2] == pion && tools.board[x - 1][y - 1] == pion && tools.board[x - 2][y - 2] == pion) {
    if (tools.board[x][y - 1] == pion2 || tools.board[x][y + 3] == pion2 || tools.board[x + 1][y + 1] == pion2 || tools.board[x - 3][y - 3] == pion2) {
      return true
    } else {
      return false
    }
  }

  if (y + 3 <= 19 && y - 1 >= 0 && x + 3 <= 19 && x - 3 >= 0 && tools.board[x + 1][y + 1] == pion &&  tools.board[x + 2][y + 2] == pion && tools.board[x - 1][y + 1] == pion && tools.board[x - 2][y + 2] == pion) {
    if (tools.board[x - 1][y - 1] == pion2 || tools.board[x + 1][y - 1] == pion2 || tools.board[x - 3][y + 3] == pion2 || tools.board[x + 3][y + 3] == pion2) {
      return true
    } else {
      return false
    }
  }

  if (y + 3 <= 19 && y - 3 >= 0 && x + 1 <= 19 && x - 3 >= 0 && tools.board[x - 1][y + 1] == pion &&  tools.board[x - 2][y + 2] == pion && tools.board[x - 1][y - 1] == pion && tools.board[x - 2][y - 2] == pion) {
    if (tools.board[x + 1][y - 1] == pion2 || tools.board[x + 1][y + 1] == pion2 || tools.board[x - 3][y + 3] == pion2 || tools.board[x - 3][y - 3] == pion2) {
      return true
    } else {
      return false
    }
  }

  if (y + 1 <= 19 && y - 3 >= 0 && x + 3 <= 19 && x - 3 >= 0 && tools.board[x - 1][y - 1] == pion &&  tools.board[x - 2][y - 2] == pion && tools.board[x + 1][y - 1] == pion && tools.board[x + 2][y - 2] == pion) {
    if (tools.board[x + 1][y + 1] == pion2 || tools.board[x - 1][y + 1] == pion2 || tools.board[x - 3][y - 3] == pion2 || tools.board[x + 3][y - 3] == pion2) {
      return true
    } else {
      return false
    }
  }

  if (y + 3 <= 19 && y - 3 >= 0 && x + 3 <= 19 && x - 1 >= 0 && tools.board[x + 1][y + 1] == pion &&  tools.board[x + 2][y + 2] == pion && tools.board[x + 1][y - 1] == pion && tools.board[x + 2][y - 2] == pion) {
    if (tools.board[x - 1][y - 1] == pion2 || tools.board[x - 1][y + 1] == pion2 || tools.board[x + 3][y + 3] == pion2 || tools.board[x + 3][y - 3] == pion2) {
      return true
    } else {
      return false
    }
  }

  return true
}
