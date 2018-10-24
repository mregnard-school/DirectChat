<template>
  <div class="login">
    <label for="pseudo">Pseudo : </label>
    <input v-model="pseudo" type="text" id="pseudo">

    <label for="password">Password : </label>
    <input v-model="password" type="password" id="password"/>

    <button @click="login">Log in</button>
  </div>
</template>

<script>
  const http = require('p2p/services/axios-wrapper').http;
  import store from '@/mutableStore';
  import Client from 'p2p/client/client';
  import {parseIpAndPortFromString} from 'p2p/services/util';
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
              this.peerCreated(peer);
              this.$router.push('/home');
            })
      },
      peerCreated(peer) {
        //peer.setOnNewConnection(this.onNewConnection);
        const port = parseIpAndPortFromString(peer.client.ips[0]).port;
        peer.runServer(port);

        peer.client.friends.forEach(friend => {
          if(friend.ips.length !== 0) {
            friend.ips.forEach(ipPort => {
              const parsed = parseIpAndPortFromString(ipPort);
              peer.node.connectTo(parsed.ip, parsed.port);
            })
          }
        });
      }
    }
  }
</script>

<style lang="scss">

</style>