const ConnectionHandler = require('./connectionHandler');
const Server = require('./server');
class ServerConnectionHandler extends ConnectionHandler {
  
  constructor(socket, client) {
    super(socket, client);
    this.handshakeHandler = new Server(socket, client);
  }
  
  getHandshakeHandler() {
    return this.handshakeHandler
  }
}

module.exports = ServerConnectionHandler;