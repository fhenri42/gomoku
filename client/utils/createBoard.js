const createBoard = () => {
  let board = []
  for (let x = 0; x < 19; x++) {
    let tmp = []
    for (let y = 0; y < 19; y++) {
      tmp.push(0)
    }
    board.push(tmp)
  }
  return board
}

export default createBoard
