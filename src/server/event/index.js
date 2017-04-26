import _ from 'lodash'
import cp from 'child_process'
import fs from 'fs'

export default (socket) => {
  socket.on('activation',(data) => aiTurn(data,socket))
}


const aiTurn = (data, socket) => {

try {

  let test = cp.spawn("ls",[])
  test.stdout.on('data',(data) => {
    console.log(data.toString());
  })
  test.stderr.on('data', (data) => {
  console.log(`stderr: ${data}`);
})
console.log(data.board);
let goScript = cp.exec(`src/Algo/main ${data.board}`, (err, stdout, stderr) => {
  if (err) {
    console.log(err);
    return
  }
  console.log(stdout.toString());

  fs.readFile('tabToSend', 'utf8', function (err,data) {
    if (err) {
      return console.log(err);
    }
    console.log(data);
     socket.emit('action', {
       type: 'AITURN',
       data,
     })
  });
})
} catch (e){
  console.log(e);
}
}
