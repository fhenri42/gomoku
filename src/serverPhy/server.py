#!/usr/bin/env python

import tornado.ioloop
import tornado.web
import tornado.websocket
import tornado.template
#from event import eventHandeler
from event import onEvent

class WSHandler(tornado.websocket.WebSocketHandler):
  def open(self):
    print 'connection opened...'
    self.write_message("The server says: 'Hello'. Connection was accepted.")

  def on_message(self, data):
    self.write_message("The server says: " + data + " back at you")
    #eventHandeler.onEvent(message)
    onEvent(self,data)

  def on_close(self):
    print 'connection closed...'

  def check_origin(self, origin):
    return True

application = tornado.web.Application([
  (r'/ws', WSHandler),
  (r"/(.*)", tornado.web.StaticFileHandler, {"path": "./resources"}),
])

if __name__ == "__main__":
  application.listen(5000)
  tornado.ioloop.IOLoop.instance().start()
  print 'lol'
