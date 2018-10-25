<template>
  <div class="thumbnail">
    <div @click="select" class="thumbnail-info"> 
      {{chatroom.name}}
    </div>
  </div>
</template>

<script>
  export default {
    name: "ChatroomThumbnail", // TODO irindul 2018-10-25 : Add last message in thubnail
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

<style lang="scss">
  @import '~styles/global';
  .thumbnail {
    padding: 5px;
    margin-top: 2px;
    border-top: 1px solid $dividerColor;
    text-align: center;

    .thumbnail-info {

    }
  }
</style>