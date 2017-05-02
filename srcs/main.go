package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

const PUSHBACK = 17
const PUSHBACKB = 28
const AVREGEREPAIR = 6
const W = 1000
const H = 1000

const PLAYER1 = 1
const PLAYER2 = 2
const DEPTH_MAX = 3
const EQUAL = 0
const MIN_BASE = 1000000
const MAX_BASE = -1000000
const AMP = 2
const SIZE = 19

type sdlTools struct {
	win *sdl.Window
	surface *sdl.Surface
	turn bool
	winPlayer1 bool
	winPlayer2 bool
	board [SIZE][SIZE]int
	exit bool
}

func intSdlTools() *sdlTools  {
	var tools *sdlTools = new(sdlTools)
	tools.win = nil
	tools.surface = nil
	tools.turn = false
	tools.winPlayer1 = false
	tools.winPlayer2 = false
	tools.board = newBoard()
	tools.exit = false
	return tools

}
func main() {
	initSdl(intSdlTools())
	return
}
