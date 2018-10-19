<template>
  <div class="chatroom">
    chat with {{name}}


    <div class="conversation">
      <div v-for="message in messages">
        <p>
          <!-- // TODO irindul 2018-10-19 : Template message with JSON object and VUe component-->
          {{message}}
        </p>
      </div>
    </div>

    <div class="sendBox">
      <label for="messageInput"></label>
      <input type="text" id="messageInput" v-model="messageToSend">
      <button @click="sendMessage">Send</button>
    </div>
  </div>
</template>

<script>

  import Client from 'p2p/client/client';

  export default {
    name: 'Chatroom',
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
        this.messages.push(data);
      },
      sendMessage() {
        this.peer.node.writeMessageTo(this.friend, this.messageToSend);
        this.messages.push(this.messageToSend);
        this.messageToSend = '';
      }
    }
  }
</script>

<style>
</style>
