package main

import (

	"math"
)
const nbAligne = 5
//Trouver si est possible de win sur horiZontal
func findHor(board[SIZE][SIZE]int,couleur int, x int, y int) (float64, float64, float64){

	var pass bool = false; //permet de voir si on a passé la case étudiée
	var compteur float64 = 0; //compte le nombre de possibilités pour une direction
	var bonus float64 = 0; //point bonus liée aux jetons alliés dans cette même direction
	var centre float64 = 0; //regarde si le jeton a de l'espace de chaque côté
	var i int = 0

	for i < SIZE {
		if i == x {
			centre = compteur + 1
			pass = true
			//continue
		}
		switch board[i][y] {
			case 0: //case vide
			compteur++
			break
			case couleur: //jeton allié
			compteur++
			bonus++
			break
			default: //jeton adverse
			if pass {
				i = SIZE //il n'y aura plus de liberté supplémentaire, on arrête la recherche ici
				} else {
					//on réinitialise la recherche
					compteur = 0
					bonus = 0
				}
			}
			i++
		}
		return centre, compteur, bonus
	}

	func iaAnalyse(board [SIZE][SIZE]int, x int, y int) int {
		var estimation float64 = 0;
		var centre float64 = 0
		var pLiberte float64 = 1; //pondération sur le nombre de liberté
		var pBonus float64 = 1; //pondération Bonus
		var pCentre float64 = 2; //pondération pour l'espace situé de chaque côté

		tmpCenter, compteur, bonus := findHor(board, board[x][y], x, y)
		centre += tmpCenter
		if compteur>=nbAligne {
			estimation += compteur * pLiberte + bonus * pBonus + (1 - math.Abs(centre / (compteur - 1) - 0.5)) * compteur * pCentre;
		}
		return int(estimation)
	}

	func iaEstimation(board [SIZE][SIZE]int) int {
		var estimation int = 0;
		var x int = 0
		for x < SIZE {
			var y int = 0
			for y < SIZE {
				switch  board[x][y] {
				case 1:
					estimation += iaAnalyse(board,x,y)
					break
				case 2:
					estimation -= iaAnalyse(board,x,y)
					break
				default:
					break
				}
				y++
			}
			x++
		}
		return estimation;
	}

	func findWeight(board [SIZE][SIZE]int, aiScore int, playerScore int) int {
		//fmt.Print(aiScore)
		//fmt.Print(playerScore)
		estimation := iaEstimation(board)
		return estimation
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
