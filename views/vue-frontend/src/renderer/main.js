import Vue from 'vue'
import axios from 'axios'

import App from './App'
import router from './router'
import store from './store'

if (!process.env.IS_WEB) Vue.use(require('vue-electron'))
Vue.http = Vue.prototype.$http = axios
Vue.config.productionTip = false

/* eslint-disable no-new */
new Vue({
  components: { App },
  router,
  store,
  template: '<App/>'
}).$mount('#app')

class Test  {
  constructor() {
    this.receiveData = () => {
      console.log('Wesh')
    };
  
    this.wopse = new Woopse();
    this.wopse.receiveData = this.receiveData;
  }
  
  runAsync() {
    let self = this;
    return new Promise((resolve, reject) => {
      this.wopse.runMotherfucker();
    })
  }
}

class Woopse {
  constructor() {
    this.receiveData = () => {console.log("Wopsed")};
  }
  
  runMotherfucker = () => {
    this.receiveData();
  }
}
const test = new Test();
test.runAsync();
test.receiveData = () => {
  console.log('oupsi doopsi');
};
test.runAsync();