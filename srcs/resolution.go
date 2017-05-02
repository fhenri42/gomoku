package main
import ("fmt")

func getBestMove(board [SIZE][SIZE]int, aiScore int, playerScore int) move {
  fmt.Print("in the resolution:")

  coups := findMoves(board) // REMPLI UN ARRAY AVEC LES COUPS POTENTIELS
  var tmpBoard [SIZE][SIZE]int
  var tmpPlayerScore int
  var tmpAiScore int

  for index,coup := range coups { // PARCOURS LES COUPS
    tmpAiScore = aiScore
    tmpPlayerScore = playerScore
    tmpBoard = moveAndEat(board, coup.x, coup.y, PLAYER2, &tmpAiScore, &tmpPlayerScore)
    coups[index].poid = calcValue(tmpBoard, 1, tmpAiScore, tmpPlayerScore)
  }
  bestCoup := maxCoup(coups)
  return *bestCoup
}

func calcValue(board [SIZE][SIZE]int, depth int, aiScore int, playerScore int) int {
  end, winner := endGame(board) // OBSERVE LA GRILLE POUR SAVOIR SI C'EST LA FIN
  var tmpBoard [SIZE][SIZE]int
  var tmpPlayerScore int
  var tmpAiScore int
  var player int

  if (end) { // SI FIN DU GAME
    if (winner == EQUAL) { // SI MATCH NUL
      return 0
    } else if (winner == PLAYER2) { // SI P2 GAGNANT
      return MIN_BASE - depth
    } else if (winner == PLAYER1) { // SI P1 GAGNANT
      return MAX_BASE + depth
    }
  } else if (depth == DEPTH_MAX) { // SINON SI ON EST ARRIVEE A LA PROFONDEUR MAX
    return evaluate(board, aiScore, playerScore) // ON EVALUE LES NOEUDS FINAUX GRACE A UNE HEURISTIQUE
  } else { // SINON ON CALCULE EN ALTERNANCE MIN/MAX PAR RAPPORT AUX EVALUATIONS DES NOEUD FINAUX EN REMONTANT JUSQUA PROFONDEUR 0

    if (depth % 2 == 1) {
      player = PLAYER1
    } else {
      player = PLAYER2
    }

    coups := findMoves(board)
    for _,coup := range coups {
      tmpAiScore = aiScore
      tmpPlayerScore = playerScore
      tmpBoard = moveAndEat(board, coup.x, coup.y, player, &tmpAiScore, &tmpPlayerScore)
      coup.poid = calcValue(tmpBoard, depth + 1, tmpAiScore, tmpPlayerScore)
    }

    if (depth % 2 == 1) {
      return minCoup(coups).poid
    } else {
      return maxCoup(coups).poid
    }
  }
  return 0
}

func evaluate(board [SIZE][SIZE]int, AiScore int, PlayerScore int) int {
//  findWeight(board)
  return 10
}
