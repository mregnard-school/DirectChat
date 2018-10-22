<template>
  <div class="message">
    <div v-if="isInformational" class="informational">
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
      }
    }
  }
</script>

<style lang="scss">
  .message {
    .informational {
      font-size: 10px;
    }
  }
</style>