const net = require('net');
const moment = require('moment');
const chain = require('./middlewares/');
const ServerHandler = require('./network/connectionHandler').Server;
const ClientHandler = require('./network/connectionHandler').Client;

class ChangeableCallback {
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
    this.onEndConnection = () => {
    };
    this.onNewConnection = () => {
    };
    
    this.callbackHandler = new ChangeableCallback(() => {
    });
  }
  
  setOnReceiveData(onReceiveData) {
    this.callbackHandler.setCallback(onReceiveData);
    return this;
  }
  
  setOnEndConnection(onEndConnection) {
    this.onEndConnection = onEndConnection;
    return this;
  }
  
  setOnNewConnection(onNewConnection) {
    this.onNewConnection = (socket) => {
      onNewConnection(socket);
    };
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
    handler.setCallbackHandler(this.callbackHandler)
        .setOnConnectionClose((socket) => {
          this.sockets = this.sockets.filter(sock => sock !== socket);
          this.onEndConnection(socket)
        })
        .setOnError((error) => {
          if (reject) {
            reject(error);
          }
          else {
            console.log('Server already running, not restarting');
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
    const socket = net.createConnection(port, ip);
    this.sockets.push(socket);
    const handler = new ClientHandler(socket, this.client);
    this.initializeConnectionHandler(handler, () => {}, () => {});
  }
  
  writeRaw(client, message) {
    let content = this.applyMiddlewares(message);
    let sockets = this.socketsAssociatedWithClient(client);
    sockets.forEach((socket) => {
      socket.write(content);
    })
  }
  
  writeMessageTo(client, message) {
    const sockets = this.socketsAssociatedWithClient(client);
    const messageObject = this.constructMesage(message);
    sockets.forEach(socket => {
      socket.write(messageObject);
    });
  }
  
  constructMesage(message) {
    message.content = this.applyMiddlewares(message.content);
    return JSON.stringify(message);
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