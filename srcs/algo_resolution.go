package main
import (
  "sync"
  "fmt"
  "math"
)

func  simulateAndGetValue(game *Game, move *Move, wg *sync.WaitGroup)  {
  defer wg.Done()
  newGame := copyGame(game)
  //fmt.Print(game)
  moveAndEat(newGame, move.x, move.y)
  if newGame.winner != NONE {
    fmt.Print(newGame)
    if (newGame.winner == newGame.friend) {
      move.poid = MIN_BASE - newGame.depth
    } else {
      move.poid = MAX_BASE + newGame.depth
    }
  } else {
    move.poid = getNextMove(newGame).poid
  }
}

func getNextMove(game *Game) Move {
  var t int = 0

  if (game.depth == DEPTH_MAX) {
    return evaluate(game)
  }

  moves := findMoves(&game.board)
  if (len(moves) == 0) {
    return newMove(9, 9, 0)
  }
  fmt.Println(moves)
  var wg sync.WaitGroup
  wg.Add(len(moves))
  for t < len(moves) {
    simulateAndGetValue(game, &moves[t], &wg)
    t++
  }
  wg.Wait()

  if (game.curPlayer == game.friend) {
    return maxCoup(moves)
  } else {
    return minCoup(moves)
  }
}

func evaluate(game *Game) Move {
  var res int = 0
  var i int = 0
  var j int = 0

  for i < game.score[game.friend - 1] {
    res += i + 1 * 10
    i++
  }
  i = 0
  for i < game.score[game.friend % 2] {
    res -= i + 1 * 10
    i++
  }
  i = 0
  for i < SIZE {
    j = 0
    for j < SIZE {
      if (game.board[i][j] == game.friend) {
        res += evaluateEach(&game.board, i, j)
      } else if (game.board[i][j] == game.friend % 2 + 1) {
        res -= evaluateEach(&game.board, i, j)
      }
      j++
    }
    i++
  }
  return (newMove(0, 0, res))
  //return newMove(0, 0, findWeight(board, aiScore, playerScore))
}

func evaluateEach(board *[SIZE][SIZE]int, x int, y int) int {
  var i int = -1
  var j int = -1
  var k int = 1
  var player = board[x][y]
  var res int = 0
  var count int = 0

  for i <= 1 {
    j = -1
    for j <= 1 {
      k = 1
      count = 0
      for k <= 4 {
        if (k == 0 || (x + i * k >= 0 && x + i * k < SIZE && y + j * k >= 0 && y + j * k < SIZE && (*board)[x + i * k][y + j * k] == player)) {
          count++
        }
        k++
      }
      j++
    }
    i++
  }

  i = 0
  for i < count {
    res += int(math.Pow(float64(i), 10))
    i++
  }
  return res
}
