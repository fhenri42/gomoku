package main
import (
  "github.com/veandco/go-sdl2/sdl"
)

func copyLastChance(lastChance *LastChance) *LastChance {
  if (lastChance == nil) {
    return nil
  }
  var newLastChance *LastChance = new(LastChance)
  newLastChance.i = lastChance.i
  newLastChance.j = lastChance.j
  newLastChance.x = lastChance.x
  newLastChance.y = lastChance.y
  newLastChance.winner = lastChance.winner
  newLastChance.player = lastChance.player
  return newLastChance
}

func copyGame(game *Game) *Game {
  var newGame *Game = new(Game)
  newGame.score = game.score
  newGame.lastChance = copyLastChance(game.lastChance)
  newGame.curPlayer = game.curPlayer
  newGame.board = game.board
  newGame.depth = game.depth + 1
  newGame.winner = game.winner
  newGame.friend = game.friend
  return newGame
}

func newGame(game *Game) {
	game.score[0] = 0
	game.score[1] = 0
	game.lastChance = nil
	game.curPlayer = PLAYER1
	game.friend = PLAYER1
	game.board = newBoard()
	game.depth = 0
	game.winner = 0
}

func initGame() *Game {
	var game *Game = new(Game)

	game.score[0] = 0
	game.score[1] = 0
	game.lastChance = nil
	game.curPlayer = PLAYER1
	game.friend = PLAYER1
	game.board = newBoard()
	game.depth = 0
	game.winner = 0
	return game
}

func newBoard() [SIZE][SIZE]int {
	var board [SIZE][SIZE]int
	var x = 0
	for x < SIZE {
		var y = 0
		for y < SIZE {
			board[x][y] = 0
			y++
		}
		x++
	}
	return board
}

func initTools() *Tools {
	var tools *Tools = new(Tools)

	tools.win,_ = sdl.CreateWindow("Gomoku", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, W, H, sdl.WINDOW_SHOWN)
	tools.surface, _ = tools.win.GetSurface()
	tools.iaStart = false
	tools.gameType = MENU
	tools.wait = false
	return tools
}
