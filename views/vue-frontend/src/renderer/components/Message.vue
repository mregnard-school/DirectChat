<template>
  <div class="message">
    <div v-if="isInformational" class="informational">
      {{message.content}}
    </div>
    <div v-else-if="isChangingName" class="informational">
      {{message.content}}
    </div>

    <div v-else class="content"
         v-bind:class="{ author: isAuthor }">
      <div class="pseudo">
        {{message.author.pseudo}}
      </div>
      <div class="message-content">
        <span>
          {{message.content}}
        </span>
      </div>
    </div>

  </div>
</template>

<script>

  import types from '@/messageTypes';
  import store from '@/mutableStore';
  export default {
    name: "Message",
    props: {
      message: {
        type: Object,
      }
    },
    computed: {
      isInformational() {
        return this.message.type && (this.message.type === types.information);
      },
      isChangingName() {
        return this.message.type && (this.message.type === types.nameChange);
      },
      isAuthor() {
        return store.state.peer.client.id === this.message.author.id;
      }
    }
  }
</script>

<style lang="scss">
  @import '~styles/global';
  .message {
    margin-bottom: 4px;
    display: block;
    .informational {
      font-size: 12px;
    }

    .content {
      margin-bottom: 10px;

      .pseudo {
        font-size: 12px;
        margin-bottom: 3px;
        padding-left: 4px;
        color: $secondaryText;
      }
      .message-content {
        background: $primaryDarkColor;
        border-radius: 5px;
        padding: 5px;
        font-size: 14px;
        display: inline;

      }
    }

    .author {
      text-align: right;
    }

  }

  /*
  .speech-bubble {
	position: relative;
	background: #00aabb;
	border-radius: .4em;
}

.speech-bubble:after {
	content: '';
	position: absolute;
	right: 0;
	top: 50%;
	width: 0;
	height: 0;
	border: 20px solid transparent;
	border-left-color: #00aabb;
	border-right: 0;
	margin-top: -20px;
	margin-right: -20px;
}
   */
</style>
