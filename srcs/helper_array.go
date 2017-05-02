package main

func newBoard() [SIZE][SIZE]int {
	var board [SIZE][SIZE]int
	var x = 0
	for x < SIZE {
		var y = 0
		for y < SIZE {
			board[x][y] = 0
			y++
		}
		x++
	}
	return board
}

func copyBoard(board [SIZE][SIZE]int) [SIZE][SIZE]int {
  var newBoard [SIZE][SIZE]int

  var x = 0
  for x < SIZE {
    var y = 0
    for y < SIZE {
      newBoard[x][y] = board[x][y]
      y++
    }
    x++
  }
  return newBoard
}
