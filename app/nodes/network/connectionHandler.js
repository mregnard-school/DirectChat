const ClientHandshake = require('./handshakeHandler').Client;
const ServerHandshake = require('./handshakeHandler').Server;

class ConnectionHandler {
  constructor(socket, client) {
    this.socket = socket;
    this.client = client;
    this.onConnectionClose = () => {};
    this.onError = () => {};
    
    this.callbackHandler = {};
  }
  
  setCallbackHandler(callbackHandler) {
    this.callbackHandler = callbackHandler;
    return this;
  }
  
  setOnConnectionClose(onConnectionClose) {
    this.onConnectionClose = onConnectionClose;
    return this;
  }
  
  setOnError(onError) {
    this.onError = onError;
    return this;
  }
  
  handleOnConnection(resolve) {
    this.socket.setEncoding("utf8");
    this.setupListeners(resolve);
  }
  
  setupListeners(resolve) {
    this.socket.on('data', (data) => {
      this.handleHandshake(data, resolve);
    });
    this.socket.on('end', () => {
      this.onConnectionClose(this.socket);
    });
    this.socket.on('error', this.onError);
  }
  
  handleHandshake(data, resolve) {
    const handshakeHandler = this.getHandshakeHandler();
    handshakeHandler.handleData(data, this.callbackHandler, resolve);
  }
  
  getHandshakeHandler() {
    return undefined;
  }
}

class Server extends ConnectionHandler {
  constructor(socket, client) {
    super(socket, client);
    this.handshakeHandler = new ServerHandshake(socket, client);
  }
  
  getHandshakeHandler() {
    return this.handshakeHandler
  }
}

class Client extends ConnectionHandler {
  constructor(socket, client) {
    super(socket, client);
    this.handshakeHandler = new ClientHandshake(socket, client);
  }
  
  handleOnConnection(resolve) {
    super.handleOnConnection(resolve);
    this.handshakeHandler.writeHandshake();
  }
  
  getHandshakeHandler() {
    return this.handshakeHandler;
  }
}

module.exports = {
  Server,
  Client,
};