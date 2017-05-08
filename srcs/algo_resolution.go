package main
import (
  //"fmt"
)

func getNextMove(game *Game, alpha int, beta int) Move {
  var t int = 0
  var shouldPrune bool = false

  // Reaching leaves
  if (game.depth == DEPTH_MAX) {
    return evaluate(game)
  }

  // Finding all possible moves around
  moves := findMoves(&game.board)

  // If only one possibility, no choice, no need to calculate
  if (len(moves) == 1) {
    return moves[0]
  }

  // Go through the array
  for t < len(moves) {

    // If the rest of the nodes are considered as useless, we consider the node as pruned
    if (shouldPrune) {
      moves[t].pruned = true

    // Else we Simulate the move for this node
    } else {
      newGame := copyGame(game)
      moveAndEat(newGame, moves[t].x, moves[t].y)

      // Check if it's a win node
      if newGame.winner == newGame.friend {
        moves[t].poid = INFINI - newGame.depth
      } else if newGame.winner == newGame.friend % 2 + 1 {
        moves[t].poid = -INFINI + newGame.depth

      // reiterate with this function with the current node
      } else {
        moves[t].poid = getNextMove(newGame, alpha, beta).poid
      }

      // We update Alpha || Beta if result from the node match the conditions
      if (game.curPlayer != game.friend && moves[t].poid < beta) {
        beta = moves[t].poid
      } else if (game.curPlayer == game.friend && moves[t].poid > alpha) {
        alpha = moves[t].poid
      }

      // If Alpha > beta, the rest of the nodes are going to be pruned
      if (alpha >= beta) {
        shouldPrune = true
      }
    }
    t++
  }

  // Once all child nodes have been calculated, we can use min / max and return the right one
  if (game.curPlayer == game.friend) {
    return maxCoup(moves)
  } else {
    return minCoup(moves)
  }
}
