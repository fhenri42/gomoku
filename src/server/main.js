import params  from '../../params'

import cors from 'cors'
import express from 'express'
import bodyParser from 'body-parser'
import event from './event/index.js'

const app = express()
var server = require('http').createServer(app);


const env = process.env.NODE_ENV || 'development'
const config = params
require("babel-core/register");
require("babel-polyfill");

app.use(cors())
app.use(bodyParser.json())
app.use(bodyParser.urlencoded({ extended: true }))
global.io = require('socket.io')(server)




app.get('/',(req, res) => res.end())

io.on('connection', (socket) => {
   	event(socket)
 })


 server.listen(5000, function(){
   console.log('listening on 5000');
 });
