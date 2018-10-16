const net = require('net');
const chain = require('./middlewares/');

const handshakeInfo = 'handshake ';

class Node {
  constructor(client) {
    this.client = client;
    this.sockets = [];
    this.serverSocket = undefined;
    this.middleWareChain = chain;
  }
  
  //Server-side code
  runServer(port) {
    this.serverSocket = net.createServer((socket) => {
      this.onConnection(socket);
      this.setupListeners(socket);
    });
    
    this.serverSocket.listen(port);
  }
  
  onConnection(socket, isClient) {
    this.sockets.push(socket);
    socket.setEncoding('utf8');
    if(isClient) {
      this.sendHandShakeClientSide(socket);
    }
  }
  
  setupListeners(socket) {
    socket.on('data', (data) => {
      this.handleData(socket, data);
    });
    
    socket.on('end', () => {
      this.onEndConnection(socket);
    })
  }
  
  handleData(socket, data, isClient, callback) {
    if (data.includes(handshakeInfo)) { // TODO irindul 2018-10-15 : Add boolean to ensure we handshake once
      if(isClient) {
        this.parseHandshakeClientSide(socket, data);
        callback();
      } else {
        this.parseAndSendHandShakeServerSide(socket, data);
      }
    } else {
      this.onReceivedData(socket, data);
    }
  }
  
  parseAndSendHandShakeServerSide(socket, data) {
    socket.client = this.parseClientFromHandshake(data);
    socket.write(this.buildHandshakeMessage());
  }
  
  parseClientFromHandshake(data) {
    let parts = data.split(' ');
    let id = Number(parts[1]);
    let pseudo = parts[2];
    return {
      "id": id,
      "pseudo": pseudo,
    };
  }
  
  buildHandshakeMessage() {
    return handshakeInfo + this.client.id + ' ' + this.client.pseudo;
  }
  
  onReceivedData(socket, data) {}
  
  onEndConnection(socket) {}
  
  closeServer() {
    this.sockets.forEach(socket => {
      socket.destroy();
    });
    this.serverSocket.close();
  }
  
  connectTo(ip, port) {
    const clientSocket = new net.Socket();
    let self = this;
    let promise = new Promise((resolve, reject) => {
      clientSocket.connect(port, ip, () => {
        this.onConnection(clientSocket, true);
      });
      
      clientSocket.on('data', (data) => {
        data = data.toString(); //Is Byte buffer otherwise;
        this.handleData(clientSocket, data, true, resolve);
        //this.onReceivedData(clientSocket, data);
      });
      
      clientSocket.on('close', function () {
      });
      
      clientSocket.on('error', (error) => {
        reject(error);
      });
    });
    
    return promise;
  }
  
  parseHandshakeClientSide(socket, data) {
    socket.client = this.parseClientFromHandshake(data);
  }
  
  writeMessageTo(client, message) {
    let content = this.applyMiddlewares(message);
    let sockets = this.socketsAssociatedWithClient(client);
    sockets.forEach(socket => {
      socket.write(content); // TODO irindul 2018-10-16 : Construct message object here (date, author, content etc..)
    });
  }
  
  applyMiddlewares(message) {
    return this.middleWareChain.applyMiddlewares(message);
  }
  
  socketsAssociatedWithClient(client) {
    let socketsAssociated = [];
    for (let i = 0; i < this.sockets.length; i++) {
      if (this.sockets[i].client.id === client.id) {
        socketsAssociated.push(this.sockets[i]);
      }
    }
    
    return socketsAssociated;
  }
  
  sendHandShakeClientSide(socket) {
    socket.write(this.buildHandshakeMessage());
  }
}

module.exports = Node;