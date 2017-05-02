package main

func minCoup(coups []move) *move {
  var min = MIN_BASE
  var minCoup *move

  for _,coup := range coups {
    if (coup.poid < min) {
      min = coup.poid
      minCoup = &coup
    }
  }
  return minCoup
}

func maxCoup(coups []move) *move {
  var max = MAX_BASE
  var maxCoup *move

  for index,coup := range coups {
    if (coup.poid > max) {
      max = coups[index].poid
      maxCoup = &coups[index]
    }
  }
  return maxCoup
}

func endGame(state [][]int) (bool, int) {
  var x = 0
  for x < SIZE {
    var y = 0
    for y < SIZE {
      if (isUnbreakableLine(x, y, state)) {
        return true, state[x][y]
      }
      y++
    }
    x++
  }
  return false, 0
}

func isUnbreakableLine(i int, j int, state [][]int) bool {
  var pion = state[i][j]

  if (pion == 0) {
    return false
  }

  if i + 4 <= SIZE && state[i][j] == pion && state[i + 1][j] == pion && state[i + 2][j] == pion && state[i + 3][j] == pion && state[i + 4][j] == pion {
    // && !eatable(state, i, j)
    // && !eatable(state, i + 1, j)
    // && !eatable(state, i + 2, j)
    // && !eatable(state, i + 3, j)
    // && !eatable(state, i + 4, j)) {
    return true

  } else if j + 4 <= SIZE && state[i][j] == pion && state[i][j + 1] == pion && state[i][j + 2] == pion && state[i][j + 3] == pion && state[i][j + 4] == pion {
    // && !eatable(state, i, j + 1)
    // && !eatable(state, i, j + 2)
    // && !eatable(state, i, j + 3)
    // && !eatable(state, i, j + 4)
    // && !eatable(state, i, j + 5)) {
    return true

  } else if j + 4 <= SIZE && i + 4 <= SIZE && state[i][j] == pion && state[i + 1][j + 1] == pion && state[i + 2][j + 2] == pion && state[i + 3][j + 3] == pion && state[i + 4][j + 4] == pion {
    // && !eatable(state, i, j)
    // && !eatable(state, i + 1, j + 1)
    // && !eatable(state, i + 2, j + 2)
    // && !eatable(state, i + 3, j + 3)
    // && !eatable(state, i + 4, j + 4)) {
    return true

  } else if j - 4 >= 0 && i + 4 <= SIZE && state[i][j] == pion && state[i + 1][j - 1] == pion && state[i + 2][j - 2] == pion && state[i + 3][j - 3] == pion && state[i + 4][j - 4] == pion {
    // && !eatable(state, i, j)
    // && !eatable(state, i + 1, j - 1)
    // && !eatable(state, i + 2, j - 2)
    // && !eatable(state, i + 3, j - 3)
    // && !eatable(state, i + 4, j - 4)) {
    return true
  }
  return false
}

// func eatable() bool {
//
// }
