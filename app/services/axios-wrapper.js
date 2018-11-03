// TODO irindul 2018-11-03 : Move into frontend folder
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
          config.baseURL = 'http://localhost:8000/api'; // TODO irindul 2018-11-03 : Retrieve from env.js
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
  
  post(url, payload) {
    if(url.includes('login')) {
      if(payload.pseudo === 'Billy' && payload.password === 'azerty') {
        return Promise.resolve({
          data: {
            id: 1, // TODO irindul 2018-10-20 : Create token here
            pseudo: payload.pseudo,
            ips: [
                "127.0.0.1:5000",
            ],
            friends: [
              {
                "id": 2,
                "pseudo": "John",
                ips: [],
              },
              {
                "id": 3,
                "pseudo": "Henry",
                ips: [],
              }
            ]
          }
        })
      }
      if(payload.pseudo === 'John' && payload.password === 'azerty') {
        return Promise.resolve({
          data: {
            id: 2,
            pseudo: payload.pseudo,
            ips: [
              "127.0.0.1:5001",
            ],
            friends: [
              {
                "id": 1,
                "pseudo": "Billy",
                ips: [
                    "127.0.0.1:5000"
                ],
              },
              {
                "id": 3,
                "pseudo": "Henry",
                ips: [
                
                ],
              }
            ]
          }
        })
      }
      
      if(payload.pseudo === 'Henry' && payload.password === 'azerty') {
        return Promise.resolve({
          data: {
            id: 3,
            pseudo: payload.pseudo,
            ips: [
              "127.0.0.1:5002",
            ],
            friends: [
              {
                "id": 1,
                "pseudo": "Billy",
                ips: [
                  "127.0.0.1:5000"
                ],
              },
              {
                "id": 2,
                "pseudo": "John",
                ips: [
                    "127.0.0.1:5001"
                ],
              }
            ]
          }
        })
      }
    }
  }
}

//const http = new Wrapper().service;
const http = new Wrapper().serviceAxios;
module.exports = {http};