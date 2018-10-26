<template>
  <div class="home">
    <div class="information">
      <h1>Hello {{this.client.pseudo}} !</h1>
    </div>

    <div class="dashboard">
      <chatroom-home class="chatrooms"/>

      <div v-for="friend in friends" class="friends"> <!--todo refact in own component-->
        Friends lists :
        <div class="friend-info">{{friend.pseudo}}
          <div v-if="friend.isConnected" class="connected">
            Connected ! <!-- todo change this with icon or smthg better-->
          </div>
        </div>
      </div>
    </div>




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
      this.disconnected = this.client.friends.slice();
      for (let i = 0; i < this.disconnected.length; i++) {
        let client = this.disconnected[i];
        if (client.ips.length > 0) {
          this.disconnected.splice(i, 1);
          client.isConnected = true;
          this.connected.push(client);
        }
      }

      this.node.setOnNewConnection(this.onNewConnection);
      this.node.setOnEndConnection(this.onEndConnection);

    },
    watch: {
      dicsonnected() {
        console.log("disconnected changed");
      }
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
          if (client.id === socket.client.id) {
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
          if (client.id === socket.client.id) {
            this.connected.splice(i, 1);
          }
        }

        this.disconnected.push(socket.client);
      }
    }
  }

</script>

<style lang="scss">
  @import '~styles/global';

  .home {
    display: flex;
    flex-direction: column;
    background: $primaryColor;
    height: 100%;

    .information {
      text-align: center;
      flex: 1;
      h1 {
        font-size: 16px;
        font-weight: 500;
      }
    }

    .dashboard {
      flex: 99;
      display: flex;
      flex-direction: row;
      border-top: 1px solid $dividerColor;

      .chatrooms {
        flex: 4;
      }
      .friends {
        flex: 1;
      }
    }

  }
</style>