const Node = require('../nodes/node');
const clients = require('../services/mock').clients;

const testNodeInstantiation = function() {
  const node = new Node(clients[0], 5000);
};

module.exports = {
  testNodeInstantiation,
};