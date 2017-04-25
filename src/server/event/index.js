require('node-go-require')
import _ from 'lodash'


export default (socket) => {
  socket.on('activation',(data) => aiTurn(data,socket))
}

const aiTurn = (data, socket) => {

try {

  var petGo = require("./main.go");

  // var pet = petGo.pet.New('my pet');
  // console.log(pet.Name());
  // pet.SetName('new name...');
  // console.log(pet.Name());
} catch (e){
  console.log(e);
}
  socket.emit('action', {
       type: 'AITURN',
       data,
     })
}
