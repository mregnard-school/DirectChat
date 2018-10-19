<template>
  <div class="chatroom">

    <!-- // TODO irindul 2018-10-18 : temp-->
    <label for="port">Port : </label>
    <input type="text" v-model="port">
    <button @click="runServer">Run server</button>
    <div class="conversation">

    </div>

    <div class="sendBox">
      <label for="messageInput"></label>
      <input type="text" id="messageInput">

      <button>Send</button>
    </div>
  </div>
</template>

<script>

  import mocks from 'p2p/services/mock'; // TODO irindul 2018-10-18 : Add as props
  import Client from 'p2p/client/client';
  export default {
    name: 'Chatroom',
    props: [
      //'peer', // TODO irindul 2018-10-18 : fetch from server
    ],
    data() {
      return {
        name: '',
        client: {},
        port: 5000,
      }
    },
    mounted() {
      this.peer = new Client(mocks.clients[0]);
      this.name = this.peer.pseudo;

    },
    methods: {
      runServer() {
        this.peer.runServer(this.port);
      }
    }
  }
</script>

<style>
</style>
