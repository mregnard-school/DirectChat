<template>
  <div class="chatroom">
    chat with {{name}}
    <!-- // TODO irindul 2018-10-18 : temp-->
    <div v-if="run">
      <label for="port">Port : </label>
      <input type="text" v-model="port" id="port">
      <button @click="runServer">Run server</button>
    </div>

    <div class="conversation">
      <div v-for="message in messages">
        <p>
          <!-- // TODO irindul 2018-10-19 : Template message with JSON object and VUe component-->
          {{message}}
        </p>
      </div>
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
        name: 'Billy',
        client: {},
        port: 5000,
        run: true,
        messages: [],
      }
    },
    mounted() {
      this.peer = new Client(mocks.clients[0]);
      this.peer.setOnReceiveData(this.onReceiveData);
      this.name = this.peer.client.pseudo;
    },
    methods: {
      runServer() {
        this.peer.runServer(this.port);
        this.run = false;
      },
      onReceiveData(data) {
        this.messages.push(data);
      }
    }
  }
</script>

<style>
</style>
