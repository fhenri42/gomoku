package main
import (
	//"fmt"
)
func moveAndEat(game *Game, x int, y int) {
  otherPlayer := game.curPlayer % 2 + 1
  var prevLastChance *LastChance = game.lastChance
  game.lastChance = nil
  var i = -1
  var j = -1
  var k = -4
  var count = 0
  var countMax = 0
  var posi = -1
  var posj = -1

  for i <= 1 {
    j = -1
    for j <= 1 {
      k = -4
      count = 0
      countMax = 0
      for k <= 4 {
        if (k == 0 || (x + i * k >= 0 && x + i * k < SIZE && y + j * k >= 0 && y + j * k < SIZE && game.board[x + i * k][y + j * k] == game.curPlayer)) {
          count++
          if (count == 5) {
            posi = x + i * k
            posj = y + j * k
          }
          countMax = count
        } else {
          count = 0
        }
        k++
      }
      if (countMax >= 5) {
        if (isUnbreakableLine(-i, -j, posi, posj, &game.board)) {
          game.winner = game.curPlayer
        } else {
          game.lastChance = new(LastChance)
					game.lastChance.player = game.curPlayer % 2 + 1
					game.lastChance.winner = game.curPlayer
					game.lastChance.i = -i
					game.lastChance.j = -j
					game.lastChance.x = posi
					game.lastChance.y = posj
      	}
      }
      if (x + i * 3 >= 0 && x + i * 3 < SIZE && y + j * 3 >= 0 && y + j * 3 < SIZE && game.board[x + i][y + j] == otherPlayer && game.board[x + i * 2][y + j * 2] == otherPlayer && game.board[x + i * 3][y + j * 3] == game.curPlayer) {
        game.board[x + i][y + j] = 0
        game.board[x + i * 2][y + j * 2] = 0
        game.score[game.curPlayer - 1] += 2
      }
      j++
    }
    i++
  }
  game.board[x][y] = game.curPlayer
  if (game.score[PLAYER2 - 1] == 10) {
		game.winner = PLAYER2
	} else if (game.score[PLAYER1 - 1] == 10) {
    game.winner = PLAYER1
  }

  if (prevLastChance != nil && stillStanding(prevLastChance, game)) {
    game.winner = prevLastChance.winner
  }
  game.curPlayer = game.curPlayer % 2 + 1
  if (game.depth == 0) {
    game.friend = game.curPlayer
  }
}

func stillStanding(lastChance *LastChance, game *Game) bool {
  var i int = 0

  for i < 5 {
    if (game.board[lastChance.x + lastChance.i * i][lastChance.y + lastChance.j * i] != lastChance.winner) {
      return false
    }
    i++
  }
  return true
}
