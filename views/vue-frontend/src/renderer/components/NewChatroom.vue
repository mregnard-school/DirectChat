<template>
  <div class="new-chatroom">
    <div class="input-container">
      <vue-tags-input
          v-model="tag"
          :tags="tags"
          :autocomplete-items="filteredItems"
          :add-only-from-autocomplete="true"
          @tags-changed="newTags => tags = newTags"
          :placeholder="'With whom ?'"
      />
    </div>

    <div class="submit">
      <button @click="submit">Create</button>
    </div>

  </div>
</template>

<script>
  import VueTagsInput from '@johmun/vue-tags-input';

  export default {
    name: "NewChatroom",
    components: {
      VueTagsInput,
    },
    data() {
      return {
        pseudo: "",
        tag: '',
        tags: [],
      }
    },
    computed: {
      connected() {
        return this.$store.state.Friends.connected;
      },
      disconnected() {
        return this.$store.state.Friends.disconnected;
      },
      autocompleteItems() {
        return this.connected.concat(this.disconnected).map(friend => {
          return {
            text: friend.pseudo,
          }
        });
      },
      filteredItems() {
        return this.autocompleteItems.filter(i => new RegExp(this.tag, 'i').test(i.text));
      },
      friendSelected() {
        return this.tags.map((tag) => tag.text);
      }
    },
    methods: {
      submit() {
        this.$emit('new-chatroom', this.friendSelected);
      }
    }
  }
</script>

<style lang="scss">
  @import '~styles/_variable';

  .new-chatroom {
    padding: 10px;
    display: flex;
    flex-direction: column;
    .input-container {
      font-size: 14px;

      input {
        border-radius: 0;
        padding: 0px;
        outline-width: 0;

        &:focus {
          box-shadow: none;
        }
      }

      .vue-tags-input {
        border-radius: 7px;

        .input {
          border-radius: 7px;
          padding: 4px 10px;
          border: none;
          border-bottom: 1px solid #ccc;
        }
      }

      .item.valid.selected-item {
        background-color: $primaryLightColor;
      }

      .tag {
        position: relative;
        border-radius: 7px;

        &.valid {
          background-color: $accentColor;
          opacity: 30%;
        }

        &.valid.deletion-mark {
          background-color: $deleter;
        }

        &:after {
          transition: transform .2s;
          position: absolute;
          content: '';
          height: 2px;
          width: 108%;
          left: -4%;
          top: calc(50% - 1px);
          background-color: #000;
          transform: scaleX(0);
        }
      }

      .deletion-mark :after {
        transform: scaleX(1);
      }

    }

    .submit {
      margin-top: 10px;
    }
  }
</style>