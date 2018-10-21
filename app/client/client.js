const Node = require('../nodes/node');
const HashTable = require('../services/util').HashTable;
const parseIp = require('../services/util').parseIpAndPortFromString;

class Client {
  constructor(client) {
    this.ipTable = new HashTable();
    this.client = client;
    this.node = new Node(client);
  }
  
  
  
  mapFriendsToIPs(friends) {
    friends.forEach((friend) => {
      friend.ips.forEach(ip => {
        this.addFriendIpToTable(friend, ip);
      });
    });
  }
  
  addFriendIpToTable(friend, ip) {
    let ipEntry = parseIp(ip);
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