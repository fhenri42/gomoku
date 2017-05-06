package main

import (
	"time"
	"github.com/veandco/go-sdl2/sdl"
	//"fmt"
)

func play(game *Game, tools *Tools, i int, j int) {
  moveAndEat(game, i, j)
  displayBoard(tools, game)
}

func iaTurn(tools *Tools, game *Game) time.Duration {
	tools.wait = true
	timeBfore := time.Now()
	bestMove := getNextMove(game)
	timeAfter := time.Now()
	time := timeAfter.Sub(timeBfore)
	play(game, tools, bestMove.x, bestMove.y)
	tools.wait = false
	return time
}

func displayHint(tools *Tools, game *Game) {
	tools.wait = true
	bestMove := getNextMove(game)

	game.board[bestMove.x][bestMove.y] = HINT
	displayBoard(tools, game)
	game.board[bestMove.x][bestMove.y] = NONE
	tools.wait = false
}

func  onClic(t *sdl.MouseButtonEvent, tools *Tools, game *Game)  {
	var j int = (int(t.X) + SQUARE / 2 - OFFSET_X) / (SQUARE + SPACING)
	var moduloj = (int(t.X) + SQUARE / 2 - OFFSET_X) % (SQUARE + SPACING)
	var i int = (int(t.Y) + SQUARE / 2 - OFFSET_Y) / (SQUARE + SPACING)
	var moduloi = (int(t.Y) + SQUARE / 2 - OFFSET_Y) % (SQUARE + SPACING)

	if (moduloj > 0 && moduloi > 0 && isPlayable(game, i, j)) {
    play(game, tools, i, j)
		if (game.winner != NONE) {
			endGame(tools, game)
		} else if (tools.gameType == SOLO) {
			time := iaTurn(tools, game)
			if (game.winner != NONE) {
				endGame(tools, game)
			} else {
				displayHint(tools, game)
			}
			displayTime(tools, time)
		} else {
			displayHint(tools, game)
		}
	}
}

func handleEvent(tools *Tools, game *Game) {

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
					onClic(t, tools, game)
				} else if (t.Type == 1025 && t.X  <= 740 && t.Y  <= 625 && t.Y  >= 492 && t.X >= 466 && tools.gameType == MENU) {
					tools.gameType = MULTI
					displayBoard(tools, game)
				} else if (t.Type == 1025 && t.X  <= 740 && t.Y  <= 822 && t.Y  >= 688 && t.X >= 466 && tools.gameType == MENU) {
					tools.gameType = SOLO
					displayBoard(tools, game)
					if (tools.iaStart) {
						iaTurn(tools, game)
					}
				} else if (t.Type == 1025 && t.X  <= 510 && t.Y  <= 921 && t.Y  >= 898 && t.X >= 485 && tools.gameType == MENU) {
					tools.iaStart = !tools.iaStart
					displayMenu(tools)
				}
				break
			}
		}
	}
}
