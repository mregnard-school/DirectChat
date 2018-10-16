const net = require('net');
const chain = require('./middlewares/');

const handshakeInfo = 'handshake ';
const handshakeAck = handshakeInfo + 'ack';

class Node {
  constructor(client) {
    this.client = client;
    this.sockets = [];
    this.serverSocket = undefined;
    this.middleWareChain = chain;
  }
  
  runServer(port) {
    this.serverSocket = net.createServer((socket) => {
      this.onConnection(socket);
      this.setupListeners(socket);
    });
    
    this.serverSocket.listen(port);
  }
  
  onConnection(socket) {
    this.sockets.push(socket);
    socket.setEncoding('utf8');
  }
  
  parseAndSendHandShakeServerSide(socket, data) {
    socket.client = this.parseHandshake(data);
    socket.write(this.buildHandshakeMessage());
  }
  
  setupListeners(socket) {
    socket.on('data', (data) => {
      if (data.includes(handshakeInfo)) { // TODO irindul 2018-10-15 : Add boolean to ensure we handshake once
        this.parseAndSendHandShakeServerSide(socket, data);
      } else {
        this.onReceivedData(socket, data);
      }
    });
    
    socket.on('end', () => {
      this.onEndConnection(socket);
    })
  }
  
  buildHandshakeMessage() {
    return handshakeInfo + this.client.id + ' ' + this.client.pseudo;
  }
  
  parseHandshake(data) {
    let parts = data.split(' ');
    let id = Number(parts[1]);
    let pseudo = parts[2];
    return {
      "id": id,
      "pseudo": pseudo,
    };
  }
  
  onReceivedData(socket, data) {
    //socket.write(data);
  }
  
  onEndConnection(socket) {
    //  this.sockets = this.sockets.filter(item => item !== socket);
  }
  
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
        this.sockets.push(clientSocket);
        this.sendHandShakeClientSide(clientSocket);
      });
      
      clientSocket.on('data', (data) => {
        data = data.toString();
        if (data.includes(handshakeInfo)) {
          self.parseHandshakeClientSide(clientSocket, data);
          resolve();
        }
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
    socket.client = this.parseHandshake(data);
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