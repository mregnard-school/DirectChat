const MiddlewareChain = require('./middlewarechain');
const NewLiner = require('./newLineMiddleware');
const newLiner = new NewLiner();

const chain = new MiddlewareChain(newLiner);

module.exports = chain;