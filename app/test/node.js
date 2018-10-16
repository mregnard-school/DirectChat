const sinon = require('sinon');
const expect = require('chai').expect;
const assert = require('chai').assert;

const Node = require('../nodes/node');
const clients = require('../services/mock').clients;

describe("Node", () => {
  beforeEach(() => {
    this.sandbox = sinon.createSandbox();
  });
  
  afterEach(() => {
    this.sandbox.restore();
  });
  
  it("should instantiate node", () => {
    const client = sinon.mock(clients[0]);
    
    const node = new Node(client);
    expect(node).to.have.deep.property('client', client);
    expect(node).to.have.deep.property('sockets', []);
  });
  
  it("should run server on port 5000", () => {
    const client = sinon.mock(clients[0]);
    const node = new Node(client);
    
    node.runServer(5000);
    node.closeServer();
  });
  
  it("should connect to another peer", (done) => {
    const nodeServer = new Node(clients[0]);
    nodeServer.runServer(5000);
    
    const nodeClient = new Node(clients[0]);
    nodeClient.connectTo('127.0.0.1', 5000);
    
    setTimeout(() => {
      nodeServer.closeServer();
      done();
    }, 1000)
  });
});