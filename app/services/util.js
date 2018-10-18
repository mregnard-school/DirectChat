function HashTable() {
  this.map = {};
  this.count = 0;
}

HashTable.prototype = {
  constructor: HashTable,
  put: function(key, value) {
    if(!key.hasOwnProperty('id')) {
      throw 'Object ' + key + 'must have an id';
    }
    
    this.map[key.id] = value;
    this.count++;
  },
  get: function(key) {
    return this.map[key.id];
  },
  
  forEach: function(callback) {
    for(const property in this.map) {
      if(this.map.hasOwnProperty(property)) {
        callback(property, this.map[property]);
      }
    }
  }
};

module.exports = {HashTable};