const createBoard = () => {
  let board = []
  for (let x = 0; x <= 19; x++) {
    let tmp = []
    for (let y = 0; y <= 19; y++) {
      tmp.push(0)
    }
    board.push(tmp)
  }
  return board
}

//TODO lastStand is not done
function lastStand(copy, x, y, pion) {
  const pion2 = pion == 1 ? 2 : 1
  let h = 0
  console.log("HERE");
  console.log(pion2);
  if (x - 1 >= 0 && copy[x - 1][y] === pion && (x - 2 >= 0  && copy[x - 2][y] !== pion || x + 1 <= 19 && copy[x + 1][y] !== pion)) {
    console.log("X: ", x);
    if (x - 2 < 0 || x + 1 > 19) {
      h = h + 3
    } else if (x - 2 >= 0 && copy[x - 2][y] === pion2) {
      h++
    } else if (x + 1 <= 19 && copy[x + 1][y] === pion2) {
      h++
    }
    if (h == 1) {
      console.log("1");
      return true
    }
    else { h = 0 }
  }
  if (x + 1 <= 19 && copy[x + 1][y] === pion && (x + 2 <= 19  && copy[x + 2][y] !== pion || x - 1 >= 0 && copy[x - 1][y] !== pion)) {
    if (x + 2 > 19 || x - 1 < 0) {
      h = h + 3
    } else if (x + 2 <= 19 && copy[x + 2][y] === pion2) {
      h++
    } else if (x - 1 >= 0 && copy[x - 1][y] === pion2) {
      h++
    }
    if (h == 1) {
      console.log("2");
      return true
    }
    else { h = 0 }
  }
  if (y + 1 <= 19 && copy[x][y + 1] === pion && (y + 2 <= 19  && copy[x][y + 2] !== pion || y - 1 >= 0 && copy[x][y - 1] !== pion)) {
    if (y + 2 > 19 || y - 1 < 0) {
      h = h + 3
    } else if (y + 2 <= 19 && copy[x][y + 2] === pion2) {
      h++
    } else if (y - 1 >= 0 && copy[x][y - 1] === pion2) {
      h++
    }
    if (h == 1) {
      console.log("3");
      return true
    }
    else { h = 0 }
  }
  if (y - 1 >= 0 && copy[x][y - 1] === pion && (y - 2 >= 0  && copy[x][y - 2] !== pion || y + 1 <= 19 && copy[x][y + 1] !== pion)) {
    if (y - 2 < 0 || y + 1 > 19) {
      h = h + 3
    } else if (y - 2 >= 0 && copy[x][y - 2] === pion2) {
      h++
    } else if (y + 1 <= 19 && copy[x][y + 1] === pion2) {
      h++
    }
    if (h == 1) {
      console.log("4");
      return true
    }
    else { h = 0 }
  }
  if (x - 1 >= 0 && y - 1 >= 0 && copy[x - 1][y - 1] === pion && (x - 2 >= 0 && y - 2 >= 0 && copy[x - 2][y - 2] !== pion || x + 1 <= 19 && y + 1 <= 19 && copy[x + 1][y + 1] !== pion)) {
    if ((x - 2 < 0 && y - 2 < 0) || (x + 1 > 19 && y + 1 > 19)) {
      h = h + 3
    } else if (x - 2 >= 0 && y - 2 >= 0 && copy[x - 2][y - 2] === pion2) {
      h++
    } else if (x + 1 <= 19  && y + 1 <= 19 && copy[x + 1][y + 1] === pion2) {
      h++
    }
    if (h == 1) { return true }
    else { h = 0 }
  }
  if (x - 1 >= 0 && y + 1 >= 0 && copy[x - 1][y + 1] === pion && (x - 2 >= 0 && y + 2 >= 0 && copy[x - 2][y + 2] !== pion || x + 1 <= 19 && y - 1 <= 19 && copy[x + 1][y - 1] !== pion)) {
    if ((x - 2 < 0 && y + 2 < 0) || (x + 1 > 19 && y - 1 > 19)) {
      h = h + 3
    } else if (x - 2 >= 0 && y + 2 >= 0 && copy[x - 2][y + 2] === pion2) {
      h++
    } else if (x + 1 <= 19  && y - 1 <= 19 && copy[x + 1][y - 1] === pion2) {
      h++
    }
    if (h == 1) { return true }
    else { h = 0 }
  }
  if (x + 1 >= 0 && y - 1 >= 0 && copy[x + 1][y - 1] === pion && (x + 2 >= 0 && y - 2 >= 0 && copy[x + 2][y - 2] !== pion || x - 1 <= 19 && y + 1 <= 19 && copy[x - 1][y + 1] !== pion)) {
    if ((x + 2 < 0 && y - 2 < 0) || (x - 1 > 19 && y + 1 > 19)) {
      h = h + 3
    } else if (x + 2 >= 0 && y - 2 >= 0 && copy[x + 2][y - 2] === pion2) {
      h++
    } else if (x - 1 <= 19  && y + 1 <= 19 && copy[x - 1][y + 1] === pion2) {
      h++
    }
    if (h == 1) { return true }
    else { h = 0 }
  }
  if (x + 1 >= 0 && y + 1 >= 0 && copy[x + 1][y + 1] === pion && (x + 2 >= 0 && y + 2 >= 0 && copy[x + 2][y + 2] !== pion || x - 1 <= 19 && y - 1 <= 19 && copy[x - 1][y - 1] !== pion)) {
    if ((x + 2 < 0 && y + 2 < 0) || (x - 1 > 19 && y - 1 > 19)) {
      h = h + 3
    } else if (x + 2 >= 0 && y + 2 >= 0 && copy[x + 2][y + 2] === pion2) {
      h++
    } else if (x - 1 <= 19  && y - 1 <= 19 && copy[x - 1][y - 1] === pion2) {
      h++
    }
    if (h == 1) { return true }
    else { h = 0 }
  }
  return false
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

const findLine = (copy, pion, SetisEatable, isEatable) => {
  console.log(isEatable);
  for (let j = 0; j < 20; j++) {
    for (let i = 0; i < 20; i++) {
      if ( i + 4 <= 19 && copy[i][j] === pion && copy[i + 1][j] === pion && copy[i + 2][j] === pion && copy[i + 3][j] === pion && copy[i + 4][j] === pion) {
        if (lastStand(copy, i, j, pion) || lastStand(copy, i + 1, j, pion) || lastStand(copy, i + 2, j, pion) || lastStand(copy, i + 3, j, pion) || lastStand(copy, i + 4, j, pion)) {
          console.log("TRUE");
          SetisEatable(isEatable + 1)
        } else {
          console.log("FALSE1");
        }
        return true
      }
      if ( j + 4 <= 19 && copy[i][j] === pion && copy[i][j + 1] === pion && copy[i][j + 2] === pion && copy[i][j + 3] === pion && copy[i][j + 4] === pion) {
        if (lastStand(copy, i, j, pion) || lastStand(copy, i, j + 1, pion) || lastStand(copy, i, j + 2, pion) || lastStand(copy, i, j + 3, pion) || lastStand(copy, i, j + 4, pion)) {
          console.log("TRUE");
          SetisEatable(isEatable + 1)
        } else {
          console.log("FALSE3");
        }
        return true
      }
      if ( j + 4 <= 19 && i + 4 <= 19 && copy[i][j] === pion && copy[i + 1][j + 1] === pion && copy[i + 2][j + 2] === pion && copy[i + 3][j + 3] === pion && copy[i + 4][j + 4] === pion) {
      if (lastStand(copy, i, j, pion) || lastStand(copy, i + 1, j + 1, pion) || lastStand(copy, i + 2, j + 2, pion) || lastStand(copy, i + 3, j + 3, pion) || lastStand(copy, i + 4, j + 4, pion)) {
          console.log("TRUE");
          SetisEatable(isEatable + 1)
        } else {
          console.log("FALSE5");
        }
        return true
      }
      if ( j - 4 >= 0 && i + 4 <= 19 && copy[i][j] === pion && copy[i + 1][j - 1] === pion && copy[i + 2][j - 2] === pion && copy[i + 3][j - 3] === pion && copy[i + 4][j - 4] === pion) {
        if (lastStand(copy, i, j, pion) || lastStand(copy, i + 1, j - 1, pion) || lastStand(copy, i + 2, j - 2, pion) || lastStand(copy, i + 3, j - 3, pion) || lastStand(copy, i + 4, j - 4, pion)) {
          console.log("TRUE");
          SetisEatable(isEatable + 1)
        } else {
          console.log("FALSE8");
        }
        return true
      }

    }
  }
  return false
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
lastStand,
}
