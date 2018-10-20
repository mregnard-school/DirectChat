<template>
  <div class="message">
    {{message.author.pseudo}} > {{message.content}}
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
  }
</script>

<style>

</style>