<template>
  <div class="home">
    <div class="information-home">

      <div class="user">
        <h1>
          <span class="welcome">
            Welcome
          </span>
          <span class="pseudo">
          {{this.client.pseudo}}
          </span>
        </h1>
      </div>

      <div class="logout" @click.stop="disconnect">
        <router-link to="/" class="link">
          Log out
        </router-link>
      </div>
    </div>

    <div class="dashboard">
      <chatroom-home class="chatrooms"/>
      <friend-list @new-connected="onNewConnection"
                   @new-disconnected="onEndConnection"
      />
    </div>

  </div>
</template>

<script>

  import ChatroomHome from 'components/ChatroomHome';
  import FriendList from 'components/FriendList';
  import store from '@/mutableStore';
  import {http} from '@/axios-wrapper';

  export default {
    name: "Home",
    components: {
      ChatroomHome,
      FriendList,
    },
    mounted() {
      this.client.friends.slice().forEach(friend => {
        if (friend.ips.length > 0) {
          friend.isConnected = true;
          this.$store.commit('connectFriend', friend);
        } else {
          friend.isConnected = false;
          this.$store.commit('disconnectFriend', friend);
        }
      });

      this.node.setOnNewConnection((socket) => {
        this.onNewConnection(socket.client);
      });
      this.node.setOnEndConnection((socket) => {
        this.onEndConnection(socket.client);
      });

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
      onNewConnection(client) {
        client.isConnected = true;
        this.$store.commit('connectFriend', JSON.parse(JSON.stringify(client)));
      },
      onEndConnection(client) {
        client.ips = [];
        client.isConnected = false;
        this.$store.commit('disconnectFriend', JSON.parse(JSON.stringify(client)));
      },
      disconnect() {
        this.$store.commit('removeFriends');
        this.$store.commit('removeToken');
        store.state.peer.node.closeServer();
        store.clean();
        http.put(`clients/${this.client.id}/logout`)
            .then((response) => {

            })
            .catch((error) => {

            });
      }
    }
  }

</script>

<style lang="scss">
  @import '~styles/global';

  .home {
    display: flex;
    flex-direction: column;
    background: $primaryLightColor;
    height: 100%;

    .information-home {
      display: flex;
      .user {
        flex: 1;
        padding: 5px 5px 5px 10px;
        h1 {
          font-size: 18px;
          font-weight: 300;
          .pseudo {
            font-weight: 700;
          }
        }
      }

      .logout {
        flex: 1;
        justify-self: flex-end;
        text-align: end;
        padding: 5px 10px 5px 5px;
        align-self: center;
        a {
          color: $accentColor;
          text-decoration: none;
          &:hover {
            text-decoration: underline;
          }
        }
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