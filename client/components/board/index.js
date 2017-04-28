import React, { Component } from 'react'

import "./style.scss"

class Board extends Component {

  componentDidMount() {
    document.getElementById('inner').focus()
  }

  render () {
    const { board, putADot } = this.props

    return (
      <div className='board' onBlur={() => document.getElementById('inner').focus()} id="inner" tabIndex="-1">
        { board.map((line,x) => {
          return (
            <div className='line'>
              { line.map((bloc, y) => {

                  if (bloc === 0) { return(<div className='bloc'> { (x != 18 && y != 18) && (<div className='noDot' onClick={()=> putADot(x,y)}/>) } </div>) }

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
