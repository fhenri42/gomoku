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

export default findCapture
