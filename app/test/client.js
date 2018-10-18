const sinon = require('sinon');
const expect = require('chai').expect;
const clients = require('../services/mock').clients;

const Client = require('../client/client');

describe('Client', () => {
  it('should instantiate a new Client', () => {
    const mock = clients[0];
    const client = new Client(mock);
    expect(client).to.have.property('node');
    expect(client).to.have.property('client');
    expect(client).to.have.property('ipTable');
  });
  
  it('should map an ip:port string to an object', () => {
    const mock = clients[0];
    const client = new Client(mock);
    
    const ipObject = client.parseIpAndPortFromString("127.0.0.1:80");
    
    expect(ipObject).to.have.property('ip');
    expect(ipObject).to.have.property('port');
    
    expect(ipObject.ip).to.equal('127.0.0.1');
    expect(ipObject.port).to.equal(80);
  });
  
  it('should add a friend\'s ip to the ip table', () => {
    const mockClient = clients[0];
    const mockFriend = clients[1];
    const client = new Client(mockClient);
    
    const friendIp = mockFriend.ips[0];
    client.addFriendIpToTable(mockFriend, friendIp);
    
    expect(client.ipTable)
  });
  
  it('should map all friends to the ip table', () => {
    const mockClient = clients[0];
    const mockFriends = [
        clients[1],
        clients[2],
    ];
    
    const client = new Client(mockClient);
    const stub = sinon.stub(client, 'addFriendIpToTable');
    
    client.mapFriendsToIPs(mockFriends);
    
    expect(stub.calledTwice).to.be.true;
  });
});