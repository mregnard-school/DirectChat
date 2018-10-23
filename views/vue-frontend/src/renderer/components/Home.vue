<template>
  <div class="home">
    connected as {{this.client.pseudo}}
    <div class="chatroom-thumbnails">
      <div class="chatroom-thumbnail" v-for="wrapper in chatrooms">
        <chatroom-thumbnail v-bind:chatroom="wrapper.conversation"
                            v-on:select-chatroom="changeChatroom"/>
      </div>
      <div>
        <button class="new-conversation" @click="handleNewConversation">New Conversation</button>
      </div>
    </div>


    <keep-alive>
      <component :is="activeComponent"
                 v-bind="activeProperties"
                 :key="activeChatroom.id"
                 v-on:new-chatroom="newChatroom"
                 v-on:new-message="newMessage"
                 class="selected-chatroom"
      />
    </keep-alive>
  </div>
</template>

<script>

  import Chatroom from 'components/Chatroom';
  import NewChatroom from 'components/NewChatroom';
  import ChatroomThumbnail from 'components/ChatroomThumbnail';
  import store from '@/mutableStore';
  import Store from 'electron-store';
  import Client from 'p2p/client/client';
  import {HashTable} from 'p2p/services/util';

  const localStore = new Store();

  export default {
    name: "Home",
    components: {
      Chatroom,
      ChatroomThumbnail,
      NewChatroom,
    },
    data() {
      return {
        chatrooms: [],
        selectedChatroom: null,
      }
    },
    created() {
      this.node.setOnReceiveData(this.onReceiveData);
      this.chatrooms = localStore.get("user-chatrooms") || [];

      this.chatrooms.forEach(wrapper => {
        wrapper.component = Chatroom;
        console.log(wrapper.conversation);
      })
    },
    methods: {
      handleNewConversation() {
        this.selectedChatroom = {
          conversation: {},
          component: NewChatroom,
        }
      },
      onNewConnection(socket) {
        //todo Handle Machin is connected here

      },
      newMessage() {
        localStore.set('user-chatrooms', this.chatrooms);
      },
      newChatroom(pseudos) {
        let friends = [];
        this.client.friends.forEach(friend => {
          if (pseudos.includes(friend.pseudo)) {
            friends.push(friend);
          }
        });
        if (friends.length > 0) {
          let friendsAndMe = friends.concat([this.client]);
          const createdMessageToSend = {
            id: 0,
            conversation: {
              id: 0,
              name: this.client.pseudo,
              friends: friendsAndMe,
            },
            type: "informational",
            content: "Chat with " + this.client.pseudo + " created !"
          };
          this.sendToAll(friends, createdMessageToSend);
          const name = pseudos.join(", ");
          const messageForMyslef = {
            id: 0,
            type: "informational",
            conversation: {
              name: name,
              friends: friendsAndMe,
            },
            content: "Chat with " + name + " created",
          };

          const conversation = this.createConversationWithWrapper(messageForMyslef);
          this.addAndSaveChatroom(conversation, messageForMyslef);
        }
      },
      sendToAll(friends, message) {
        friends.forEach(friend => {
          if (friend.id !== this.client.id) {
            this.node.writeMessageTo(friend, message);
          }
        });
      },
      createConversation(message) {
        const conversation = {
          'id': this.chatrooms.length + 1,// TODO irindul 2018-10-19 : Define id (maybe SHA-256 of all pseudos concatenated),
          'name': message.conversation.name,
          'friends': message.conversation.friends || [],
          'messages': [],
        };

        return conversation;
      },
      createConversationWithWrapper(message) {
        const conversation = this.createConversation(message);

        let conversationWrapper = {
          conversation: conversation,
          component: Chatroom,
        };

        this.chatrooms.push(conversationWrapper);
        return conversation;
      },
      onReceiveData(data) {
        const message = JSON.parse(data);
        let conversation_id = message.conversation.id;
        let conversation = this.chatrooms.map((wrapper) => wrapper.conversation).find(
            (conv) => conv.id === conversation_id);

        if (!conversation) {
          conversation = this.createConversationWithWrapper(message);
        }
        this.addAndSaveChatroom(conversation, message);
      },
      addAndSaveChatroom(chatroom, message) {
        chatroom.messages.push(message);
        localStore.set('user-chatrooms', this.chatrooms);
      },
      changeChatroom(chatroom) {
        let wrapper = this.chatrooms.find(wrapper => wrapper.conversation.id === chatroom.id);
        this.selectedChatroom = wrapper;
      }
    },
    computed: {
      client() {
        return store.state.peer.client;
      },
      node() {
        return store.state.peer.node;
      },
      activeComponent() {
        if (this.selectedChatroom) {
          return this.selectedChatroom.component;
        }

        return null
      },
      activeProperties() {
        if (this.selectedChatroom) {
          return {
            conversation: this.selectedChatroom.conversation
          }
        }
        return {};

      },
      activeChatroom() {
        if (this.selectedChatroom) {
          return this.selectedChatroom.conversation
        }

        return {
          id: 0,
        }
      },
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