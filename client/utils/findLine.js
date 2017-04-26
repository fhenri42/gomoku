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

export default findLine
