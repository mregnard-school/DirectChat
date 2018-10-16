const handshakeInfo = 'handshake ';

class HandshakeHandler {
  constructor(socket, client) {
    this.socket = socket;
    this.client = client;
    this.handshaked = false;
  }
  
  handleData(data, onReceiveData, resolve) {
    if(!this.handshaked && data.includes(handshakeInfo)) {
      this.handleHandshake(data);
      if(resolve) {
        resolve();
      }
    } else {
      onReceiveData();
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
    this.socket.write(this.buildHandshakeMessage());
  }
  
  buildHandshakeMessage() {
    return handshakeInfo + this.client.id + ' ' + this.client.pseudo;
  }
}

module.exports = HandshakeHandler;