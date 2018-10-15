const net = require('net');
const clients = require('../services/mock').clients;

class Node {
  constructor() {
    this.client = clients[0];
    this.sockets = [];
    
    this.initializeServer(5000); // TODO irindul 2018-10-15 : Don't hardcode it
  }
  
  initializeServer(port) {
    net.createServer((socket) => {
      this.onConnection(socket);
      
      socket.on('data', (data) => {
        this.onReceivedData(socket, data);
      });
      
      socket.on('end', () => {
        this.onEndConnection(socket);
      })
    }).listen(port);
  }
  
  onConnection(socket) {
    this.sockets.push(socket);
    socket.write("Hello friend" + "\n");
  }
  
  onReceivedData(socket, data) {
    console.log("Received : " + data);
    socket.write(data);
  }
  
  onEndConnection(socket) {
    this.sockets.remove(socket);
    console.log('Client disconected');
  }
}

module.exports = Node;