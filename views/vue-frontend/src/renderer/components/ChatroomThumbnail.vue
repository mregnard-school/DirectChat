<template>
  <div class="thumbnail">
    <div @click="select" class="thumbnail-info">
      <div class="thumbnail-icon">
        <div class="thumbnail-name">
          {{name}}
        </div>
      </div>

      <div class="thumbnail-info">
        <div class="chatroom-infos">
          <div class="chatroom-name">
            {{chatroomName}}
          </div>
          <div class="chatroom-date">
            {{hour}}
          </div>
        </div>
        <div class="thumbnail-content">
          <div class="pseudo">
            <div v-if="isAuthor">
              You:
            </div>
            <div v-else>
              {{chatroom.last_message.author.pseudo}}:
            </div>
          </div>
          <div class="content">
            {{trimmed}}
          </div>
        </div>
      </div>

    </div>
  </div>
</template>

<script>
  import store from '@/mutableStore';
  import moment from 'moment';

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
      },
      chatroomName() {
        const name =  this.chatroom.name;
        return name.length > 10 ?
            name.substring(0, 10) + '...'
            : name;
      },
      name() {
        return this.chatroom.name.charAt(0);
      },
      hour() {
        let date = moment(this.message.date);
        if(date.isBefore(moment())) {
          return moment(date).format('HH:mm');
        } else {
          return moment(date).format('ddd');
        }
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
    text-align: center;

    &:hover {
      background: $primaryBox;
    }

    .thumbnail-info {
      display: flex;
      flex-direction: row;
      .thumbnail-icon {
        $iconSize: 30px;
        background: $chathead;
        border-radius: 50%;
        text-align: center;
        min-width: $iconSize;
        max-width: $iconSize;
        min-height: $iconSize;
        max-height: $iconSize;
        flex: 1;
        padding: 5px;
        .thumbnail-name {
          padding-top: 15%;
        }
      }

      .thumbnail-info {
        display: flex;
        flex-direction: column;
        padding-left: 10px;
        flex: 2;
        .chatroom-infos {
          display: flex;
          flex-direction: row;
          flex: 1;

          min-width: 100%;
          .chatroom-name {
            flex: 1;
            font-weight: 600;
            font-size: 16px;
            text-align: left;
            align-self: center;
          }

          .chatroom-date {
            align-self: center;
            justify-items: flex-end;
            justify-self: flex-end;
            flex-grow: 2;
            font-weight: normal;
            font-size: 14px;
            text-align: right;
            color: $lightGrey;
            padding-right: 3px;
          }
        }

        .thumbnail-content {
          flex: 1;
          align-self: flex-start;
          font-size: 14px;

          .pseudo {
            font-weight: normal;
            margin-right: 3px;
            display: inline-block;
          }
          .content {
            text-overflow: ellipsis;
            display: inline-block;
            font-weight: normal;
          }
        }
      }
    }
  }
</style>