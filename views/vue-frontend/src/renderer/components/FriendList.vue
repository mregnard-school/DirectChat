<template>
  <div class="friends">
    <div class="added">
      <h1>Friends</h1>
      <div v-for="friend in friends" class="friend-container">
        <div class="friend-info">
          {{friend.pseudo}}
        </div>
        <div v-if="friend.isConnected" class="connected">
          <div class="icon-container">
            <span class="connected-icon"></span>
          </div>
        </div>
      </div>
    </div>
    <div class="pending">
    </div>

    <div>
      <button class="fab" id="show-modal" @click="showModal = true">+</button>
    </div>

    <modal v-if="showModal" @close="showModal = false">
      <div slot="header">
        <h3>Add a new friend ! </h3>
      </div>


      <div slot="body" class="add-input">
        <input type="text"
               v-model.trim="addFriend"
               id="adder"
               placeholder="Pseudo"
               autofocus
               @keyup.enter="handleAddNewFriend"
        >
      </div>

      <div slot="footer">
        <button class="btn" @click.stop="handleAddNewFriend">
          OK
        </button>
      </div>
    </modal>
  </div>

</template>

<script>
  import Modal from 'components/Modal';
  import {http} from '@/axios-wrapper';
  import store from '@/mutableStore';

  export default {
    name: "FriendList",
    components: {
      Modal,
    },
    data() {
      return {
        showModal: false,
        addFriend: '',
      }
    },
    computed: {
      connected() {
        return this.$store.state.Friends.connected;
      },
      disconnected() {
        return this.$store.state.Friends.disconnected;
      },
      friends() {
        return this.connected.concat(this.disconnected);
      },
      clientId() {
        return store.state.peer.client.id;
      }
    },
    methods: {
      handleAddNewFriend() {
        if (this.addFriend !== '') {
          http.post(`/clients/${this.clientId}/friends`, {
            pseudo: this.addFriend,
          }).then((response) => {
            let client = response.data;
            let oldFriendIds = this.friends.map((friend) => {
              return friend.id;
            });

            client.friends
                .filter(friend => {
                  return !oldFriendIds.includes(friend.id);
                })
                .forEach(friend => {
                  if (friend.ips && friend.ips.length > 0) {
                    this.$emit('new-connected', friend);
                  } else {
                    this.$emit('new-disconnected', friend);
                  }
                })
          }).catch((error) => {
            console.log(error);
          });

          this.showModal = false;
        }
      },
    }
  }
</script>

<style lang="scss">
  @import '~styles/global';

  .friends {
    .added {
      display: flex;
      padding-left: 10px;
      flex-direction: column;
      h1 {
        text-align: center;
        font-size: 16px;
      }

      .friend-container {
        display: flex;
        flex-direction: row;
        align-items: center;
        .friend-info {
          flex: 1;
        }
        .connected {

          flex: 1;
          .icon-container {
            .connected-icon {
              height: 10px;
              width: 10px;
              background-color: $accentColor;
              border-radius: 50%;
              display: inline-block;
            }
          }
        }
      }
    }
    .fab {
      position: absolute;
      bottom: 2%;
      right: 3%;
      border-radius: 100px;
      height: 30px;
      width: 30px;
      font-size: 20px;
    }

    .modal-mask {
      position: fixed;
      z-index: 9998;
      top: 0;
      left: 0;
      width: 100%;
      height: 100%;
      background-color: rgba(0, 0, 0, .5);
      display: table;
      transition: opacity .3s ease;
    }

    .modal-wrapper {
      display: table-cell;
      vertical-align: middle;
    }

    .modal-container {
      width: 300px;
      margin: 0px auto;
      padding: 20px 30px;
      background-color: #fff;
      border-radius: 5px;
      box-shadow: 0 2px 8px rgba(0, 0, 0, .33);
      transition: all .3s ease;
      font-family: Helvetica, Arial, sans-serif;
    }

    .modal-header {
      h3 {
        margin-top: 0;
        color: $accentColor;
      }

    }

    .modal-body {
      padding-left: auto;
      padding-right: auto;
      input {
        font-size: $fontSize;;
        max-width: 300px;
        min-height: 25px;
        padding: 10px;
        border-radius: 3px;
        border: 1px solid $dividerColor;
      }
    }

    .modal-footer {
      float: right;
    }

    .modal-default-button {
      //float: right;
    }

    /*
     * The following styles are auto-applied to elements with
     * transition="modal" when their visibility is toggled
     * by Vue.js.
     *
     * You can easily play with the modal transition by editing
     * these styles.
     */

    .modal-enter {
      opacity: 0;
    }

    .modal-leave-active {
      opacity: 0;
    }

    .modal-enter .modal-container,
    .modal-leave-active .modal-container {
      -webkit-transform: scale(1.1);
      transform: scale(1.1);
    }
  }
</style>