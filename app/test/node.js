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
    nodeClient.connectTo('127.0.0.1', 5000).then(() => {
          nodeServer.closeServer();
          done();
        }
    )
  });
  
  it("should return the socket for associated client", (done) => {
    const nodeServer = new Node(clients[0]);
    nodeServer.runServer(5000);
    const nodeClient = new Node(clients[1]);
    nodeClient.connectTo("127.0.0.1", 5000)
        .then(() => {
              let sockets = nodeServer.socketsAssociatedWithClient(clients[1]);
              expect(sockets.length).to.deep.equal(1);
              nodeServer.closeServer();
              done();
            }
        );
  });
  
  it("server should send a message to a client", (done) => {
    const nodeServer = new Node(clients[0]);
    const nodeClient = new Node(clients[1]);
    const receiveStub = sinon.stub(nodeClient, 'onReceivedData');
    
    nodeServer.runServer(5000);
    
    nodeClient.connectTo('127.0.0.1', 5000).then(() => {
      nodeServer.writeRaw(clients[1], "Hello");
      expect(receiveStub.calledOnce);
      nodeServer.closeServer();
      done();
    });
  });
  
  // TODO irindul 2018-10-20 : Redo - Adapt to changes
  it("client should send a message to server", (done) => {
    const nodeServer = new Node(clients[0]);
    const nodeClient = new Node(clients[1]);
    const receiveStub = sinon.stub(nodeServer.callbackHandler, 'getCallback');
    
    nodeServer.runServer(5000);
    
    nodeClient.connectTo('127.0.0.1', 5000).then(() => {
      nodeClient.writeRaw(clients[0], "Hello");
      expect(receiveStub.calledOnce);
      nodeServer.closeServer();
      done();
    })
  });
  
});