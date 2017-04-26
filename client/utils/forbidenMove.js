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

export default forbidenMove
