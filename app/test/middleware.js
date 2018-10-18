const sinon = require('sinon');
const expect = require('chai').expect;
const NewLiner = require('../nodes/middlewares/newLineMiddleware');
const MiddlewareChain = require('../nodes/middlewares/middlewarechain');


describe("Middleware test", () => {
  it('should chain operations on a message', () => {
    const newliner1 = new NewLiner();
    const newliner2 = new NewLiner();
    newliner1.setNext(newliner2);
    
    let message = "Hello";
    let expected = message + "\n" + "\n";
    message = newliner1.applyMiddleware(message);
    
    expect(message).to.equal(expected);
  });
  
  it('should chain operations on a message using a chain', () => {
    const newliner1 = new NewLiner();
    const newliner2 = new NewLiner();
    
    const chain = new MiddlewareChain(newliner1);
    chain.addMiddleware(newliner2);
    
    let message = "Hello";
    let expected = message + "\n" + "\n";
    message = chain.applyMiddlewares(message);
    
    expect(message).to.equal(expected);
  })
});