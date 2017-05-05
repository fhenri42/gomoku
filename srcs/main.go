package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"time"
)

type move struct {
  x int
  y int
  poid int
}

const W = 1205
const H = 1205
const OFFSET_Y = 160
const OFFSET_SCORE_Y = 70
const SPACING_SCORE = 2
const OFFSET_X = 110
const SPACING = 5
const SQUARE = 50
const PION = 40
const OFFSET_LAST_SCORE_X = 574
const OFFSET_LAST_SCORE_Y = 34
const OFFSET_ARRAY_Y = 32
const OFFSET_ARRAY_LEFT_X = 275
const OFFSET_ARRAY_RIGHT_X = 860

const PLAYER1 = 1
const PLAYER2 = 2
const HINT = 3
const DEPTH_MAX = 5
const EQUAL = 0
const MIN_BASE = 1000000
const MAX_BASE = -1000000
const AMP = 1
const SIZE = 19
const MULTI = 2
const SOLO = 1

type sdlTools struct {
	win *sdl.Window
	surface *sdl.Surface
	player int
	scorePlayer1 int
	scorePlayer2 int
	board [SIZE][SIZE]int
	exit bool
	iaStart bool
	gameState bool
	gameType int
	wait bool
	time time.Duration
	hint *move
}


func initSdlTools(tools *sdlTools) {

	tools.iaStart = false
	tools.player = PLAYER1
	tools.scorePlayer1 = 0
	tools.scorePlayer2 = 0
	tools.board = newBoard()
	tools.exit = false
	tools.gameState = false
	tools.gameType = 0
	tools.wait = false
	tools.hint = nil
}

func initSdl(tools *sdlTools) {

	tools.win,_ = sdl.CreateWindow("Gomokou", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, W, H, sdl.WINDOW_SHOWN)
	tools.surface, _ = tools.win.GetSurface()
	loadMenu(tools, "ressources/menu.bmp")
	handleEvent(tools)
	sdl.Quit()
	return
}

func main() {
	var tools *sdlTools = new(sdlTools)
	initSdlTools(tools)
	initSdl(tools)
	return
}
