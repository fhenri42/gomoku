package main

import ("fmt")


func findScoreLine(board[SIZE][SIZE]int, numberAligne int ,pion int) {
var y = 0

for y < SIZE {
  var x = 0
    for x < SIZE {
      var t = 0
      var continu = true
      for t <= numberAligne  && continu {
        if y + numberAligne <= SIZE && board[y][x] === pion && board[y + t][x] === pion {
          t++
          } else {
            continu = false
          }
        if t == numberAligne {
          return numberAligne * 100
        }
      }
      t = 0
      continu = true
      for t <= numberAligne  && continu {
        if (x + 4 <= SIZE && board[y][x] === pion  && board[y][x + t] === pion) {
          t++
          } else {
            continu = false
          }
        if t == numberAligne {
          return numberAligne * 100
        }
      }

      //
      // if (x + 4 <= SIZE && board[y][x] === pion && board[y][x + 1] === pion && board[y][x + 2] === pion && board[y][x + 3] === pion && board[y][x + 4] === pion) {
      //   return true
      // }
      // if (x + 4 <= SIZE && y + 4 <= SIZE && board[y][x] === pion && board[y + 1][x + 1] === pion && board[y + 2][x + 2] === pion && board[y + 3][x + 3] === pion && board[y + 4][x + 4] === pion) {
      //   return true
      // }
      // if ( x - 4 >= 0 && y + 4 <= SIZE && board[y][x] === pion && board[y + 1][x - 1] === pion && board[y + 2][x - 2] === pion && board[y + 3][x - 3] === pion && board[y + 4][x - 4] === pion)  {
      //   return true
      // }
    }
  }
}

func findEatMove(board [SIZE][SIZE]) {

}
func findWeight(board [SIZE][SIZE]int, aiScore int, playerScore int) int {

    if aiScore == 10 {
      return 2000
    }
    if playerScore == 10 {
      return -2000
    }
    if findScoreLine(board, numberAligne, 2) {
      return 1800
    }
    if findScoreLine(board, 1) {
      return -1800
    }
    if findEatMove(board) {
      return 1700
    }
    return 10
}



/*
1) cherher a manger un pyon si on est a 9 sur 10                    Poids: 2000
2) cherher un allignmen de 5 sans possibiliter detre contrer        Poids: SIZE00
3) cherher a contrer un alignments de 5 du joueur                   Poids: 1800
4) cherher a manger un pion                                         Poids: 1700
a ajouter

*/
/*
CORRECTION :
1) current alignments into account
2) check whether an alignment has enough space to develop into a 5-in-a-row
3) heuristic weigh an alignment according to its freedom
4) potential captures into account
5) heuristic take current captured stones into account
6) heuristic check for advanteageous combinations
7) heuristic take both players into account
HARDCORE
8)heuristic take past player actions into account to identify patterns and weigh board boards accordingl
*/
