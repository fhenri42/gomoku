import { fromJS } from 'immutable'

const intialState = {
  name: null,
}

const reducer = (state = fromJS(intialState) , action) => {
  
  switch(action.type) {
      case 'START_GAME':
        return state.setIn(['room'], fromJS(action.data)).setIn(['currentPiece'], state.getIn(['room', 'pieces', 0]))
      case 'NEXT_TURN' :
        return state.setIn(['room'], fromJS(action.data)).setIn(['currentPiece'], state.getIn(['room', 'pieces', 0]))
      default:
        return state
    }
  }


export default reducer
