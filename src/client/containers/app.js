import React, { Component, PropTypes } from 'react'
import { connect } from 'react-redux'
import Board from '../components/board'

import './style.scss'

const mapDispatchToProps = (dispatch) => {
  return

}

class App extends Component {

  static propTypes = {  }



  render () {
    return (
      <div className='container'>
        <div className='content'>
          <Board />
        </div>
      </div>
    )
  }

}

export default connect(null, mapDispatchToProps)(App)
