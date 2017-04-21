import React, { Component, PropTypes } from 'react'
import { fromJS } from 'immutable'
import "./style.scss"
/*
RULE: http://jeuxdesociete.free.fr/jeux/jeu-gomoku.html
Capture: Captures are made by flanking a pair of the opponent’s stones
Free-threes: look example
win if 5 are aligne and no Capture possibole
win if 10 are remove from the board.
//last stand
//forbiden
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

function lastStand() {

}
//TODO manque la diagonalle
function forbidenMove(copy, x, y, pion) {

  if (x - 2 > 0 && y + 2 < 19 && y - 2 > 0 && copy[x][y + 1] === pion &&  copy[x][y + 2] === pion && copy[x - 1][y - 1] === pion && pion[x - 2][y - 2]) {
    return false
  }

  if (x - 2 > 0 && y + 2 < 19 && y - 2  > 0 && copy[x][y - 1] === pion &&  copy[x][y - 2] === pion && copy[x - 1][y + 1] === pion && pion[x - 2][y + 2]) {
    return false
  }

  if ( x + 2 < 19 && y + 2 < 19 && y - 2  > 0 && copy[x][y + 1] === pion &&  copy[x][y + 2] === pion && copy[x + 1][y - 1] === pion && pion[x + 2][y - 2]) {
    return false
  }
  if ( x + 2 < 19 && y + 2 < 19 && y - 2  > 0 && copy[x][y - 1] === pion &&  copy[x][y - 2] === pion && copy[x + 1][y + 1] === pion && pion[x + 2][y + 2]) {
    return false
  }

  return true

}

function findLine(copy, x, y, pion) {

  if ( x + 4 < 19 && copy[x][y] === pion && copy[x + 1][y] === pion && copy[x + 2][y] === pion && copy[x + 3][y] === pion && copy[x + 4][y] === pion) {
    return true
  }
  if ( x - 4 > 0 && copy[x][y] === pion && copy[x - 1][y] === pion && copy[x - 2][y] === pion && copy[x - 3][y] === pion && copy[x - 4][y] === pion) {
    return true
  }
  if ( y + 4 < 19 && copy[x][y] === pion && copy[x][y + 1] === pion && copy[x][y + 2] === pion && copy[x][y + 3] === pion && copy[x][y + 4] === pion) {
    return true
  }
  if ( y - 4 > 0 && copy[x][y] === pion && copy[x][y - 1] === pion && copy[x][y - 2] === pion && copy[x][y - 3] === pion && copy[x][y - 4] === pion) {
    return true
  }


  if ( y + 4 < 19 && x + 4 < 19 && copy[x][y] === pion && copy[x + 1][y + 1] === pion && copy[x + 2][y + 2] === pion && copy[x + 3][y + 3] === pion && copy[x + 4][y + 4] === pion) {
    return true
  }
  if ( y - 4 > 0 && x - 4 > 0 && copy[x][y] === pion && copy[x - 1][y - 1] === pion && copy[x - 2][y - 2] === pion && copy[x - 3][y - 3] === pion && copy[x - 4][y - 4] === pion) {
    return true
  }

  if ( y + 4 < 19 && x - 4 > 0 && copy[x][y] === pion && copy[x - 1][y + 1] === pion && copy[x - 2][y + 2] === pion && copy[x - 3][y + 3] === pion && copy[x - 4][y + 4] === pion) {
    return true
  }

  if ( y - 4 > 0 && x + 4 < 19 && copy[x][y] === pion && copy[x + 1][y - 1] === pion && copy[x + 2][y - 2] === pion && copy[x + 3][y - 3] === pion && copy[x + 4][y - 4] === pion) {
    return true
  }

  false
}
function findCapture(copy,x, y, pion, pion2) {

  if ( x + 3 < 19 && copy[x + 3][y] === pion && copy[x + 1][y] === pion2 && copy[x + 2][y] === pion2) {
    return { x1: x + 1, y1: y, x2: x + 2,y2:y }
  }

  if ( x - 3 > 0 &&  copy[x - 3][y] === pion && copy[x - 1][y] === pion2 && copy[x - 2][y] === pion2) {
    return { x1: x - 1, y1: y, x2: x - 2, y2: y }
  }

  if ( y + 3 < 19 && copy[x][y + 3] === pion && copy[x][y + 1] === pion2 && copy[x][y + 2] === pion2) {
    return { x1: x, y1: y + 1 , x2: x, y2:y + 2 }
  }

  if ( y - 3 > 0 && copy[x][y - 3] === pion && copy[x][y - 1] === pion2 && copy[x][y - 2] === pion2) {
    return { x1: x, y1: y - 1, x2: x, y2: y - 2 }
  }

  if ( x + 3 < 19 && y + 3 < 19 && copy[x + 3][y + 3] === pion && copy[x + 1][y + 1] === pion2 && copy[x + 2][y + 2] === pion2) {
    return { x1: x + 1, y1: y + 1, x2: x + 2, y2: y + 2 }
  }

  if ( x - 3 > 0 && y - 3 > 0 && copy[x - 3][y - 3] === pion && copy[x - 1][y - 1] === pion2 && copy[x - 2][y - 2] === pion2) {
    return { x1: x - 1, y1: y - 1, x2: x - 2, y2: y - 2 }
  }

  if ( x - 3 > 0 && y + 3 < 19 && copy[x - 3][y + 3] === pion && copy[x - 1][y + 1] === pion2 && copy[x - 2][y + 2] === pion2) {
    return { x1: x - 1, y1: y + 1, x2: x - 2, y2: y + 2 }
  }

  if ( x + 3 < 19 && y - 3 > 0 && copy[x + 3][y - 3] === pion && copy[x + 1][y - 1] === pion2 && copy[x + 2][y - 2] === pion2) {
    return { x1: x + 1, y1: y - 1, x2: x + 2, y2: y - 2 }
  }



  return false
}

class Board extends Component {

  static propTypes = {}

  state = {
    turn :  false,
    board: createBoard(),
    player1: 0,
    player2: 0,
    winLine: '',
  }

  putADot = (x, y) => {
    const { turn, player1, player2 } = this.state
    let copy = this.state.board
    if (turn === true) {

      if (!forbidenMove(copy, x, y, 1)) { return }
      copy[x][y] = 1

      if (findLine(copy, x, y, 1)) {
        this.setState({winLine: 'player 2'})
        return;
      }
      const cord = findCapture(copy,x, y, 1, 2)
      if (cord) {
        console.log('blackDot');
        console.log(require('util').inspect(cord, { depth: null }));
        copy[cord.x1][cord.y1] =  0
        copy[cord.x2][cord.y2] =  0
        this.setState({ player2: player2 + 1 })
      }
      this.setState({ board: copy })
      this.setState({turn: !turn})
    } else {
      if (!forbidenMove(copy, x, y, 2)) { return }
      copy[x][y] = 2
      if (findLine(copy, x, y, 2)) {
        this.setState({winLine: 'player 1'})
        return;
      }
      const cord = findCapture(copy,x, y, 2, 1)
      console.log(cord);
      if (cord) {

        console.log(require('util').inspect(cord, { depth: null }))
        copy[cord.x1][cord.y1] =  0
        copy[cord.x2][cord.y2] =  0
        this.setState({ player1: player1 + 1 })

      }
      this.setState({ board: copy })
      this.setState({turn: !turn})
    }
  }

  render () {
    const { board, turn, player1, player2, winLine } = this.state

    return (
      <div className="game">
      <h1 className="h1">
      {turn === true ? "Player 2" : "Player 1" }
      </h1>
      {(player1 === 10 || player2 === 10) && ( <div className="message"> {player1 === 10 ? "Player 1 Win" : "Player 2 win" }</div>)}
      {winLine && ( <div className="message"> {winLine === 'player 1' ? "Player 1 Win" : "Player 2 win" }</div>)}
      <div className='score'>
        <div>
          {`Player 1 score = ${player1}`}
        </div>
        <div>
          {` Player 2 score = ${player2}`}
        </div>
      </div>
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
