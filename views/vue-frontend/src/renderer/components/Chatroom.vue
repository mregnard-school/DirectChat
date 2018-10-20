<template>
  <div class="chatroom">
    chat with {{name}}


    <div class="conversation">
      <div v-for="message in messages">
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

  export default {
    name: 'Chatroom',
    components: {
      Message,
    },
    props: {
      friend: {
        type: Object,
      },
      peer: {
        type: Client,
      }
    },
    data() {
      return {
        messages: [],
        messageToSend: '',
        name: '',
      }
    },
    mounted() {
      this.name = this.friend.pseudo;
      this.peer.node.setOnReceiveData(this.onReceiveData)
    },
    methods: {
      onReceiveData(data) {
        this.messages.push(JSON.parse(data));
      },
      sendMessage() {
        this.peer.node.writeMessageTo(this.friend, this.messageToSend)
            .then(message => {
              this.messages.push(JSON.parse(message));
            });

        this.messageToSend = '';
      }
    }
  }
</script>

<style>
</style>
