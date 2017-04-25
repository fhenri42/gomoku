import _ from 'lodash'

export default (socket) => {
  socket.on('activation',(data) => aiTurn(data,socket))
}

const aiTurn = (data, socket) => {
  socket.emit('action', {
       type: 'AITURN',
       data,
     })
}
