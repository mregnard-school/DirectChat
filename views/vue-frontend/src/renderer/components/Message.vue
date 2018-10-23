<template>
  <div class="message">
    <div v-if="isInformational" class="informational">
      {{message.content}}
    </div>
    <div v-else-if="isChangingName" class="name-changing">
      {{message.content}}
    </div>
    <div v-else>
      {{message.author.pseudo}} > {{message.content}}
    </div>

  </div>
</template>

<script>
  export default {
    name: "Message",
    props: {
      message: {
        type: Object,
      }
    },
    computed: {
      isInformational() {
        return this.message.type === "informational";
      },
      isChangingName() {
        return this.message.type === "name-changing"
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
</style>