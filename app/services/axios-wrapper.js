const axios = require('axios');
const mock = require('./mock');
class Wrapper {
  constructor() {
    this.initService();
  }
  
  
  
  initService() {
    this.serviceAxios = axios.create({});
    this.serviceAxios.interceptors.request.use(
        config => {
          const token = ''; // TODO irindul 2018-10-15 : Add JWT token
          if (token) {
            config.headers.Authorization = 'bearer ' + token;
          }
          
          return config;
        },
        error => {
          return Promise.reject(error);
        }
    );
    this.serviceAxios.interceptors.response.use(null, this.handleError);
  }
  
  handleError(error) {
    if (error.response.status === 401) {
      //Handle unauthorized
    }
    return Promise.reject(error);
  }
}

class MockWrapper {
  constructor() {
  
  }
  
  get(url) {
    if(url.includes('ipList')) {
      return Promise.resolve({
        data: mock.ipList,
      })
    }
  }
}

//const http = new Wrapper().service;
const http = new MockWrapper();
module.exports = {http};