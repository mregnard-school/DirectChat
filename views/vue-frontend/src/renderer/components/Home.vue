<template>
  <div class="home">
    <div class="informations">
      Hello {{this.client.pseudo}} !
    </div>

    <div v-for="friend in friends">
      {{friend}}
    </div>

    <chatroom-home/>
  </div>
</template>

<script>

  import ChatroomHome from 'components/ChatroomHome'
  import store from '@/mutableStore';
  export default {
    name: "Home",
    components: {
      ChatroomHome,
    },
    data() {
      return {
        connected: [],
        disconnected: [],
      }
    },
    mounted() {
      this.disconnected = this.client.friends;
      for (let i = 0; i < this.disconnected.length; i++) {
        let client = this.disconnected[i];
        if(client.ips.length > 0) {
          this.disconnected.splice(i, 1);
          client.isConnected = true;
          this.connected.push(client);
        }
      }

      this.node.setOnNewConnection(this.onNewConnection);

    },
    computed: {
      friends() {
        return this.connected.concat(this.disconnected);
      },
      client() {
        return store.state.peer.client;
      },
      node() {
        return store.state.peer.node;
      }
    },
    methods: {
      onNewConnection(socket) {
        socket.client.isConnected = true;
        for (let i = 0; i < this.disconnected.length; i++) {
          let client = this.disconnected[i];
          if(client.id === socket.client.id) {
            this.disconnected.splice(i, 1);
          }
        }
        this.connected.push(socket.client);
      }
    }
  }

</script>

<style scoped>

</style>