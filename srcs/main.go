package main

import ()

type move struct {
  x int
  y int
  poid int
}

const PUSHBACK = 17
const PUSHBACKB = 28
const AVREGEREPAIR = 6
const W = 1000
const H = 1000
var TURN = true

const PLAYER1 = 1
const PLAYER2 = 2
const DEPTH_MAX = 3
const EQUAL = 0
const MIN_BASE = 1000000
const MAX_BASE = -1000000
const AMP = 2
const SIZE = 19

func main() {
	board := newBoard()
	initSdl(&board)
}
