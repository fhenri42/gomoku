import React, { Component } from 'react'

import "./style.scss"

class Board extends Component {

  render () {
    const { board, putADot } = this.props

    return (
      <div>
        { board.map((line,x) => {
          return (
            <div className='line'>
              { line.map((bloc, y) => {

                  if (bloc === 0) { return(<div className='bloc'> <div className='noDot' onClick={()=> putADot(x,y)}/> </div>) }

                  else if (bloc === 1) { return(< div className='bloc'> <div className='blackDot'/> </div>) }

                  else { return(<div className='bloc'><div className='whiteDot'/></div>) }

                })
              }
            </div>
          )}
        )}
      </div>
    )
  }
}

export default Board

/*
RULE: http://jeuxdesociete.free.fr/jeux/jeu-gomoku.html
Capture: Captures are made by flanking a pair of the opponent’s stones
Free-threes: look example
win if 5 are aligne and no Capture possibole
win if 10 are remove from the board.
//last stand
//forbiden
*/
