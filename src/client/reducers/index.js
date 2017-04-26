import { fromJS } from 'immutable'
import { createBoard } from '../components/utils.js'
const intialState = {
  name: null,
  board: createBoard()
}

const formatTAb = (tab) => {
      let newTab = []
      let tmpNewBoard = tab.split("")
      let m = 0
      for (let t = 0; t <= 19; t++) {
        let tmp = []
        for (let b = 0; b <= 19; b++) {
          tmp.push(parseInt(tmpNewBoard[m]))
          m++
        }
        newTab.push(tmp)
      }
  return newTab
}
const reducer = (state = fromJS(intialState) , action) => {

  switch(action.type) {
      case 'START_GAME':
        return state.setIn(['room'], fromJS(action.data)).setIn(['currentPiece'], state.getIn(['room', 'pieces', 0]))
      case 'NEXT_TURN' :
        return state.setIn(['room'], fromJS(action.data)).setIn(['currentPiece'], state.getIn(['room', 'pieces', 0]))
      case 'AITURN':
        return state.setIn(['board'],fromJS(formatTAb(action.data)))
        case 'client/boardSet':
          return state.setIn(['board'],fromJS(action.data))
      default:
        return state
    }
  }

export default reducer
