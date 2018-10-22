<template>
  <div class="home">
    <div class="chatroom-thumbnails">
      <div class="chatroom-thumbnail" v-for="chatroom in chatrooms">
        <chatroom-thumbnail v-bind:chatroom="chatroom"
                            v-on:select-chatroom="changeChatroom"/>
      </div>
    </div>

    <div v-if="selectedChatroom" class="selected-chatroom">
      <chatroom v-bind:conversation="selectedChatroom"></chatroom>
    </div>
  </div>
</template>

<script>

  import Chatroom from 'components/Chatroom';
  import ChatroomThumbnail from 'components/ChatroomThumbnail';
  import store from '@/mutableStore';
  import Client from 'p2p/client/client';
  import {HashTable} from 'p2p/services/util';
  import {parseIpAndPortFromString} from 'p2p/services/util';

  export default {
    name: "Home",
    components: {
      Chatroom,
      ChatroomThumbnail,
    },
    data() {
      return {
        chatrooms: [],
        selectedChatroom: false,
      }
    },
    mounted() {
      store.state.peer.setOnNewConnection(this.onNewConnection);
      const port = parseIpAndPortFromString(store.state.peer.client.ips[0]).port;
      store.state.peer.runServer(port);
    },
    methods: {
      onNewConnection(socket) {
        // TODO irindul 2018-10-19 : Define proper conversation structure
        const conversation = {
          'id': this.chatrooms.length+1, // TODO irindul 2018-10-19 : Define id (maybe SHA-256 of all pseudos concatenated)
          'friend': socket.client,
          'name': socket.client.pseudo,
          'socket': socket, // TODO irindul 2018-10-19 : See if useful
          'messages': [], // TODO irindul 2018-10-19 : Fetch from local
        };

        this.chatrooms.push(conversation)
      },
      changeChatroom(chatroom) {
        this.selectedChatroom = chatroom;
        console.log(this.selectedChatroom);
      }
    },
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