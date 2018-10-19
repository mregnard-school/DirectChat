const Node = require('../nodes/node');
const HashTable = require('../services/util').HashTable;

class Client {
  constructor(client) {
    this.ipTable = new HashTable();
    this.client = client;
    this.node = new Node(client);
  }
  
  parseIpAndPortFromString(ipString) {
    let splited = ipString.split(':');
    return {
      "ip": splited[0],
      "port": parseInt(splited[1]),
    }
  }
  
  mapFriendsToIPs(friends) {
    friends.forEach((friend) => {
      friend.ips.forEach(ip => {
        this.addFriendIpToTable(friend, ip);
      });
    });
  }
  
  addFriendIpToTable(friend, ip) {
    let ipEntry = this.parseIpAndPortFromString(ip);
    this.ipTable.put(friend, ipEntry);
  }
  
  connectToEachClient() {
    this.ipTable.forEach((key, value) => {
      this.node.connectTo(value.ip, value.port);
    });
  }
  
  runServer(port) {
    this.node.runServer(port);
  }
  
  setOnReceiveData(onReceiveData) {
    this.node.setOnReceiveData(onReceiveData);
    return this;
  }
  
  setOnEndConnection(onEndConnection) {
    this.node.setOnEndConnection(onEndConnection);
    return this;
  }
  
  setOnNewConnection(onNewConnection) {
    this.node.setOnNewConnection(onNewConnection);
  }
}

module.exports = Client;