package main
import ("fmt")

const PLAYER1 = 1
const PLAYER2 = 2
const DEPTH_MAX = 1
const EQUAL = 0
const MIN_BASE = 1000000
const MAX_BASE = -1000000
const AMP = 2
const SIZE = 19

type move struct {
  x int
  y int
  poid int
}

func getBestState(state [][]int, AiScore int, PlayerScore int) [][]int {
  fmt.Print("in the resolution:")

  coups := findMoves(state) // REMPLI UN ARRAY AVEC LES COUPS POTENTIELS
  var bestCoup *move

  for _,coup := range coups { // PARCOURS LES COUPS

    state[coup.x][coup.y] = PLAYER2
    coup.poid = calcValue(state, 1, AiScore, PlayerScore) // CALCULE LE POID DE CE COUP
    state[coup.x][coup.y] = 0
  }

  bestCoup = maxCoup(coups)
  state[bestCoup.x][bestCoup.y] = PLAYER2
  return state // RETURN LA GRILLE ASSOCIEE A CE COUP
}

func calcValue(state [][]int, depth int, AiScore int, PlayerScore int) int {
  end, winner := endGame(state) // OBSERVE LA GRILLE POUR SAVOIR SI C'EST LA FIN
  var player = 0

  if (end) { // SI FIN DU GAME
    if (winner == EQUAL) { // SI MATCH NUL
      return 0
    } else if (winner == PLAYER2) { // SI P2 GAGNANT
      return MIN_BASE - depth
    } else if (winner == PLAYER1) { // SI P1 GAGNANT
      return MAX_BASE + depth
    }
  } else if (depth == DEPTH_MAX) { // SINON SI ON EST ARRIVEE A LA PROFONDEUR MAX
    return evaluate(state, AiScore, PlayerScore) // ON EVALUE LES NOEUDS FINAUX GRACE A UNE HEURISTIQUE
  } else { // SINON ON CALCULE EN ALTERNANCE MIN/MAX PAR RAPPORT AUX EVALUATIONS DES NOEUD FINAUX EN REMONTANT JUSQUA PROFONDEUR 0
    if (depth % 2 == 1) {
      player = PLAYER1
    } else {
      player = PLAYER2
    }

    coups := findMoves(state)
    for _,coup := range coups {

      state[coup.x][coup.y] = player
      coup.poid = calcValue(state, depth + 1, AiScore, PlayerScore)
      state[coup.x][coup.y] = 0
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
