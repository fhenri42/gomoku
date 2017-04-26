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

export default lastStand
