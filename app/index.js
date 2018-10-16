const Node = require('./nodes/node');
const clients = require('./services/mock/clients');
const node = new Node(clients[0]);

node.runServer(5000);