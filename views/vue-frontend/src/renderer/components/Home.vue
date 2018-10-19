<template>
  <div class="home">
    <!-- // TODO irindul 2018-10-18 : temp-->
    <div v-if="run">
      <label for="port">Port : </label>
      <input type="text" v-model="port" id="port">
      <button @click="runServer">Run server</button>
    </div>

    <div v-for="chatroom in chatrooms">

      <chatroom v-bind:peer="peer" v-bind:friend="chatroom.client"></chatroom>
    </div>
  </div>
</template>

<script>

  import Chatroom from './Chatroom';
  import mocks from 'p2p/services/mock'; // TODO irindul 2018-10-18 : Fetch after auth (create your own object instead of mock)
  import Client from 'p2p/client/client';
  import {HashTable} from 'p2p/services/util';

  export default {
    name: "Home",
    components: {
      Chatroom,
    },
    data() {
      return {
        run: true,
        peer: {},
        chatrooms: [],
        port: 5000,
      }
    },
    mounted() {
      // TODO irindul 2018-10-19 : Fetch conversations from local storage
      this.client = mocks.clients[0];
      this.peer = new Client(this.client);
      this.peer.setOnNewConnection(this.onNewConnection)
    },
    methods: {
      runServer() {
        this.peer.runServer(this.port);
        this.run = false;
      },
      onNewConnection(socket) {
        // TODO irindul 2018-10-19 : Define proper conversation structure
        const conversation = {
          'id': 1, // TODO irindul 2018-10-19 : Define id
          'client': socket.client,
          'socket': socket, // TODO irindul 2018-10-19 : See if useful
          'messages': [], // TODO irindul 2018-10-19 : Fetch from local
        };

        this.chatrooms.push(conversation)
      }
    }
  }
</script>

<style scoped>

</style>