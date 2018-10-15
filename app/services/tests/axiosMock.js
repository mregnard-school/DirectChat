const wrapper = require('../axios-wrapper');

const testAxiosMockGet = function() {
  wrapper.http.get('ipList')
      .then((response) => {
        console.log(response.data);
      });
};


module.exports = {
  testAxiosMockGet,
};
