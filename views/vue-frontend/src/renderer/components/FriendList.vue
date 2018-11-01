<template>
  <div class="friends">
    <p>Friends :</p>
    <div v-for="friend in friends" class="friend-container">
      <div class="friend-info">
        {{friend.pseudo}}
      </div>
      <div v-if="friend.isConnected" class="connected">
        <div class="icon-container">
          <span class="connected-icon"></span>
        </div>
      </div>
    </div>
  </div>

</template>

<script>
  export default {
    name: "FriendList",
    computed: {
      connected() {
        return this.$store.state.Friends.connected;
      },
      disconnected() {
        return this.$store.state.Friends.disconnected;
      },
      friends() {
        return this.connected.concat(this.disconnected);
      }
    },
  }
</script>

<style lang="scss">
  @import '~styles/global';

  .friends {
    display: flex;
    padding-left: 10px;
    flex-direction: column;
    .friend-container {
      display: flex;
      flex-direction: row;
      .friend-info {
        flex: 2;
      }

      .connected {
        flex: 1;

        .icon-container {
          .connected-icon {
            height: 10px;
            width: 10px;
            background-color: lawngreen;
            border-radius: 50%;
            display: inline-block;
          }
        }
      }
    }
  }
</style>