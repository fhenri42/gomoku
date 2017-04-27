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
        <header>
          <h1> <a href="http://localhost:8080/">{'Gomoku'}</a></h1>
        </header>

        <div className='content'>
        {!chose && (
          <div className='game'>
            <button onClick={() => this.handeClic(2)}><p>{'1 vs 1' }</p></button>
            <button onClick={() => this.handeClic(1)}><p>{'vs the ai'}</p></button>
          </div>
        )}

        { chose && (<Game nbPlayer={numberPlayer}/>) }
        </div>

        <footer>
          <h1> <a href="http://localhost:8080/">{'Swag'}</a></h1>
        </footer>
      </div>
    )
  }

}

export default App
