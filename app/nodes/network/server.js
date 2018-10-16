const Connector = require('./handshakeHandler');

class Server extends Connector {
  
  constructor(socket, client) {
    super(socket, client);
    
  }
  
  handleHandshake(data) {
    super.handleHandshake(data);
    this.writeHandshake();
  }
}

module.exports = Server;