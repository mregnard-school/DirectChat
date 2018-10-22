<template>
  <div class="home">
    <div class="chatroom-thumbnails">
      <div class="chatroom-thumbnail" v-for="wrapper in chatrooms">
        <chatroom-thumbnail v-bind:chatroom="wrapper.conversation"
                            v-on:select-chatroom="changeChatroom"/>
      </div>
    </div>


    <keep-alive>
      <component :is="activeComponent"
                 v-bind="activeProperties"
                 :key="activeChatroom.id"
      />
    </keep-alive>
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
        selectedChatroom: null,
      }
    },
    mounted() {
      store.state.peer.setOnNewConnection(this.onNewConnection);
      const port = parseIpAndPortFromString(store.state.peer.client.ips[0]).port;
      store.state.peer.runServer(port);
      store.state.peer.node.setOnReceiveData(this.onReceiveData)
    },
    methods: {
      onNewConnection(socket) {
        // TODO irindul 2018-10-22 : Check here if convo contains the client otherwise create new
        // TODO irindul 2018-10-19 : Define proper conversation structure
        const conversation = {
          'id': this.chatrooms.length + 1, // TODO irindul 2018-10-19 : Define id (maybe SHA-256 of all pseudos concatenated)
          'friends': [socket.client],
          'name': socket.client.pseudo,
          'messages': [], // TODO irindul 2018-10-19 : Fetch from local
        };

        let conversationWrapper = {
          conversation: conversation,
          component: Chatroom,
        };
        this.chatrooms.push(conversationWrapper)
      },
      onReceiveData(data){
        const message = JSON.parse(data);
        let conversation_id = message.conversation_id;
        let conversation = this.chatrooms.map((wrapper) => wrapper.conversation).find((conv) => conv.id === conversation_id)
        conversation.messages.push(message);
      },
      changeChatroom(chatroom) {
        let wrapper = this.chatrooms.find(wrapper => wrapper.conversation.id === chatroom.id);
        this.selectedChatroom = wrapper;
        console.log(this.selectedChatroom.component.data);
      }
    },
    computed: {
      activeComponent() {
        if(this.selectedChatroom) {
          return this.selectedChatroom.component;
        }

        return null
      },
      activeProperties() {
        if(this.selectedChatroom) {
          return {
            conversation: this.selectedChatroom.conversation
          }
        }

        return {};

      },
      activeChatroom() {
        if(this.selectedChatroom) {
          return this.selectedChatroom.conversation
        }

        return {
          id: 0,
        }
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