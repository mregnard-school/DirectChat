const handshakeInfo = 'handshake';

class HandshakeHandler {
  constructor(socket, client) {
    this.socket = socket;
    this.client = client;
    this.handshaked = false;
  }
  
  handleData(data, callbackHandler, resolve) {
    if (!this.handshaked) {
      try {
        const dataObj = JSON.parse(data);
        if (dataObj.type === handshakeInfo) {
          this.handleHandshake(dataObj);
          if (resolve) {
            resolve();
          }
        }
      } catch (e) {
        //Do not handle message if no handshake has been made
      }
    } else {
      let callback = callbackHandler.getCallback();
      callback(data);
    }
  }
  
  handleHandshake(data) {
    this.socket.client = HandshakeHandler.parseClientFromHandshake(data);
    console.log(this.socket.client);
    this.handshaked = true;
  }
  
  static parseClientFromHandshake(data) {
   return data.friend;
  }
  
  writeHandshake() {
    this.socket.write(this.buildHandshakeMessage());
  }
  
  buildHandshakeMessage() {
    const message = {
      type: 'handshake',
      friend: this.getClientWithoutFriends(),
    };
    
    return JSON.stringify(message);
  }
  
  getClientWithoutFriends() {
    return {
      id: this.client.id,
      pseudo: this.client.pseudo,
      ips: this.client.ips,
    }
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