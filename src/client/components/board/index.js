import React, { Component, PropTypes } from 'react'
import { fromJS } from 'immutable'
import "./style.scss"
/*
RULE: http://jeuxdesociete.free.fr/jeux/jeu-gomoku.html
Capture: Captures are made by flanking a pair of the opponent’s stones
Free-threes: look example
win if 5 are aligne and no Capture possibole
win if 10 are remove from the board.
*/

function createBoard() {
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
// TODO a finir  !  quand hugo rentre
function findCapture(copy,x, y, pion, pion2) {

  if ( x < 19 && copy[x + 3][y] === pion && copy[x + 1][y] === pion2 && copy[x + 2][y] === pion2) {
    return { x1: x + 1, y1: y, x2: x + 2,y2:y }
  }
  if ( x > 0 &&  copy[x - 3][y] === pion && copy[x - 1][y] === pion2 && copy[x - 2][y] === pion2) {
    return { x1: x - 1, y1: y, x2: x - 2, y2: y }
  }
  if ( x < 19 && y < 19 && copy[x + 3][y + 3] === pion && copy[x + 1][y + 1] === pion2 && copy[x + 2][y + 2] === pion2) {
    return { x1: x + 1, y1: y + 1, x2: x + 2, y2: y + 2 }
  }
  if ( x > 0 && y > 0 && copy[x - 3][y - 3] === pion && copy[x - 1][x - 1] === pion2 && copy[x - 2][y - 2] === pion2) {
    return { x1: x - 1, y1: y - 1, x2: x - 2, y2: y - 2 }
  }
  if ( y < 19 && copy[x][y + 3] === pion && copy[x][y + 1] === pion2 && copy[x][y + 2] === pion2) {
    return { x1: x, y1: y + 1 , x2: x, y2:y + 2 }
  }
  if ( y > 0 && copy[x][y - 3] === pion && copy[x][y - 1] === pion2 && copy[x][y - 2] === pion2) {
    return { x1: x, y1: y - 1, x2: x, y2: y - 2 }
  }
  return false
}

class Board extends Component {

  static propTypes = {}

  state = {
    turn :  false,
    board: createBoard()
  }


  putADot = (x, y) => {
    const { turn } = this.state
    let copy = this.state.board
    if (turn === true) {
      copy[x][y] = 1
      const cord = findCapture(copy,x, y, 1, 2)
      console.log(cord);
      if (cord) {
        console.log(require('util').inspect(cord, { depth: null }));
        copy[cord.x1][cord.y1] =  0
        copy[cord.x2][cord.y2] =  0
      }
      this.setState({ board: copy })
      this.setState({turn: !turn})
    } else {
      copy[x][y] = 2
      const cord = findCapture(copy,x, y, 2, 1)
      console.log(cord);
      if (cord) {
        console.log(require('util').inspect(cord, { depth: null }));

        copy[cord.x1][cord.y1] =  0
        copy[cord.x2][cord.y2] =  0
      }
      this.setState({ board: copy })
      this.setState({turn: !turn})
    }
  }

  render () {
    const { board, turn } = this.state

    return (
      <div className="game">
      <h1 className="h1">
      {turn === true ? "Player 2" : "Player 1" }
      </h1>
      { board.map((line,x) => {
        return (
          <div className='line'>
          {
            line.map((bloc, y) => {
              if (bloc === 0) { return (<div className='bloc'> <div className='noDot' onClick={()=> this.putADot(x,y)}/> </div>) }
              else if(bloc === 1) {return(< div className='bloc'> <div className='blackDot'/> </div>) }
              else { return (<div className='bloc'> <div className='whiteDot'/></div>) }
            })

          } </div>
        )})}
        </div>
      )
    }
  }

  export default Board
