const net = require('net');

class Node {
  constructor(client) {
    this.client = client;
    this.sockets = [];
  }
  
  runServer(port) {
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