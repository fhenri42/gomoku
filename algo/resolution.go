package main
import ("fmt")



type move struct {
  x int
  y int
  poid int
}

func simulateMove(state [][]int, coup move, player int, aiScore *int, playerScore *int) [][]int {
  otherPlayer := player % 2 + 1
  newState := make([][]int, SIZE)
  var i = 0
  var j = 0

  for i < SIZE {
    j = 0
    newLine := make([]int, SIZE)
    for j < SIZE {
      newLine[j] = state[i][j]
      j++
    }
    newState[i] = newLine
    i++
  }

  i = -1
  for i <= 1 {
    j = -1
    for j <= 1 {
      if (coup.x + i * 3 > 0 && coup.x + i * 3 < SIZE - 1 && coup.y + j * 3 > 0 && coup.y + j * 3 < SIZE - 1 && newState[coup.x + i][coup.y + j] == otherPlayer && newState[coup.x + i * 2][coup.y + j * 2] == otherPlayer && newState[coup.x + i * 3][coup.y + j * 3] == player) {
        newState[coup.x + i][coup.y + j] = 0
        newState[coup.x + i * 2][coup.y + j * 2] = 0
        if (player == 1) {
          *playerScore += 2
        } else {
          *aiScore += 2
        }
      }
      j++
    }
    i++
  }
  newState[coup.x][coup.y] = player
  return newState
}

func getBestState(state [][]int, aiScore int, playerScore int) ([][]int, int, int) {
  fmt.Print("in the resolution:")

  coups := findMoves(state) // REMPLI UN ARRAY AVEC LES COUPS POTENTIELS
  var tmpState [][]int
  var tmpPlayerScore int
  var tmpAiScore int

  for index,coup := range coups { // PARCOURS LES COUPS
    tmpAiScore = aiScore
    tmpPlayerScore = playerScore
    tmpState = simulateMove(state, coup, PLAYER2, &tmpAiScore, &tmpPlayerScore)
    coups[index].poid = calcValue(tmpState, 1, tmpAiScore, tmpPlayerScore)
  }
  bestCoup := maxCoup(coups)
  tmpState = simulateMove(state, *bestCoup, PLAYER2, &aiScore, &playerScore) // RETURN LA GRILLE ASSOCIEE A CE COUP
  return tmpState, aiScore, playerScore
}

func calcValue(state [][]int, depth int, aiScore int, playerScore int) int {
  end, winner := endGame(state) // OBSERVE LA GRILLE POUR SAVOIR SI C'EST LA FIN
  var tmpState [][]int
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
    return evaluate(state, aiScore, playerScore) // ON EVALUE LES NOEUDS FINAUX GRACE A UNE HEURISTIQUE
  } else { // SINON ON CALCULE EN ALTERNANCE MIN/MAX PAR RAPPORT AUX EVALUATIONS DES NOEUD FINAUX EN REMONTANT JUSQUA PROFONDEUR 0

    if (depth % 2 == 1) {
      player = PLAYER1
    } else {
      player = PLAYER2
    }

    coups := findMoves(state)
    for _,coup := range coups {
      tmpAiScore = aiScore
      tmpPlayerScore = playerScore
      tmpState = simulateMove(state, coup, player, &tmpAiScore, &tmpPlayerScore)
      coup.poid = calcValue(tmpState, depth + 1, tmpAiScore, tmpPlayerScore)
    }

    if (depth % 2 == 1) {
      return minCoup(coups).poid
    } else {
      return maxCoup(coups).poid
    }
  }
  return 0
}

func evaluate(state [][]int, AiScore int, PlayerScore int) int {
//  findWeight(state)
  return 10
}
