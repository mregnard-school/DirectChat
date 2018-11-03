<template>
  <div class="login">
    <div class="login-form" @keypress.enter="login">
      <div class="login-input">
        <label for="pseudo">Pseudo</label>
        <input v-model.trim="pseudo" autofocus type="text" id="pseudo" placeholder="Enter your pseudo...">
      </div>

      <div class="login-input">
        <label for="password">Password</label>
        <input v-model.trim="password" type="password" id="password" placeholder="Enter password"/>
      </div>

      <div class="login-submit">
        <button @click="login">
          Log in
        </button>
      </div>

    </div>
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
        if (this.pseudo !== '' && this.password !== '') {
          const payload = {
            pseudo: this.pseudo,
            password: this.password,
          };
          http.post('/clients/login', payload)
              .then((response) => {
                let client = response.data;
                let peer = new Client(client);
                // TODO irindul 2018-11-03 : store token somewhere
                store.push({
                  peer: peer,
                });
                this.peerCreated(peer);
                this.$router.push('/home');
              })
        }
      },
      peerCreated(peer) {
        const port = parseIpAndPortFromString(peer.client.ips[0]).port;
        peer.runServer(port);

        peer.client.friends.forEach(friend => {
          if (friend.ips.length !== 0) {
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
  @import '~styles/global';

  .login {
    display: flex;
    position: absolute;
    justify-content: center;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    background: $primaryColor;
    .login-form {
      display: flex;
      flex-direction: column;
      margin: auto;
      padding: 30px;
      border-radius: 5px;
      background: $primaryLightColor;
      .login-input {
        display: flex;
        flex-direction: column;
        color: $primaryText;
        margin-bottom: 15px;

        input {
          font-size: 16px;
          max-width: 300px;
        }

        label {
          font-size: 15px;
        }
      }

      .login-submit {
        min-width: 55.2px;
        max-width: 100px;
        align-self: flex-end;
      }
    }
  }
</style>