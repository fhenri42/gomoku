
import express from 'express'
import event from './event/index.js'
import socketio from 'socket.io'
import http from 'http'

const app = express()
const server = http.createServer(app)
const io = socketio(server)

app.get('/',(req, res) => res.end())
io.on('connection', (socket) => event(socket))
server.listen(5000, () => console.log('listening on 5000'))
