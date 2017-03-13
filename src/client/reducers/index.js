import { fromJS } from 'immutable'

const intialState = {
  name: null,
}

const reducer = (state = fromJS(intialState) , action) => {
      return state
  }


export default reducer
