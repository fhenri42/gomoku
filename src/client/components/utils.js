const createBoard = () => {
  let board = [[]]
  for (let x = 0; x < 19; x++) {
    let tmp = []
    for (let y = 0; y < 19; y++) {
      tmp.push(0)
    }
    board.push(tmp)
  }
  return board
}

//TODO lastStand is not done
function lastStand() {

}

const forbidenMove = (copy, x, y, pion) => {

  if (x - 2 >= 0 && y + 2 <= 19 && y - 2 >= 0 && copy[x][y + 1] === pion &&  copy[x][y + 2] === pion && copy[x - 1][y - 1] === pion && copy[x - 2][y - 2] === pion) {
    return false
  }

  if (x - 2 >= 0 && y + 2 <= 19 && y - 2  >= 0 && copy[x][y - 1] === pion &&  copy[x][y - 2] === pion && copy[x - 1][y + 1] === pion && copy[x - 2][y + 2] === pion) {
    return false
  }

  if ( x + 2 <= 19 && y + 2 <= 19 && y - 2  >= 0 && copy[x][y + 1] === pion &&  copy[x][y + 2] === pion && copy[x + 1][y - 1] === pion && copy[x + 2][y - 2] === pion) {
    return false
  }
  if ( x + 2 <= 19 && y + 2 <= 19 && y - 2  >= 0 && copy[x][y - 1] === pion &&  copy[x][y - 2] === pion && copy[x + 1][y + 1] === pion && copy[x + 2][y + 2] === pion) {
    return false
  }

  if (x + 2 <= 19 && x - 2 >= 0 && y - 2 >= 0 && copy[x - 1][y] === pion &&  copy[x - 2][y] === pion && copy[x + 1][y - 1] === pion && copy[x + 2][y - 2] === pion) {
    return false
  }

  if (x - 2 >= 0 && x + 2 <= 19 && y + 2 <= 19 && copy[x - 1][y] === pion &&  copy[x - 2][y] === pion && copy[x + 1][y + 1] === pion && copy[x + 2][y + 2] === pion) {
    return false
  }

  if (x + 2 <= 19 && x - 2 >= 0 && y + 2 <= 19 && copy[x + 1][y] === pion &&  copy[x + 2][y] === pion && copy[x - 1][y + 1] === pion && copy[x - 2][y + 2] === pion) {
    return false
  }

  if (x + 2 <= 19 && x - 2 >= 0 && y - 2 >= 0 && copy[x + 1][y] === pion &&  copy[x + 2][y] === pion && copy[x - 1][y - 1] === pion && copy[x - 2][y - 2] === pion) {
    return false
  }


  return true

}

const findLine = (copy, x, y, pion) => {

  if ( x + 4 <= 19 && copy[x][y] === pion && copy[x + 1][y] === pion && copy[x + 2][y] === pion && copy[x + 3][y] === pion && copy[x + 4][y] === pion) {
    return true
  }
  if ( x - 4 >= 0 && copy[x][y] === pion && copy[x - 1][y] === pion && copy[x - 2][y] === pion && copy[x - 3][y] === pion && copy[x - 4][y] === pion) {
    return true
  }
  if ( y + 4 <= 19 && copy[x][y] === pion && copy[x][y + 1] === pion && copy[x][y + 2] === pion && copy[x][y + 3] === pion && copy[x][y + 4] === pion) {
    return true
  }
  if ( y - 4 >= 0 && copy[x][y] === pion && copy[x][y - 1] === pion && copy[x][y - 2] === pion && copy[x][y - 3] === pion && copy[x][y - 4] === pion) {
    return true
  }


  if ( y + 4 <= 19 && x + 4 <= 19 && copy[x][y] === pion && copy[x + 1][y + 1] === pion && copy[x + 2][y + 2] === pion && copy[x + 3][y + 3] === pion && copy[x + 4][y + 4] === pion) {
    return true
  }
  if ( y - 4 >= 0 && x - 4 >= 0 && copy[x][y] === pion && copy[x - 1][y - 1] === pion && copy[x - 2][y - 2] === pion && copy[x - 3][y - 3] === pion && copy[x - 4][y - 4] === pion) {
    return true
  }

  if ( y + 4 <= 19 && x - 4 >= 0 && copy[x][y] === pion && copy[x - 1][y + 1] === pion && copy[x - 2][y + 2] === pion && copy[x - 3][y + 3] === pion && copy[x - 4][y + 4] === pion) {
    return true
  }

  if ( y - 4 >= 0 && x + 4 <= 19 && copy[x][y] === pion && copy[x + 1][y - 1] === pion && copy[x + 2][y - 2] === pion && copy[x + 3][y - 3] === pion && copy[x + 4][y - 4] === pion) {
    return true
  }
  false
}
const findCapture = (copy,x, y, pion, pion2) => {

  if ( x + 3 <= 19 && copy[x + 3][y] === pion && copy[x + 1][y] === pion2 && copy[x + 2][y] === pion2) {
    return { x1: x + 1, y1: y, x2: x + 2,y2:y }
  }

  if ( x - 3 >= 0 &&  copy[x - 3][y] === pion && copy[x - 1][y] === pion2 && copy[x - 2][y] === pion2) {
    return { x1: x - 1, y1: y, x2: x - 2, y2: y }
  }

  if ( y + 3 <= 19 && copy[x][y + 3] === pion && copy[x][y + 1] === pion2 && copy[x][y + 2] === pion2) {
    return { x1: x, y1: y + 1 , x2: x, y2:y + 2 }
  }

  if ( y - 3 >= 0 && copy[x][y - 3] === pion && copy[x][y - 1] === pion2 && copy[x][y - 2] === pion2) {
    return { x1: x, y1: y - 1, x2: x, y2: y - 2 }
  }

  if ( x + 3 <= 19 && y + 3 <= 19 && copy[x + 3][y + 3] === pion && copy[x + 1][y + 1] === pion2 && copy[x + 2][y + 2] === pion2) {
    return { x1: x + 1, y1: y + 1, x2: x + 2, y2: y + 2 }
  }

  if ( x - 3 >= 0 && y - 3 >= 0 && copy[x - 3][y - 3] === pion && copy[x - 1][y - 1] === pion2 && copy[x - 2][y - 2] === pion2) {
    return { x1: x - 1, y1: y - 1, x2: x - 2, y2: y - 2 }
  }

  if ( x - 3 >= 0 && y + 3 <= 19 && copy[x - 3][y + 3] === pion && copy[x - 1][y + 1] === pion2 && copy[x - 2][y + 2] === pion2) {
    return { x1: x - 1, y1: y + 1, x2: x - 2, y2: y + 2 }
  }

  if ( x + 3 <= 19 && y - 3 >= 0 && copy[x + 3][y - 3] === pion && copy[x + 1][y - 1] === pion2 && copy[x + 2][y - 2] === pion2) {
    return { x1: x + 1, y1: y - 1, x2: x + 2, y2: y - 2 }
  }


  return false
}

module.exports = {
findCapture,
findLine,
forbidenMove,
createBoard,
}
