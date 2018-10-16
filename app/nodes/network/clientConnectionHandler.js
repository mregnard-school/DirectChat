const ConnectionHandler = require('./connectionHandler');
const Client = require('./client');
class ClientConnectionHandler extends ConnectionHandler{
  
  constructor(socket, client) {
    super(socket, client);
    this.handshakeHandler = new Client(socket, client);
  }
  
  handleOnConnection(resolve) {
    super.handleOnConnection(resolve);
    this.handshakeHandler.writeHandshake();
  }
  
  getHandshakeHandler() {
    return this.handshakeHandler;
  }
}

module.exports = ClientConnectionHandler;