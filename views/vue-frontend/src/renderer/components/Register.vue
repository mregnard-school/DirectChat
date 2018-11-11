<template>
  <div class="register">
    <div class="register-form" @keypress.enter="register">
      <div class="register-input">
        <label for="pseudo">Pseudo</label>
        <input v-model.trim="pseudo" autofocus type="text" id="pseudo" placeholder="Enter your pseudo...">
      </div>

      <div class="register-input">
        <label for="password">Password</label>
        <input v-model.trim="password" type="password" id="password" placeholder="Enter password"/>
      </div>

      <div class="register-input">
        <label for="passwordConfirm">Password Confirmation</label>
        <input v-model.trim="passwordConfirm" type="password" id="passwordConfirm" placeholder="Enter password"/>
      </div>

      <div class="register-submit">
        <button @click="register">
          Register
        </button>
      </div>

      <div class="login-link">
        <router-link to="/">
          Already have an account ? Log in !
        </router-link>
      </div>

    </div>
  </div>
</template>

<script>
  import {http} from '@/axios-wrapper';
  import store from '@/mutableStore';
  import {ip} from '@/util';
  import Client from 'p2p/client/client';
  import {userAuthed} from '@/util';

  export default {
    name: "Register",
    data() {
      return {
        pseudo: '',
        password: '',
        passwordConfirm: '',
      }
    },
    methods: {
      register() {
        if(this.passwordConfirm === this.password) {
          http.post('/clients/new', {
            pseudo: this.pseudo,
            password: this.password,
            ips: [ip],
          }).then((response) => {
            let client = response.data;
            userAuthed(client);
          }).catch((error) => {
            console.error(error);
          })
        } else {
          //Notify user with validation error
        }
      },
    }
  }
</script>

<style lang="scss">
  @import '~styles/global';

  .register {
    display: flex;
    position: absolute;
    justify-content: center;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    background: $primaryColor;
    .register-form {
      display: flex;
      flex-direction: column;
      margin: auto;
      padding: 30px;
      border-radius: 5px;
      background: $primaryLightColor;
      .register-input {
        display: flex;
        flex-direction: column;
        color: $primaryText;
        margin-bottom: 15px;

        input {
          font-size: 16px;
          max-width: 300px;
        }

        label {
          font-size: 15px;
        }
      }

      .register-submit {
        min-width: 55.2px;
        max-width: 100px;
        align-self: flex-end;
      }
    }
  }
</style>