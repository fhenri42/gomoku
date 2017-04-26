import React, { Component } from 'react'
import Game from '../components/Game'

import './style.scss'

class App extends Component {

  state = {
    chose :  false,
    numberPlayer: 0,
  }
  handeClic = (numberPlayer) => {
    this.setState({ chose: true })
    this.setState({ numberPlayer })
  }

  render () {
    const { numberPlayer, chose } =  this.state
    return (
      <div className='container'>
        <div className='content'>
        {!chose && (
          <div className='game'>
            <button onClick={() => this.handeClic(2)}>{'1 vs 1' }</button>
            <button onClick={() => this.handeClic(1)}>{'vs the ai'}</button>
          </div>
        )}

        { chose && (<Game nbPlayer={numberPlayer}/>) }
        </div>
      </div>
    )
  }

}

export default App
