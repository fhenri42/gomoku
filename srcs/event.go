package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
)

func play(tools *sdlTools, i int, j int) {
  tools.board = moveAndEat(tools.board, i, j, tools.player, &(tools.scorePlayer2), &(tools.scorePlayer1))
  isEnd, _ := endGame(tools.board)
  if (isEnd) {
    fmt.Println("THE END")
  }
  tools.player = tools.player % 2 + 1
  printBoard(tools)
}

func  onClic(t *sdl.MouseButtonEvent, tools *sdlTools)  {
	var j int = (int(t.X) + SQUARE / 2 - OFFSET) / SQUARE
	var i int = (int(t.Y) + SQUARE / 2 - OFFSET) / SQUARE

	if (isPlayable(tools.board, i, j)) {
    play(tools, i, j)
		if (tools.gameType == SOLO) {
			tools.wait = true
			bestMove := getBestMove(tools.board, tools.scorePlayer1, tools.scorePlayer2)
			play(tools, bestMove.x, bestMove.y)
			tools.wait = false
		}
	}
}

func handleEvent(tools *sdlTools) {

	for !tools.exit {

		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch t := event.(type) {
			case *sdl.QuitEvent:
				tools.exit = true
				break
			case *sdl.MouseMotionEvent:
				//fmt.Println(t.X, " ", t.Y)
				break
			case *sdl.KeyDownEvent:
				if t.Keysym.Sym == 27 {
					tools.exit = true
				}
				break
			case *sdl.MouseButtonEvent:
				if t.Type == 1025 && tools.gameState && !tools.wait {
					onClic(t, tools)
					break
				} else if (t.Type == 1025 && t.X  <= 680 && t.Y  <= 460 && t.Y  >= 240 && t.X >= 380 && !tools.gameState) {
					tools.gameState = true
					tools.gameType = 2
					loadMap(tools)
				} else if (t.Type == 1025 && t.X  <= 680 && t.Y  <= 719 && t.Y  >= 501 && t.X >= 380 && !tools.gameState) {
					tools.gameState = true
					tools.gameType = 1
					loadMap(tools)
				}
			break
			}
		}
	}
}
