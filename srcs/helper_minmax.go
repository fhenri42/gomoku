package main

func minCoup(coups []move) move {
  var min = MIN_BASE
  var minCoup *move
  var t = 0

  for t < len(coups) {
    if (coups[t].poid < min) {
      min = coups[t].poid
      minCoup = &coups[t]
    }
    t++
  }
  return *minCoup
}

func maxCoup(coups []move) move {
  var max = MAX_BASE
  var maxCoup *move
  var t = 0

  for t < len(coups) {
    if (coups[t].poid > max) {
      max = coups[t].poid
      maxCoup = &coups[t]
    }
    t++
  }
  return *maxCoup
}

func newMove(x int, y int, poid int) move {
  var tmp move

  tmp.x = x
  tmp.y = y
  tmp.poid = poid
  return tmp
}
