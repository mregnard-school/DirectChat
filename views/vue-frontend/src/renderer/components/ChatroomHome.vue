<template>
  <div class="chatroom-home">
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
                 v-on:new-message="save"
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
  import types from '@/messageTypes';

  const localStore = new Store();

  export default {
    name: "ChatroomHome",
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
     // this.node.setOnNewConnection(this.onNewConnection);
      this.chatrooms = localStore.get(this.storeFile) || [];
      this.chatrooms.forEach(wrapper => {
        wrapper.component = Chatroom;
      })
    },
    methods: {
      changeChatroom(chatroom) {
        let wrapper = this.chatrooms.find(wrapper => wrapper.conversation.id === chatroom.id);
        this.selectedChatroom = wrapper;
      },
      handleNewConversation() {
        this.selectedChatroom = {
          conversation: {},
          component: NewChatroom,
        }
      },
      onNewConnection(socket) {
        this.peer.handleFriendConnection(socket.client);
        //Alert user here with something visual
      },
      save() {
        localStore.set(this.storeFile, this.chatrooms);
      },
      newChatroom(pseudos) {
        let friends = this.peer.getFriendsWithPseudos(pseudos);
        if (friends.length > 0) {
          let friendsAndMe = friends.concat([this.client]);
          const messageTempalte = {
            id: 0,
            conversation: {id: 0, name: "", friends: friendsAndMe,},
            type: types.information,
            content: "Chat with {0} created ! "
          };

          this.createAndSendChatroomCreationMessage(messageTempalte, friends, this.client.pseudo);
          const name = pseudos.join(", ");
          const messageForMyself = this.createMessage(messageTempalte, name);

          const conversation = this.createConversationWithWrapper(messageForMyself);
          this.addAndSaveChatroom(conversation, messageForMyself);
        }
      },
      createAndSendChatroomCreationMessage(messageTemplate, friends, name) {
        const createdMessageToSend = this.createMessage(messageTemplate, name);
        this.sendToAll(friends, createdMessageToSend);
      },
      createMessage(messageTemplate, chatroomName) {
        const createdMessage = {
          ...messageTemplate,
        };
        createdMessage.conversation.name = chatroomName
        createdMessage.content = messageTemplate.content.format(chatroomName);
        return createdMessage
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
        let conversation = this.chatrooms.map((wrapper) => wrapper.conversation)
            .find((conv) => conv.id === conversation_id);
        if (!conversation) {
          conversation = this.createConversationWithWrapper(message);
        }

        if(message.type === types.nameChange) {
          conversation.name = message.conversation.name;
        }

        this.addAndSaveChatroom(conversation, message);
      },
      addAndSaveChatroom(chatroom, message) {
        chatroom.messages.push(message);
        this.save();
      },
    },
    computed: {
      storeFile() {
        return this.client.pseudo+'-chatroom';
      },
      client() {
        return this.peer.client
      },
      node() {
        return this.peer.node;
      },
      peer() {
        return store.state.peer;
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

  .chatroom-home {
   // position: absolute;
    //bottom: 0;
    //top: 0;
    //left: 0;
    //right: 0;
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