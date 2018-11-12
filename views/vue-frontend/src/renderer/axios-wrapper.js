import axios from 'axios';
import store from './store';
import env from './env';
class Wrapper {
  constructor() {
    this.env = env;
    this.initService();
    console.log(env);
  }
  
  initService() {
    this.serviceAxios = axios.create({});
    this.serviceAxios.interceptors.request.use(
        config => {
          config.baseURL = env.url;
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