const net = require('net');
const chain = require('./middlewares/');
const ServerHandler = require('./network/connectionHandler').Server;
const ClientHandler = require('./network/connectionHandler').Client;

class Callback {
  constructor(callback) {
    this.callback = callback;
  }
  
  setCallback(callback) {
    this.callback = callback;
  }
  
  getCallback() {
    return this.callback;
  }
}

class Node {
  constructor(client) {
    this.client = client;
    this.sockets = []; //Maybe change this with hashmap client/socket
    this.serverSocket = undefined;
    this.middleWareChain = chain;
    
    this.onReceivedData = () => {};
    this.onEndConnection = () => {
    };
    this.onNewConnection = () => {
    };
  
    this.callbackHandler = new Callback(this.onReceivedData);
  }
  
  setOnReceiveData(onReceiveData) {
    this.onReceivedData = onReceiveData;
    this.callbackHandler.setCallback(onReceiveData);
    return this;
  }
  
  setOnEndConnection(onEndConnection) {
    this.onEndConnection = onEndConnection;
    return this;
  }
  
  setOnNewConnection(onNewConnection) {
    this.onNewConnection = onNewConnection;
    return this;
  }
  
  runServer(port) {
    this.serverSocket = net.createServer((socket) => {
      this.sockets.push(socket);
      const serverConnectionHandler = new ServerHandler(socket,
          this.client);
      this.initializeConnectionHandler(serverConnectionHandler, () => {
        this.onNewConnection(socket);
      });
      
    });
    
    this.serverSocket.listen(port);
  }
  
  initializeConnectionHandler(handler, resolve, reject) {
    handler.setOnReceiveData(this.callbackHandler)
        .setOnConnectionClose(this.onEndConnection)
        .setOnError((error) => {
          if (reject) {
            reject(error);
          }
          else {
            console.error(error);
          }
        })
        .handleOnConnection(resolve);
  }
  
  closeServer() {
    this.sockets.forEach(socket => {
      socket.destroy();
    });
    this.serverSocket.close();
  }
  
  connectTo(ip, port) {
    const clientSocket = new net.Socket();
    return new Promise((resolve, reject) => {
      clientSocket.connect(port, ip, () => {
        this.sockets.push(clientSocket);
        const handler = new ClientHandler(clientSocket, this.client);
        this.initializeConnectionHandler(handler, resolve, reject);
      });
    });
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
}

module.exports = Node;