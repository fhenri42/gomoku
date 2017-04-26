import React, { Component } from 'react'
import Board from '../components/board'
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

        {(chose && numberPlayer === 2) && (<Board/>)}
        {(chose && numberPlayer === 1) && (<Board/>)}
        </div>
      </div>
    )
  }

}

export default App
