class Middleware {
  constructor(next) {
    this.next = next;
  }
  
  applyMiddleware(message) {
    let messageTreated = this.modifyMessage(message);
    if(this.next) {
      messageTreated = this.next.applyMiddleware(messageTreated);
    }
    
    return messageTreated;
  }
  
  setNext(middleware) {
    this.next = middleware;
  }
  
  modifyMessage(message) {
    return message;
  }
}

module.exports = Middleware;