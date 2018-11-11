<template>
  <div class="chatroom">
    <div class="header">
      <div v-if="changingName">
        <input v-model="name" type="text"
               ref="name"
               v-on:keyup.enter="toggleChangingName"
               v-on:keyup.esc="cancel"
        />
      </div>
      <div v-else @click="toggleChangingName">
        <h1>{{conversation.name}}</h1>
      </div>
    </div>

    <div class="chatroom-container" ref="chatroom">
      <div ref="messagesDisplay" class="conversation">
        <div v-for="message in conversation.messages">
          <message v-bind:message="message"/>
        </div>
      </div>
    </div>


    <div class="sendBox">
      <div class="sendBox-input">
        <label for="messageInput"></label>
        <input type="text" id="messageInput" v-model.trim="messageToSend"
               v-on:keyup.enter="sendMessage">
      </div>
      <div>
        <button @click="sendMessage">Send</button>
      </div>

    </div>
  </div>
</template>
<script>

  const http = require('@/axios-wrapper').http;
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
        name: '',
        messageToSend: '',
        changingName: false,
      }
    },
    mounted() {
      this.name = this.conversation.name;
      this.scrollDown();
    },
    computed: {
      nameHasChanged() {
        return this.name !== this.conversation.name;
      }
    },
    methods: {
      cancel() {
        if (this.changingName) {
          this.changingName = false;
        }
      },
      toggleChangingName() {
        if (this.name !== '') {
          this.changingName = !this.changingName;
          if (this.changingName) {
            //Focus on input component so we can start to write right away
            //Doesn't work without nextTick()
            this.$nextTick(() => this.$refs.name.focus());
          }
          if (!this.changingName && this.nameHasChanged) {
            let message = this.constructMessage(
                store.state.peer.client.pseudo + ' renamed the conversation to ' + this.name,
            );
            message = {
              ...message,
              type: types.nameChange,
            };
            message.conversation.name = this.name;
            this.conversation.name = this.name;
            this.writeMessageToAll(message);
          }
        }
      },
      sendMessage() {
        if (this.messageToSend !== '') {
          this.constructAndWriteMessageToAll(this.messageToSend);
          this.messageToSend = '';
          this.scrollDown();
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
        this.$emit('last-message', this.conversation, message);
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
      writeMessageTo(friend, message) {
        this.$store.dispatch('isConnected', friend).then(isConnected => {
          if (!isConnected) {
            this.handleNotConnected(friend, message);
          }
        });
        store.state.peer.node.writeMessageTo(friend, message);
      },
      handleNotConnected(friend, message) {
        //todo test
        http.post(`/client/${friend.id}/messages`, message)
            .then(() => {
              //Yeah message was stored on the server
            }).catch(() => {

        })
      },
      scrollDown() {
        this.$nextTick(() => {
          this.$refs.messagesDisplay.scrollTop =
              this.$refs.messagesDisplay.scrollHeight;
        });

      }
    },
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
