const Connector = require('./handshakeHandler');

class Client extends Connector {
  constructor(socket, client) {
    super(socket, client);
  }
}

module.exports = Client;