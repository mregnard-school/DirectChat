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
        this.changingName = !this.changingName;
        // TODO irindul 2018-10-22 : find a way to propagate change for all members of chat
      },
      onReceiveData(data) {
        this.addNewMessage(data);
      },
      sendMessage() {
        if(this.messageToSend !== '') {
          this.writeMessageToAll(this.messageToSend);
          this.messageToSend = '';
        }


      },
      writeMessageToAll(content) {
        const message = {
          id: this.conversation.messages.length+1, // TODO irindul 2018-10-20 : Genererate id... (sha-256 of content + date(for unicity) is the best)
          conversation_id: this.conversation.id,
          date: moment().format("YYYY-mm-DD HH:mm:ss"),
          author: {
            id: store.state.peer.client.id,
            pseudo: store.state.peer.client.pseudo,
          },
          content: content,
        };
        this.conversation.friends.forEach(friend => {
          if(friend.id !== store.state.peer.client.id) {
            this.writeMessageTo(friend, message);
          }
        });
        this.conversation.messages.push(message);
      },
      writeMessageTo(friend, message) {
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
        font-size: Medium; // TODO irindul 2018-10-20 : Change with custom font-size
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
