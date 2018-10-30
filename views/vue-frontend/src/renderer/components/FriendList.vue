<template>
  <div>
    Friends lists :
    <div v-for="friend in friends" class="friends">
      <div class="friend-info">{{friend.pseudo}}
        <div v-if="friend.isConnected" class="connected">
          Connected ! <!-- todo change this with icon or smthg better-->
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

<style scoped>

</style>