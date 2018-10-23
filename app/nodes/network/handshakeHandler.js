const handshakeInfo = 'handshake ';

class HandshakeHandler {
  constructor(socket, client) {
    this.socket = socket;
    this.client = client;
    this.handshaked = false;
  }
  
  handleData(data, callbackHandler, resolve) {
    if(!this.handshaked && data.includes(handshakeInfo)) {
      this.handleHandshake(data);
      if(resolve) {
        resolve();
      }
    } else {
      let callback = callbackHandler.getCallback();
      callback(data);
    }
  }
  
  handleHandshake(data) {
    this.socket.client = HandshakeHandler.parseClientFromHandshake(data);
    this.handshaked = true;
  }
  
  static parseClientFromHandshake(data) {
    let parts = data.split(' ');
    let id = Number(parts[1]);
    let pseudo = parts[2];
    return {
      "id": id,
      "pseudo": pseudo,
    };
  }
  
  writeHandshake() {
    console.log('writing handshake');
    this.socket.write(this.buildHandshakeMessage());
  }
  
  buildHandshakeMessage() {
    return handshakeInfo + this.client.id + ' ' + this.client.pseudo + '\n';
  }
}

class Client extends HandshakeHandler {
  constructor(socket, client) {
    super(socket, client);
  }
}

class Server extends HandshakeHandler {
  constructor(socket, client) {
    super(socket, client);
    
  }
  
  handleHandshake(data) {
    super.handleHandshake(data);
    this.writeHandshake();
  }
}

module.exports = {
  Server,
  Client,
};