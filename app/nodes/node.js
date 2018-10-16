const net = require('net');

class Node {
  constructor(client) {
    this.client = client;
    this.sockets = [];
    this.serverSocket = undefined;
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
    socket.write("Hello friend " + "\n");
  }
  
  setupListeners(socket) {
    socket.on('data', (data) => {
      this.onReceivedData(socket, data);
    });
    
    socket.on('end', () => {
      this.onEndConnection(socket);
    })
  }
  
  onReceivedData(socket, data) {
    socket.write(data);
  }
  
  onEndConnection(socket) {
    this.sockets = this.sockets.filter(item => item !== socket);
  }
  
  closeServer() {
    this.sockets.forEach(socket => {
      socket.destroy();
    });
    this.serverSocket.close();
  }
  
  connectTo(ip, port) {
    const clientSocket = new net.Socket();
    
    clientSocket.connect(port, ip);
    
    clientSocket.on('data', function(data) {
      clientSocket.destroy(); // kill client after server's response
    });
  
    clientSocket.on('close', function() {
    });
  
    clientSocket.on('error', (error) => {
      throw (error);
    });
  }
}

module.exports = Node;