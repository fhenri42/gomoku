package main
import (
  "sync"
  "fmt"
)

var k = 0

func  simulateAndGetValue(coup *move, aiScore int, playerScore int, board *[SIZE][SIZE]int, player int, depth int, wg *sync.WaitGroup)  {
  defer (*wg).Done()
   tmpBoard, isEnd := moveAndEat(*board, coup.x, coup.y, player, &aiScore, &playerScore)
   if isEnd {
     (*coup).poid = getNextMove(tmpBoard, aiScore, playerScore, player, DEPTH_MAX).poid
   } else {
     (*coup).poid = getNextMove(tmpBoard, aiScore, playerScore, player, depth + 1).poid
   }
}

func getNextMove(board [SIZE][SIZE]int, aiScore int, playerScore int, player int, depth int) move {
  var t int = 0
  if (depth == 0) {
    k = 0
  }

  if (depth == DEPTH_MAX) {
    fmt.Println(k)
    k++
    return evaluate(board, aiScore, playerScore)
  }

  coups := findMoves(board)
  var wg sync.WaitGroup

  wg.Add(len(coups))
  for t < len(coups) {
    simulateAndGetValue(&coups[t], aiScore, playerScore, &board, player, depth, &wg)
    t++
  }
  wg.Wait()

  if (depth % 2 == 1) {
    return minCoup(coups)
  } else {
    return maxCoup(coups)
  }
}

func evaluate(board [SIZE][SIZE]int, aiScore int, playerScore int) move {
  return newMove(0, 0, 10)
  //return findWeight(board, aiScore, playerScore)
}
