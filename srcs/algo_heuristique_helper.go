package main

import (
)

//Trouver si est possible de win sur horiZontal
func findHor(board[SIZE][SIZE]int,couleur int, x int, y int) (float64, float64, float64) {

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

  //Trouver si est possible de win sur vertical
  func findVer(board[SIZE][SIZE]int,couleur int, x int, y int) (float64, float64, float64) {

    var pass bool = false; //permet de voir si on a passé la case étudiée
    var compteur float64 = 0; //compte le nombre de possibilités pour une direction
    var bonus float64 = 0; //point bonus liée aux jetons alliés dans cette même direction
    var centre float64 = 0; //regarde si le jeton a de l'espace de chaque côté
    var i int = 0

    for i < SIZE {
      if i == y {
        centre = compteur + 1
        pass = true
        //continue
      }
      switch board[x][i] {
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

    func findDigo1(board[SIZE][SIZE]int,couleur int, x int, y int) (float64, float64, float64) {

      var compteur float64 = 0; //compte le nombre de possibilités pour une direction
      var bonus float64 = 0; //point bonus liée aux jetons alliés dans cette même direction
      var centre float64 = 0; //regarde si le jeton a de l'espace de chaque côté
      var i int = x - 1
      var j int = y - 1

      for (i > 0 && j > 0) {
        switch board[i][j] {
          case 0: //case vide
          compteur++
          break
          case couleur: //jeton allié
          compteur++
          bonus++
          break
          default: //jeton adverse
            i = 0
          }
          i--
          j--
      }
      centre = compteur + 1
      i = x + 1
      j = y + 1

      for (i < SIZE  && j < SIZE) {
        switch board[i][j] {
          case 0: //case vide
          compteur++
          break
          case couleur: //jeton allié
          compteur++
          bonus++
          break
          default: //jeton adverse
            i = SIZE
          }
          i++
          j++
      }
        return centre, compteur, bonus
      }

      func findDigo2(board[SIZE][SIZE]int,couleur int, x int, y int) (float64, float64, float64) {

        var compteur float64 = 0; //compte le nombre de possibilités pour une direction
        var bonus float64 = 0; //point bonus liée aux jetons alliés dans cette même direction
        var centre float64 = 0; //regarde si le jeton a de l'espace de chaque côté
        var i int = x - 1
        var j int = y - 1

        for (i > 0 && j > 0) {
          switch board[i][j] {
            case 0: //case vide
            compteur++
            break
            case couleur: //jeton allié
            compteur++
            bonus++
            break
            default: //jeton adverse
              i = 0
            }
            i--
            j--
        }

        centre = compteur + 1
        i = x + 1
        j = y - 1

        for (i < SIZE  && j > 0) {
          switch board[i][j] {
            case 0: //case vide
            compteur++
            break
            case couleur: //jeton allié
            compteur++
            bonus++
            break
            default: //jeton adverse
              i = SIZE
            }
            i++
            j--
        }

        return centre, compteur, bonus
      }
