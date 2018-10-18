const Middleware = require('./middleware');

class NewLiner extends Middleware {
  
  modifyMessage(message) {
    return message + "\n";
  }
}

module.exports = NewLiner;