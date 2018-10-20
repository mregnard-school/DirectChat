<template>
  <div class="home">
    <!-- // TODO irindul 2018-10-18 : temp-->
    <div v-if="run">
      <label for="port">Port : </label>
      <input type="text" v-model="port" id="port">
      <button @click="runServer">Run server</button>
    </div>


    <div class="chatroom-thumbnails">
      <div class="chatroom-thumbnail" v-for="chatroom in chatrooms">
        <div @click="selectChatroom(chatroom)">
          {{chatroom.name}}
        </div>
      </div>
    </div>


    <div v-if="selectedChatroom" class="selected-chatroom">
      <chatroom v-bind:peer="peer" v-bind:friend="selectedChatroom.client"></chatroom>
    </div>
  </div>
</template>

<script>

  import Chatroom from 'components/Chatroom';
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
        selectedChatroom: false,
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
          'name': socket.client.pseudo,
          'socket': socket, // TODO irindul 2018-10-19 : See if useful
          'messages': [], // TODO irindul 2018-10-19 : Fetch from local
        };

        this.chatrooms.push(conversation)
      },
      selectChatroom(chatroom) {
        this.selectedChatroom = chatroom;
      }
    }
  }
</script>

<style lang="scss">
  @import '~styles/global';
  .home {
    position: absolute;
    bottom: 0;
    top: 0;
    left: 0;
    right: 0;
    padding: 2px;
    display: flex;
    flex-wrap: wrap;
    .chatroom-thumbnails {
      display: flex;
      flex-direction: column;
      flex: 1;
      border-radius: 2px;
      padding-top: 5px;
      padding-left: 5px;
      padding-bottom: 5px;
      border-right: 1px solid $lightGrey;
      max-width: 100px;
      .chatroom-thumbnail {
        flex: 1;
        border: 4px;
        max-height: 20px;
        border-bottom: 1px solid $lightGrey;
      }
    }

    .selected-chatroom {
      flex: 5;
    }

  }
</style>