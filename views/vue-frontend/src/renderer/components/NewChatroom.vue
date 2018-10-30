<template>
  <div class="new-chatroom">
    <div class="input">
      <label for="friends">With whom ?</label>
      <input type="text" autofocus="true" id="friends" v-model="pseudo" v-on:keyup.enter="submit">
    </div>

    <div class="submit">
      <button @click="submit">Create</button>
    </div>

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

<style lang="scss">
  @import '~styles/global';
  .new-chatroom {
    padding: 10px;
    display: flex;
    flex-direction: column;
    .input {
      font-size: 14px;
    }

    .submit {
      margin-top: 10px;
    }
  }
</style>