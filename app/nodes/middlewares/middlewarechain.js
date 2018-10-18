class MiddlewareChain {
  constructor(first) {
    this.first = first;
    this.last = first;
  }
  
  addMiddleware(middleware) {
    this.last.setNext(middleware);
    
    if(middleware) {
      this.last = middleware;
    }
  }
  
  applyMiddlewares(message) {
    return this.first.applyMiddleware(message);
  }
  
}

module.exports = MiddlewareChain;