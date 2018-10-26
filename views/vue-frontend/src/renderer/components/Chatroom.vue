<template>
  <div class="chatroom">
    <div class="header">
      <div v-if="changingName"> <!--todo Exit without change with escape -->
        <input v-model="conversation.name" type="text" v-on:keyup.enter="toggleChangingName"/>
      </div>
      <div v-else @click="toggleChangingName">
        <h1>{{conversation.name}}</h1>
      </div>
    </div>

    <div class="chatroom-container">
      <div ref="messagesDisplay" class="conversation">
        <div v-for="message in conversation.messages">
          <message
                   v-bind:message="message" />
        </div>

      </div>
    </div>


    <div class="sendBox">
      <div class="sendBox-input">
        <label for="messageInput"></label>
        <input type="text" id="messageInput" v-model="messageToSend" v-on:keyup.enter="sendMessage">
      </div>
      <div>
        <button @click="sendMessage">Send</button>
      </div>

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
    display: flex;
    flex-direction: column;
    padding: 5px;
    height: 90%;
    .header {
      flex: 1;
      text-align: center;
      h1 {
        font-size: 17px;
        font-weight: 600;
      }
    }

    .chatroom-container {
      flex: 8;
      background: $primaryLightColor;
      padding: 10px;
      border-radius: 5px;
      margin-bottom: 10px;
      max-height: 80%;

      .conversation {
        overflow: auto;
        max-height: 100%;
      }
    }

    .sendBox {
      flex: 2;
      display: flex;
      flex-direction: row;

      .sendBox-input {
        margin-right: 5px;
      }

    }
  }
</style>
