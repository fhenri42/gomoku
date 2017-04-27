package main
import ("fmt")


const AMPCHOSSE = 2
type moveNoeued struct {
  x int
  y int
}

func expand(x int, y int, newTab[20][20]int, moves *[]moveNoeued) {
  var tmp moveNoeued
  fmt.Print("oeuo")
  var amp = 1
  for amp <= AMPCHOSSE {

  if (x + amp < 20 && newTab[x + amp][y] == 0) { tmp.x = x + amp; tmp.y = y; *moves = append(*moves, tmp) }
  if (x - amp >= 0 && newTab[x - amp][y] == 0) { tmp.x = x - amp; tmp.y = y; *moves = append(*moves, tmp) }
  if (y + amp < 20 && newTab[x][y + amp] == 0) { tmp.x = x; tmp.y = y + amp; *moves = append(*moves, tmp) }
  if (y - amp >= 0 && newTab[x][y - amp] == 0) { tmp.x = x; tmp.y = y - amp; *moves = append(*moves, tmp) }
  if (x + amp < 20 && y + amp < 20 && newTab[x + amp][y + amp] == 0) { tmp.x = x + amp; tmp.y = y + amp; *moves = append(*moves, tmp) }
  if (x - amp >= 0 && y - amp >= 0 && newTab[x - amp][y - amp] == 0) { tmp.x = x - amp; tmp.y = y; *moves = append(*moves, tmp) }
  if (x - amp >= 0 && y + amp < 20 && newTab[x - amp][y + amp] == 0) { tmp.x = x - amp; tmp.y = y + amp; *moves = append(*moves, tmp) }
  if (x + amp < 20 && y - amp >= 0 && newTab[x + amp][y - amp] == 0) { tmp.x = x + amp; tmp.y = y - amp; *moves = append(*moves, tmp) }
  amp++
}

}

func findMoves(newTab[20][20]int) ([]moveNoeued) {

  moves := make([]moveNoeued, 1)
  var x = 0
  for x < 20 {
    var y = 0
    for y < 20 {
      fmt.Print(newTab[x][y])
      if newTab[x][y] != 0 { expand(x, y, newTab, &moves) }
      y++
    }
    x++
  }
  fmt.Print(moves)
  return moves
}
