<template>
  <div class="thumbnail">
    <div @click="select" class="thumbnail-info">
      <div class="thumbnail-name">
        {{chatroom.name}}
      </div>
      <div class="thumbnail-message">
        <div class="pseudo">
          <div v-if="isAuthor">
            You :
          </div>
          <div v-else>
            {{chatroom.last_message.author.pseudo}} :
          </div>
        </div>
        <div class="content">
          {{trimmed}}
        </div>

      </div>
    </div>
  </div>
</template>

<script>
  import store from '@/mutableStore';

  export default {
    name: "ChatroomThumbnail",
    props: {
      chatroom: {
        type: Object,
      }
    },
    computed: {
      isAuthor() {
        return store.state.peer.client.id === this.message.author.id;
      },
      message() {
        return this.chatroom.last_message;
      },
      trimmed() {
        //let content = this.message.content;
        return this.message.content.length > 7 ?
            this.message.content.substring(0, 7) + '...'
            : this.message.content;
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

      .thumbnail-name {
      }

      .thumbnail-message {
        .pseudo {
          font-weight: bold;
          margin-right: 3px;
          display: inline-block;
        }
        .content {
          display: inline-block;
          font-weight: normal;
        }
      }
    }
  }
</style>