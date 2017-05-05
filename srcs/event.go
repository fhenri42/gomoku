package main

import (
	"time"
	"github.com/veandco/go-sdl2/sdl"
)

func play(tools *sdlTools, i int, j int) {
	var isEnd bool

  tools.board, isEnd = moveAndEat(tools.board, i, j, tools.player, &(tools.scorePlayer2), &(tools.scorePlayer1))
	tools.player = tools.player % 2 + 1
  printBoard(tools)
	if (isEnd) {
		playAgain(tools, tools.player % 2 + 1)
		initSdlTools(tools)
	}
}

func iaTurn(tools *sdlTools) {
	tools.wait = true
	timeBfore := time.Now()
	bestMove := getNextMove(tools.board, tools.scorePlayer1, tools.scorePlayer2, tools.player, 0)
	timeAfter := time.Now()
	tools.time = timeAfter.Sub(timeBfore)
	play(tools, bestMove.x, bestMove.y)
	displayTime(tools)
	tools.wait = false
}

func displayHint(tools *sdlTools) {
	tools.wait = true
	bestMove := getNextMove(tools.board, tools.scorePlayer1, tools.scorePlayer2, tools.player, 0)
	tools.board[bestMove.x][bestMove.y] = HINT
	printBoard(tools)
	tools.board[bestMove.x][bestMove.y] = 0
	tools.wait = false
}

func  onClic(t *sdl.MouseButtonEvent, tools *sdlTools)  {
	var j int = (int(t.X) + SQUARE / 2 - OFFSET_X) / (SQUARE + SPACING)
	var moduloj = (int(t.X) + SQUARE / 2 - OFFSET_X) % (SQUARE + SPACING)
	var i int = (int(t.Y) + SQUARE / 2 - OFFSET_Y) / (SQUARE + SPACING)
	var moduloi = (int(t.Y) + SQUARE / 2 - OFFSET_Y) % (SQUARE + SPACING)

	if (moduloj > 0 && moduloi > 0 && isPlayable(tools, i, j)) {
    play(tools, i, j)
		if (tools.gameType == SOLO) {
			iaTurn(tools)
		}
		if (tools.gameType != MENU) {
			displayHint(tools)
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
				if t.Type == 1025 && tools.gameType != MENU && !tools.wait {
					onClic(t, tools)
					break
				} else if (t.Type == 1025 && t.X  <= 740 && t.Y  <= 625 && t.Y  >= 492 && t.X >= 466 && tools.gameType == MENU) {
					tools.gameType = 2
					loadMap(tools, "ressources/board.bmp")
				} else if (t.Type == 1025 && t.X  <= 740 && t.Y  <= 822 && t.Y  >= 688 && t.X >= 466 && tools.gameType == MENU) {
					tools.gameType = 1
					loadMap(tools, "ressources/board.bmp")
					if (tools.iaStart) {
						bestMove := getNextMove(tools.board, tools.scorePlayer1, tools.scorePlayer2, tools.player, 0)
						play(tools, bestMove.x, bestMove.y)
					}

				} else if (t.Type == 1025 && t.X  <= 510 && t.Y  <= 921 && t.Y  >= 898 && t.X >= 485 && tools.gameType == MENU) {
					if (tools.iaStart) {
						loadMenu(tools, "ressources/menu.bmp")
					} else {
						loadMenu(tools, "ressources/menu1.bmp")
					}
					tools.iaStart = !tools.iaStart
			 }
			break
			}
		}
	}
}
