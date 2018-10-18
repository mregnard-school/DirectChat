const sinon = require('sinon');
const expect = require('chai').expect;
const HashTable = require('../services/util').HashTable;

describe('Utility methods', () => {
  it('should create a new HashTable', () => {
    let table = new HashTable();
    expect(table).to.have.property('map');
  });
  
  it('should map a key to a value', () => {
    let mock = {
      id: 1,
      value: 'Hello',
    };
    let mock2 = {
      id: 2,
      value: 'Foo'
    };
    
    let table = new HashTable();
    table.put(mock, "value");
    expect(Object.keys(table.map).length).to.equal(1);
    
    table.put(mock2, "value2");
    expect(Object.keys(table.map).length).to.equal(2);
  });
  
  it('should map a key to a value and retrieve it', () => {
    let table = new HashTable();
    let key = {
      id: 1,
      value: "Hello",
    };
    table.put(key, "value");
    expect(table.get(key)).to.equal("value");
  });
  
  it('should apply an operation on each key/value pair', () => {
    let table = new HashTable();
    let key1 = {
      id: 1,
      value: "Hello",
    };
    let key2 = {
      id: 2,
      value: "Wesh"
    };
    
    let spy = sinon.spy();
    
    table.put(key1, "value1");
    table.put(key2, "value2");
    
    
    table.forEach((key, value) => {
      spy(key, value);
    });
    
    
    expect(spy.getCall(0).args[0] == key1.id).to.be.true;
    expect(spy.getCall(0).args[1]).to.equal("value1");
    
    expect(spy.getCall(1).args[0] == key2.id).to.be.true;
    expect(spy.getCall(1).args[1]).to.equal("value2");
  });
});