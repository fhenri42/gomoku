import React, { Component, PropTypes } from 'react'
import { fromJS } from 'immutable'
import { connect } from 'react-redux'
import activation from '../../actions/activation.js'
import "./style.scss"
import { createBoard, forbidenMove, findLine, findCapture } from '../utils.js'

/*
RULE: http://jeuxdesociete.free.fr/jeux/jeu-gomoku.html
Capture: Captures are made by flanking a pair of the opponent’s stones
Free-threes: look example
win if 5 are aligne and no Capture possibole
win if 10 are remove from the board.
//last stand
//forbiden
*/

function mapDispatchToProps(dispatch) {
  return {
    activation: (data) => dispatch(activation(data)),
  }
}
  function  mapStateToProps(state) {
    const board = state.game.getIn(['board']).toJS()
  return {
    board
  }
}

export class BoardAi extends Component {

  static propTypes = {
    activation: PropTypes.func,
    board: PropTypes.Array,
  }

  state = {
    turn :  false,
    player1: 0,
    player2: 0,
    winLine: '',
  }

  newGame = () => {
    this.setState({ player1: 0 })
    this.setState({ player2: 0 })
    this.setState({ winLine: '' })
    this.setState({ board: createBoard() })
  }

  putADot = (x, y) => {
    const { activation, board } = this.props
    const { turn, player1, player2, winLine,  } = this.state
    let copy = board
    if(player1 != 10 && player2 != 10 && !winLine)  {

      if (!forbidenMove(copy, x, y, 1)) { return }
      copy[x][y] = 1

      if (findLine(copy, x, y, 1)) {
        this.setState({winLine: 'player 2'})
        return;
      }
      const cord = findCapture(copy,x, y, 1, 2)
      if (cord) {
        copy[cord.x1][cord.y1] =  0
        copy[cord.x2][cord.y2] =  0
        this.setState({ player2: player2 + 1 })
      }
      this.setState({ board: copy })
      this.setState({turn: !turn})

      activation({board: copy})
  }
  }
  render () {
    const { turn, player1, player2, winLine,  } = this.state
    const { board } = this.props
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
        {winLine && ( <button onClick={() => this.newGame()}> { 'play again'}</button>)}
        </div>
      )
    }
  }

export default connect(mapStateToProps, mapDispatchToProps)(BoardAi)
