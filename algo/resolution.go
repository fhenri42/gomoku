package main
//import ("fmt")

const PLAYER1 = 1
const PLAYER2 = 2
const DEPTH_MAX = 4
const EQUAL = 0
const MIN_BASE = 1000000
const MAX_BASE = -1000000

func getBestState(state [20][20]int) [20][20]int {
  fmt.Printf("in the resolution:")

  coups := findMoves(state) // REMPLI UN ARRAY AVEC LES COUPS POTENTIELS
  var bestCoup *moveNode = nil
  var max = 0

  for _,coup := range coups { // PARCOURS LES COUPS
    newState := getNewState(state, coup, PLAYER2) // SIMULE LE COUP DANS UNE GRILLE
    poid := calcValue(newState, 1) // CALCULE LE POID DE CE COUP
    if (bestCoup == nil || poid > max) {
      max = poid
      bestCoup = coup // MET A JOUR LE MEILLEUR COUP A JOUER (MAX)
    }
  }
  return getNewState(state, bestCoup, PLAYER2) // RETURN LA GRILLE ASSOCIEE A CE COUP
}

func getNewState(state [20][20]int, move moveNode, player int) [20][20]int {
  var newState [20][20]int

  for i, line := range state {
    var newLine [20]int
    for j, bloc := range line {
      if (i == coup.x && j == coup.y) {
        newLine[j] = player
      } else {
        newLine[j] = bloc
      }
    }
    newState[i] = newLine
  }
  return newState
}

func calcValue(state [20][20]int, depth int) int {
  end, winner := endGame(state) // OBSERVE LA GRILLE POUR SAVOIR SI C'EST LA FIN
  var player = 0
  var max = MAX_BASE
  var min = MIN_BASE

  if (end) { // SI FIN DU GAME
    if (winner == EQUAL) { // SI MATCH NUL
      return 0
    } else if (winner == PLAYER2) { // SI P2 GAGNANT
      return MIN_BASE - depth
    } else if (winner == PLAYER1) { // SI P1 GAGNANT
      return MAX_BASE + depth
    }
  } else if (depth == DEPTH_MAX) { // SINON SI ON EST ARRIVEE A LA PROFONDEUR MAX
    return evaluate(state) // ON EVALUE LES NOEUDS FINAUX GRACE A UNE HEURISTIQUE
  } else { // SINON ON CALCULE EN ALTERNANCE MIN/MAX PAR RAPPORT AUX EVALUATIONS DES NOEUD FINAUX EN REMONTANT JUSQUA PROFONDEUR 0
    if (depth % 2 == 1) {
      player = PLAYER1
    } else {
      player = PLAYER2
    }

    coups := findMoves(state)
    for _,coup := range coups {
      newState := getNewState(state, coup, player)
      poid := calcValue(newState, depth + 1)
      if (poid > max) {
        max = poid
      }
      if (poid < min) {
        min = poid
      }
    }
  }
  if (depth % 2 == 1) {
    return min
  } else {
    return max
  }
}

func endGame(state [20][20]int) (bool, int) {
  return false, 1
}

func evaluate(state [20][20]int) int {
  return 10
}
