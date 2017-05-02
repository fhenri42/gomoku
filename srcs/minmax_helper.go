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
