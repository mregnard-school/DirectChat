const net = require('net');
const chain = require('./middlewares/');

const ServerConnectionHandler = require('./network/serverConnectionHandler');
const ClientConnectionHandler = require('./network/clientConnectionHandler');

class Node {
  constructor(client) {
    this.client = client;
    this.sockets = []; //Maybe change this with hashmap client/socket
    this.serverSocket = undefined;
    this.middleWareChain = chain;
    
  }
  
  //Server-side code
  runServer(port) {
    this.serverSocket = net.createServer((socket) => {
      this.sockets.push(socket);
      const serverConnectionHandler = new ServerConnectionHandler(socket, this.client);
      serverConnectionHandler
          .setOnReceiveData(this.onReceivedData)
          .setOnConnectionClose(this.onEndConnection)
          .setOnError((error) => console.log(error))
          .handleOnConnection();
    });
    
    this.serverSocket.listen(port);
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
        this.sockets.push(clientSocket);
        const handler = new ClientConnectionHandler(clientSocket, this.client);
        handler
            .setOnReceiveData(this.onReceivedData)
            .setOnConnectionClose(this.onEndConnection)
            .setOnError((error) => { reject(error)})
            .handleOnConnection(resolve);
      });
    });
    
    return promise;
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