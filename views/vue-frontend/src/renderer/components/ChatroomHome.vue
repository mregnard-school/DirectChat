<template>
  <div class="chatroom-home">

    <div class="chatroom-thumbnails">
      <div class="chatrooms-search">
        <input type="text" placeholder="Search...">
      </div>
        <div  v-for="chatroom in conversations">
          <chatroom-thumbnail class="chatroom-thumbnail"
                              :chatroom="chatroom"
                              @select-chatroom="changeChatroom"
          />
        </div>
        <div>
          <button class="fab" @click="handleNewConversation">+</button>
        </div>
    </div>

    <div class="selected-chatroom">
      <keep-alive>
        <component :is="activeComponent"
                   v-bind="activeProperties"
                   :key="activeChatroom.id"
                   @new-chatroom="newChatroom"
                   @last-message="setLastMessage"
                   @new-message="save"
        />
      </keep-alive>
    </div>

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
      setLastMessage(chatroom, message) {
        this.conversations.filter(chtroom => chtroom.id === chatroom.id)
            .forEach(chat => chat.last_message = message);
      },
      removeChatroom(chatroom) {
        for (let i = 0; i < this.chatrooms.length; i++) {
          if(this.chatrooms[i].id === chatroom.id) {
            this.removeChatroomWithIndex(i);
          }
        }
      },
      removeChatroomWithIndex(i) {
        this.chatrooms.splice(i,1);
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
            content: "Chat with {0} created ! ",
            author: {
              id: this.client.id,
              pseudo: this.client.pseudo,
            }
          };

          this.createAndSendChatroomCreationMessage(messageTempalte, friends, this.client.pseudo);
          const name = pseudos.join(", ");
          const messageForMyself = this.createMessage(messageTempalte, name);

          const conversation = this.createConversationWithWrapper(messageForMyself, true);
          this.setLastMessage(conversation, messageForMyself);
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
        createdMessage.conversation.name = chatroomName;
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
          'id': this.chatrooms.length + 1,
          'name': message.conversation.name,
          'friends': message.conversation.friends || [],
          'messages': [],
          'last_message': message,
        };

        return conversation;
      },
      createConversationWithWrapper(message, setCurrent) {
        const conversation = this.createConversation(message);

        let conversationWrapper = {
          conversation: conversation,
          component: Chatroom,
        };

        this.chatrooms.push(conversationWrapper);
        if(setCurrent) {
          this.selectedChatroom = conversationWrapper;
        }
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

        conversation.last_message = message;
        if(this.selectedChatroom && this.selectedChatroom.conversation) {
          if(conversation.id !== this.selectedChatroom.conversation.id) {
            conversation.read = false;
          }
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
      conversations() {
        return this.chatrooms.map(wrapper => wrapper.conversation);
      },
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
    display: flex;
    height: 100%;
    flex-direction: row;
    border-right: 1px solid $dividerColor;

    .chatroom-thumbnails {
      flex: 1;
      display: flex;
      flex-direction: column;
      border-right: 1px solid $dividerColor;
      position: relative;
      overflow: auto;

      .chatrooms-search {
        margin: 4px;
      }
      .chatroom-thumbnail {
      }

      .fab {
        position: absolute;
        bottom: 2%;
        right: 3%;
        border-radius: 100px;
        height: 30px;
        width: 30px;
        font-size: 20px;
      }
    }

    .selected-chatroom {
      flex: 3;
    }
  }
</style>