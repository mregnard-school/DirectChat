<template>
  <div class="login">
    <label for="pseudo">Pseudo : </label>
    <input v-model="pseudo" type="text" id="pseudo">

    <label for="password">Password : </label>
    <input v-model="password" type="password" id="password"/>

    <button @click="login">Log in</button>

    <!-- // TODO irindul 2018-10-20 : Add register -->
  </div>
</template>

<script>
  const http = require('p2p/services/axios-wrapper').http;
  import store from '@/mutableStore';
  import Client from 'p2p/client/client';
  export default {
    name: "Login",
    data() {
      return {
        pseudo: 'Billy',
        password: 'azerty',
      }
    },
    methods: {
      login() {
        const payload = {
          pseudo: this.pseudo,
          password: this.password,
        };
        http.post('/login', payload)
            .then((response) => {
              let client = response.data;
              let peer = new Client(client);
              store.push({
                peer: peer,
              });
              this.$router.push('/home');
            })
      },
    }
  }
</script>

<style lang="scss">

</style>