<template>
  <div class="home">
    <div class="informations">
      Hello {{this.client.pseudo}} !
    </div>

    <div v-for="friend in friends" class="friends">
      Friends lists :
      <div class="friend-info">{{friend.pseudo}}
        <div v-if="friend.isConnected" class="connected">
          Connected ! <!-- todo change this with icon or smthg better-->
        </div>
      </div>

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
      this.node.setOnEndConnection(this.onEndConnection);

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
      },
      onEndConnection(socket) {
        socket.client.ips = [];
        socket.client.isConnected = false;
        for (let i = 0; i < this.connected.length; i++) {
          let client = this.connected[i];
          if(client.id === socket.client.id) {
            this.connected.splice(i, 1);
          }
        }

        this.disconnected.push(socket.client);
      }
    }
  }

</script>

<style lang="scss">
.home {
  position: absolute;
  top: 0;
  bottom: 0;
  left: 0;
  right: 0;
  padding: 10px;

  .friends {
    .friend-info {

    }
    .active {

    }
  }
}
</style>