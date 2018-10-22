<template>
  <div class="new-chatroom">
    <label for="friends">With whom ?</label>
    <input type="text" id="friends" v-model="pseudo">
    <button @click="submit">Create</button>
  </div>
</template>

<script>
  export default {
    name: "NewChatroom",
    data() {
      return {
        pseudo: "",
      }
    },
    methods: {
      submit() {
        // TODO irindul 2018-10-22 : Check if pseudo exist before emiting it
        // TODO irindul 2018-10-22 : Handle multiple accounts
        this.$emit('new-chatroom', [this.pseudo]);
      }
    }
  }
</script>

<style scoped>

</style>