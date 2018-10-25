<template>
  <div class="message">
    <div v-if="isInformational" class="informational">
      {{message.content}}
    </div>
    <div v-else-if="isChangingName" class="name-changing">
      {{message.content}}
    </div>
    <div v-else class="">
      {{message.author.pseudo}} > {{message.content}}
    </div>

  </div>
</template>

<script>

  import types from '@/messageTypes';

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
      }
    }
  }
</script>

<style lang="scss">
  .message {
    .informational {
      font-size: 10px;
    }
    .name-changing {
      font-size: 10px;
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
