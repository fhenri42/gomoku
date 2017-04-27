import React, { Component } from 'react'
import { createBoard, forbidenMove, findLine, findCapture, lastStand } from '../../utils/index.js'
import request from 'superagent'
import Score from '../Score'
import Board from '../Board'

import "./style.scss"

class Game extends Component {

  state = {
    turn :  false,
    player1: 0,
    player2: 0,
    winLine: '',
    board: createBoard(),
    isEatable: 0
  }

  nextTurn = (tab) => {
    return new Promise((resolve, reject) => {
      request
      .post(`http://localhost:5000/aiturn`)
      .set('Accept', 'application/json')
      .send({board})
      .end((err, res) => {
        if (err) { reject(err) }
        resolve(res.body)
      })
    })
  }

  newGame = () => {
    this.setState({player1: 0})
    this.setState({player2: 0})
    this.setState({winLine: ''})
    this.setState({board: createBoard()})
    this.setState({turn: false})
  }

  putADotSolo = (x, y) => {
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

  putADotMulti = (x, y) => {
    const { board, turn } = this.state
    board[x][y] = 2
    this.setState({ board })
    this.setState({turn: !turn})
    console.log('hello');
    this.nextTurn(board).then(res => {
      console.log(res);
      //whatever

    })
  //   const { turn, player1, player2, winLine } = this.state
  //   let copy = this.props.board
  //   const { boardSet, SetisEatable, isEatable } = this.props
  //   if(player1 != 10 && player2 != 10 && !winLine)  {
  //   if (turn === true) {
  //
  //     if (!forbidenMove(copy, x, y, 1)) { return }
  //     copy[x][y] = 1
  //
  //     if (findLine(copy, 1, SetisEatable, isEatable)) {
  //       console.log("LA");
  //       this.setState({winLine: 'player 2'})
  //       return;
  //     }
  //     const cord = findCapture(copy,x, y, 1, 2)
  //     if (cord) {
  //       copy[cord.x1][cord.y1] =  0
  //       copy[cord.x2][cord.y2] =  0
  //       this.setState({ player2: player2 + 1 })
  //     }
  //     this.setState({turn: !turn})
  //     boardSet(copy)
  //   } else {
  //     if (!forbidenMove(copy, x, y, 2)) { return }
  //     copy[x][y] = 2
  //
  //     if (findLine(copy, 2, SetisEatable, isEatable)) {
  //       this.setState({winLine: 'player 1'})
  //       return;
  //     }
  //     const cord = findCapture(copy,x, y, 2, 1)
  //     if (cord) {
  //       copy[cord.x1][cord.y1] =  0
  //       copy[cord.x2][cord.y2] =  0
  //       this.setState({ player1: player1 + 1 })
  //
  //     }
  //     this.setState({turn: !turn})
  //     boardSet(copy)
  //   }
  // }

  }

  render () {
    const { turn, player1, player2, winLine, board } = this.state
    const { nbPlayer } = this.props

    return (
      <div className="game">
        <Score player1={player1} player2={player2} winLine={winLine} turn={turn}/>
        <Board board={board} putADot={nbPlayer == 1 ? this.putADotMulti : this.putADotSolo} />
        { winLine && (<button onClick={() => this.newGame()}>{ 'play again' }</button>) }
      </div>
    )
  }
}

export default Game

/*
RULE: http://jeuxdesociete.free.fr/jeux/jeu-gomoku.html
Capture: Captures are made by flanking a pair of the opponent’s stones
Free-threes: look example
win if 5 are aligne and no Capture possibole
win if 10 are remove from the board.
//last stand
//forbiden
*/
