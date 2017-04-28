package main

import ("fmt")

func findWeight(state [][]int, noeuds []move ) int {

  for _,noeud := range noeuds {
    findWInByValue(noeud)
    findWInByLine(noeud)
    findCounterWin(noeud)
    findEatMove(noeud)
  }
  fmt.Print(state)

  return 9
}



/*
1) cherher a manger un pyon si on est a 9 sur 10                    Poids: 2000
2) cherher un allignmen de 5 sans possibiliter detre contrue        Poids: 1900
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
8)heuristic take past player actions into account to identify patterns and weigh board states accordingl
*/
