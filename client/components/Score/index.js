import React, { Component } from 'react'
import "./style.scss"

export class Score extends Component {

  render () {
    const { player1, player2, winLine, turn } = this.props
    return (
      <div>
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
      </div>
    )
  }
}

export default Score
