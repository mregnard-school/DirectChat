import axios from 'axios';
import store from './store';

class Wrapper {
  constructor() {
    this.initService();
  }
  
  initService() {
    this.serviceAxios = axios.create({});
    this.serviceAxios.interceptors.request.use(
        config => {
          config.baseURL = 'http://localhost:8000/api';// TODO irindul 2018-11-03 : Retrieve from env.js
          return config;
        },
        error => {
          return Promise.reject(error);
        }
    );
    this.serviceAxios.interceptors.response.use(null, this.handleError);
  }
  
  setToken() {
    axios.defaults.headers.common['Authorization'] = `Bearer ${store.state.Auth.token}`;
    /*this.serviceAxios.interceptors.request.use(
        config => {
          config.baseURL = 'http://localhost:8000/api'; // TODO irindul 2018-11-03 : Retrieve from env.js
          const token = store.state.Auth.token;
          if (token) {
            config.headers.Authorization = 'Bearer ' + token;
          }
          config.headers.Samere = 'Swag';
        
          return config;
        },
        error => {
          return Promise.reject(error);
        }
    );*/
  }
  
  handleError(error) {
    if (error.response.status === 401) {
      //Handle unauthorized
      return Promise.reject();
    }
    return Promise.reject(error);
  }
}

//const http = new Wrapper().service;
const wrapper = new Wrapper();
const http = wrapper.serviceAxios;
export {http, wrapper};