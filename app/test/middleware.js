const sinon = require('sinon');
const expect = require('chai').expect;
const NewLiner = require('../nodes/middlewares/newLineMiddleware');

describe("Middleware test", () => {
  it('should chain operations on a message', () => {
    const newliner1 = new NewLiner();
    const newliner2 = new NewLiner();
    newliner1.setNext(newliner2);
    
    let message = "Hello";
    let expected = message + "\n" + "\n";
    message = newliner1.applyMiddleware(message);
    
    expect(message).to.equal(expected);
  })
});