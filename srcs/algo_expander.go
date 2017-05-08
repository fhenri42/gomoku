package main

import (
  //"fmt"
)

// Check if there is a pion around the move we attempt to play
func isClose(x int, y int, board *[SIZE][SIZE]int) bool {
  if (board[x][y] != 0) {
    return false
  }
  var i = -AMP
  for i <= AMP {
    var j = -AMP
    for j <= AMP {
      if (x + i >= 0 && x + i < SIZE && y + j >= 0 && y + j < SIZE && board[x + i][y + j] != 0) {
        return true
      }
      j++
    }
    i++
  }
  return false
}

// Get array of all playable moves around
func findMoves(board *[SIZE][SIZE]int) ([]Move) {

  moves := make([]Move, 0)
  var x = 0
  for x < SIZE {
    var y = 0
    for y < SIZE {
      if (isClose(x, y, board)) {
        moves = append(moves, newMove(x, y, 0))
      }
      y++
    }
    x++
  }
  if (len(moves) == 0) {
    moves = append(moves, newMove(9, 9, 0))
  }
  return moves
}


  // type afluence struct {
  //   conteur int
  //   arrayMove Move
  // }
  //
  //
  // func isClose(x int, y int, board *[SIZE][SIZE]int) afluence {
  //
  //
  //   var conteur int
  //   conteur = 0
  //   var i int = -AMP
  //   for i <= AMP {
  //     var j int = -AMP
  //     for j <= AMP {
  //       if (x + i >= 0 && x + i < SIZE && y + j >= 0 && y + j < SIZE && board[x + i][y + j] != 0) {
  //         conteur++
  //       }
  //       j++
  //     }
  //     i++
  //   }
  //   var tmp afluence
  //   tmp.conteur = conteur
  //   tmp.arrayMove = newMove(x, y, 0)
  //   return tmp
  // }
  //
  // func findMoves(board *[SIZE][SIZE]int) ([]Move) {
  //
  //   arrayAflu := make([]afluence, 0)
  //
  //   var x int = 0
  //   for x < SIZE {
  //     var y int = 0
  //     for y < SIZE {
  //       if board[x][y] == 0 {
  //         arrayAflu = append(arrayAflu, isClose(x,y, board))
  //        }
  //       y++
  //     }
  //     x++
  //   }
  //
  //   tmpBoard := make([]Move, 0)
  //   var tmpConter int = 0
  //   var t int = 0
  //   for t < len(arrayAflu) {
  //     if (tmpConter < arrayAflu[t].conteur) {
  //       tmpConter = arrayAflu[t].conteur
  //     }
  //       t++
  //     }
  //     t = 0
  //     for t < len(arrayAflu) {
  //       if tmpConter == arrayAflu[t].conteur {
  //         tmpBoard = append(tmpBoard, arrayAflu[t].arrayMove)
  //       }
  //       t++
  //     }
  //     return tmpBoard
  //   }

// ancinne version
//
