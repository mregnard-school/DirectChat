<template>
  <div class="home">
    <div class="information">
      <h1>Hello {{this.client.pseudo}} !</h1>

      <div class="loggout" @click.stop="disconnect">
        <router-link to="/" class="link">
          Log out
        </router-link>
      </div>
    </div>

    <div class="dashboard">
      <chatroom-home class="chatrooms"/>
      <friend-list />
    </div>

  </div>
</template>

<script>

  import ChatroomHome from 'components/ChatroomHome';
  import FriendList from 'components/FriendList';
  import store from '@/mutableStore';

  export default {
    name: "Home",
    components: {
      ChatroomHome,
      FriendList,
    },
    mounted() {
      this.client.friends.slice().forEach(friend => {
        if(friend.ips.length > 0) {
          friend.isConnected = true;
          this.$store.commit('connectFriend', friend);
        } else {
          friend.isConnected = false;
          this.$store.commit('disconnectFriend', friend);
        }
      });

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
        this.$store.commit('connectFriend', JSON.parse(JSON.stringify(socket.client)));
      },
      onEndConnection(socket) {
        socket.client.ips = [];
        socket.client.isConnected = false;
        this.$store.commit('disconnectFriend', JSON.parse(JSON.stringify(socket.client)));
      },
      disconnect() {
        this.$store.commit('removeFriends');
        this.$store.commit('removeToken');
        store.clean();
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