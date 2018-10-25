<template>
  <div class="chatroom">
    <div class="header">
      <div v-if="changingName">
        <input v-model="conversation.name" type="text" v-on:keyup.enter="toggleChangingName"/>
      </div>
      <div v-else @click="toggleChangingName">
        <h1>{{conversation.name}}</h1>
      </div>
    </div>

    <div ref="messagesDisplay" class="conversation">
      <div v-for="message in conversation.messages">
        <message v-bind:message="message"></message>
      </div>
    </div>

    <div class="sendBox">
      <label for="messageInput"></label>
      <input type="text" id="messageInput" v-model="messageToSend" v-on:keyup.enter="sendMessage">
      <button @click="sendMessage">Send</button>
    </div>
  </div>
</template>
<script>

  import Client from 'p2p/client/client';
  import Message from 'components/Message';
  import store from '@/mutableStore';
  import types from '@/messageTypes';
  import moment from 'moment';

  export default {
    name: 'Chatroom',
    components: {
      Message,
    },
    props: {
      conversation: {
        type: Object,
      }
    },
    data() {
      return {
        messageToSend: '',
        changingName: false,
      }
    },
    mounted() {
      this.name = this.name === '' ? this.conversation.friend.pseudo : this.name;

    },
    methods: {
      toggleChangingName() {
        if (this.conversation.name != '') {
          this.changingName = !this.changingName;
          if (!this.changingName) { //Check if new name is different than old one
            let message = this.constructMessage(
                store.state.peer.client.pseudo + ' renamed the conversation to '
                + this.conversation.name
            );
            message = {
              ...message,
              type: types.nameChange,
            };
            this.writeMessageToAll(message);
          }
        }
      },
      onReceiveData(data) {
        this.addNewMessage(data);
      },
      sendMessage() {
        if (this.messageToSend !== '') {
          this.constructAndWriteMessageToAll(this.messageToSend);
          this.messageToSend = '';
        }
      },
      constructMessage(content) {
        return {
          id: this.conversation.messages.length + 1, //Maybe change with unique id (SHA-256 of content + date)
          conversation: {
            id: this.conversation.id,
            name: this.conversation.name,
            friends: this.conversation.friends,
          },
          date: moment().format("YYYY-mm-DD HH:mm:ss"),
          author: {
            id: store.state.peer.client.id,
            pseudo: store.state.peer.client.pseudo,
          },
          content: content,
        };
      },
      constructAndWriteMessageToAll(content) {
        const message = this.constructMessage(content);
        this.writeMessageToAll(message);
      },
      writeMessageToAll(message) {
        this.conversation.friends.forEach(friend => {
          if (friend.id !== store.state.peer.client.id) {
            this.writeMessageTo(friend, message);
          }
        });
        this.conversation.messages.push(message);
        this.$emit('new-message');
      },
      writeMessageTo(friend, message) { // TODO irindul 2018-10-25 : Handle not connected
        store.state.peer.node.writeMessageTo(friend, message);
      },
      addNewMessage(message) {
        this.conversation.messages.push(JSON.parse(message));
        let messagesDisplay = this.$refs.messagesDisplay;
        messagesDisplay.scrollTop = messagesDisplay.scrollHeight;
      }
    }
  }
</script>

<style lang="scss">
  @import '~styles/global';

  .chatroom {
    position: relative;
    display: flex;
    flex-direction: column;
    max-height: 100%;
    padding: 10px;
    .header {
      flex: 1;
      h1 {
        font-family: $main-font;
        font-size: 10px;
      }
    }

    .conversation {
      display: flex;
      flex: 8;
      overflow: auto;
      flex-direction: column;
      box-shadow: 0 0 5px $lightGrey;
      border-radius: 5px;
      padding: 10px;
      margin-bottom: 10px;
      min-height: 80%;
      .message {
        flex: 1;
      }
    }

    .sendBox {
      flex: 1;
    }
  }
</style>
