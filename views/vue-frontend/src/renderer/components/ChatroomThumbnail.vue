<template>
  <div class="thumbnail">
    <div @click="select">
      {{chatroom.name}}
    </div>
  </div>
</template>

<script>
  export default {
    name: "ChatroomThumbnail",
    props: {
      chatroom: {
        type: Object,
      }
    },
    methods: {
      select() {
        this.$emit('select-chatroom', this.chatroom);
      }
    }
  }
</script>

<style scoped>

</style>