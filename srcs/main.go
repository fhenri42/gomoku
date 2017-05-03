package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"fmt"
)

type move struct {
  x int
  y int
  poid int
}

const W = 1205
const H = 1205
const OFFSET_Y = 160
const OFFSET_X = 110
const SPACING = 5
const PLAYER1 = 1
const PLAYER2 = 2
const DEPTH_MAX = 3
const EQUAL = 0
const MIN_BASE = 1000000
const MAX_BASE = -1000000
const AMP = 2
const SIZE = 19
const SQUARE = 50
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
	gameState bool
	gameType int
	wait bool
}

func intSdlTools() *sdlTools  {
	var tools *sdlTools = new(sdlTools)
	tools.win = nil
	tools.surface = nil
	tools.player = PLAYER1
	tools.scorePlayer1 = 0
	tools.scorePlayer2 = 0
	tools.board = newBoard()
	tools.exit = false
	tools.gameState = false
	tools.gameType = 0
	tools.wait = false
	return tools
}

func initSdl(tools *sdlTools) {

	tools.win,_ = sdl.CreateWindow("Gomokou", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, W, H, sdl.WINDOW_SHOWN)
	tools.surface, _ = tools.win.GetSurface()

	errCenter := putImageCenter("ressources/menu.bmp", tools.surface)
	if errCenter != nil {
		fmt.Println("err imageCenter", errCenter)
	}

	tools.win.UpdateSurface()
	handleEvent(tools)
	sdl.Quit();

	return;
}

func main() {
initSdl(intSdlTools())
	return
}
