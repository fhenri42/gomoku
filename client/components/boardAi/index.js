import React, { Component, PropTypes } from 'react'
import { createBoard, forbidenMove, findLine, findCapture } from '../../utils.js'
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

export class BoardAi extends Component {

  state = {
    turn :  false,
    player1: 0,
    player2: 0,
    winLine: '',
    board: createBoard()
  }

  newGame = () => {
    this.setState({ player1: 0 })
    this.setState({ player2: 0 })
    this.setState({ winLine: '' })
    this.setState({ turn: false })
    this.setState({ board: createBoard() })
  }

  putADot = (x, y) => {
    const { turn, player1, player2, winLine, board } = this.state
    if (player1 != 10 && player2 != 10 && !winLine) {

      if (!forbidenMove(board, x, y, 1)) { return }
      board[x][y] = 1

      if (findLine(board, x, y, 1)) {
        this.setState({ winLine: 'player 2' })
        return;
      }
      const coord = findCapture(board, x, y, 1, 2)
      if (coord) {
        board[coord.x1][coord.y1] = 0
        board[coord.x2][coord.y2] = 0
        this.setState({ player2: player2 + 1 })
      }
      this.setState({ board })
      this.setState({ turn: !turn })
    }
  }
  render () {
    const { turn, player1, player2, winLine, board } = this.state
    return (
      <div className="game">

        <h1 className="h1">
          { turn === true ? "Player 2" : "Player 1" }
        </h1>

        {(player1 === 10 || player2 === 10 || winLine) && (
          <div className="message">
            { (player1 === 10 || winLine === 'player 1') ? "Player 1 Win" : "Player 2 win" }
          </div>
        )}

        <div className='score'>
          <div> {`Player 1 score = ${player1}`} </div>
          <div> {` Player 2 score = ${player2}`} </div>
        </div>

        { board.map((line,x) => {
          return (
            <div className='line'>
              { line.map((bloc, y) => {

                  if (bloc === 0) { return(<div className='bloc'> <div className='noDot' onClick={()=> this.putADot(x,y)}/> </div>) }

                  else if (bloc === 1) { return(< div className='bloc'> <div className='blackDot'/> </div>) }

                  else { return(<div className='bloc'><div className='whiteDot'/></div>) }

                })
              }
            </div>
          )}
        )}

        { winLine && (<button onClick={() => this.newGame()}>{ 'play again' }</button>) }

      </div>
    )
  }
}

export default BoardAi
