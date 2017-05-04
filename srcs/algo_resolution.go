package main
import (
  "fmt"
  "sync"
)
var FIND_COUP_ROUTINE sync.WaitGroup
var CALCULE_VALUE_ROUTINE sync.WaitGroup



func  findCoup(index int, coups[]move,coup *move, aiScore int, playerScore int, board*[SIZE][SIZE]int, player int)  {


  var tmpBoard [SIZE][SIZE]int
  var tmpPlayerScore int
  var tmpAiScore int

  CALCULE_VALUE_ROUTINE.Add(DEPTH_MAX)

  tmpAiScore, tmpPlayerScore = aiScore, playerScore
  tmpBoard = moveAndEat(*board, coup.x, coup.y, player, &tmpAiScore, &tmpPlayerScore)
  //coups[index].poid = go calcValue(tmpBoard, 1, tmpAiScore, tmpPlayerScore, player, *coups[index].poid)
  go calcValue(tmpBoard, 1, tmpAiScore, tmpPlayerScore, player, &coups[index].poid)


  CALCULE_VALUE_ROUTINE.Wait()

  fmt.Println("this is takin time like hell lot of time",coup)
  defer FIND_COUP_ROUTINE.Done()

}

func getBestMove(board [SIZE][SIZE]int, aiScore int, playerScore int, player int) move {
  coups := findMoves(board) // REMPLI UN ARRAY AVEC LES COUPS POTENTIELS

  FIND_COUP_ROUTINE.Add(len(coups))
  var t int = 0
  for t < len(coups) {
    go findCoup(t, coups ,&coups[t], aiScore, playerScore, &board, player)
    t++
    fmt.Println("go lunch a routine in BESTMOVE")
  }
  FIND_COUP_ROUTINE.Wait()
  fmt.Println("All routine are done the code continue")
  if len(coups) != 0 {
    bestCoup := maxCoup(coups)
    return *bestCoup
  }
  var tmp move
  tmp.x = 9
  tmp.y = 9
  return tmp
}

func calcValue(board [SIZE][SIZE]int, depth int, aiScore int, playerScore int, player int, test* int) int {
  defer CALCULE_VALUE_ROUTINE.Done()
  end, winner := endGame(board, playerScore, aiScore) // OBSERVE LA GRILLE POUR SAVOIR SI C'EST LA FIN
  var tmpBoard [SIZE][SIZE]int
  var tmpPlayerScore int
  var tmpAiScore int

  player = player % 2 + 1

  if (end) { // SI FIN DU GAME
    if (winner == EQUAL) { // SI MATCH NUL
      *test = 0
      return 0
    } else if (winner == PLAYER2) { // SI P2 GAGNANT
      *test = MIN_BASE - depth
      return MIN_BASE - depth
    } else if (winner == PLAYER1) { // SI P1 GAGNANT
      *test = MAX_BASE + depth
      return MAX_BASE + depth
    }
  } else if (depth == DEPTH_MAX) { // SINON SI ON EST ARRIVEE A LA PROFONDEUR MAX
    *test = evaluate(board, aiScore, playerScore)
    return evaluate(board, aiScore, playerScore) // ON EVALUE LES NOEUDS FINAUX GRACE A UNE HEURISTIQUE
  } else { // SINON ON CALCULE EN ALTERNANCE MIN/MAX PAR RAPPORT AUX EVALUATIONS DES NOEUD FINAUX EN REMONTANT JUSQUA PROFONDEUR 0

    // coups := findMoves(board)
    // CALCULE_VALUE_ROUTINE.Add(len(coups))
    // var t int = 0
    // for t < len(coups) {
    //    findCoup1(t, coups ,&coups[t], aiScore, playerScore, &board, player)
    //   t++
    //   fmt.Println("go lunch a routine in CALCULE")
    // }
    // CALCULE_VALUE_ROUTINE.Wait()
    coups := findMoves(board)
 for index,coup := range coups {
   tmpAiScore, tmpPlayerScore = aiScore, playerScore
   tmpBoard = moveAndEat(board, coup.x, coup.y, player, &tmpAiScore, &tmpPlayerScore)
   //coups[index].poid = go calcValue(tmpBoard, depth + 1, tmpAiScore, tmpPlayerScore, player)
    go calcValue(tmpBoard, depth + 1, tmpAiScore, tmpPlayerScore, player, &coups[index].poid)

 }

    if (depth % 2 == 1) {
      *test = minCoup(coups).poid
      return minCoup(coups).poid
    } else {
      *test = maxCoup(coups).poid
      return maxCoup(coups).poid
    }
  }
  *test = 0
  return 0
}

func evaluate(board [SIZE][SIZE]int, aiScore int, playerScore int) int {
  //fmt.Print("\nHELLO\n")
  return 10
  //return findWeight(board, aiScore, playerScore)
}
