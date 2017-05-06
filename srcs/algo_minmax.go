package main

import (
  //"fmt"
)

func minCoup(moves []Move) Move {
  var min = MIN_BASE
  var minCoup *Move
  var t = 0

  //fmt.Println(moves)
  for t < len(moves) {
    if (moves[t].poid < min) {
      min = moves[t].poid
      minCoup = &moves[t]
    }
    t++
  }
  return *minCoup
}

func maxCoup(moves []Move) Move {
  var max = MAX_BASE
  var maxCoup *Move
  var t = 0

  //fmt.Println(moves)
  for t < len(moves) {
    if (moves[t].poid > max) {
      max = moves[t].poid
      maxCoup = &moves[t]
    }
    t++
  }
  return *maxCoup
}

func newMove(x int, y int, poid int) Move {
  var tmp Move

  tmp.x = x
  tmp.y = y
  tmp.poid = poid
  return tmp
}
