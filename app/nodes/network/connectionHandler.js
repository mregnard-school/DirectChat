const Server = require('./server');

class ConnectionHandler {
  constructor(socket, client) {
    this.socket = socket;
    this.client = client;
    this.onReceiveData = () => {};
    this.onConnectionClose = () => {};
    this.onError = () => {};
  }
  
  setOnReceiveData(onReceiveData) {
    this.onReceiveData = onReceiveData;
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
    this.socket.on('end', this.onConnectionClose);
    this.socket.on('error', this.onError);
  }
  
  handleHandshake(data, resolve) {
    const handshakeHandler = this.getHandshakeHandler();
    handshakeHandler.handleData(data, this.onReceiveData, resolve);
  }
  
  getHandshakeHandler() {
    return undefined;
  }
}

module.exports = ConnectionHandler;