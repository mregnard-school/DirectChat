const sinon = require('sinon');
const expect = require('chai').expect;

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
  
  it("should run server on port 5000 (will exit on test exit)", () => {
    const client = sinon.mock(clients[0]);
    const node = new Node(client);
    
    node.runServer(5000);
  });
});