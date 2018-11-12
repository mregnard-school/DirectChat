<template>
  <div class="register">

    <div class="information">
      <h1>
        DirectChat
      </h1>

      <p>
        Create an account now ! And join others you could chat with :)
      </p>
    </div>
    <div class="register-form" @keypress.enter="register">
      <div class="register-input">
        <input v-model.trim="pseudo" autofocus type="text" id="pseudo" placeholder="Pseudo">
      </div>

      <div class="register-input">
        <input v-model.trim="password" type="password" id="password" placeholder="Password"/>
      </div>

      <div class="register-input">
        <input v-model.trim="passwordConfirm" type="password" id="passwordConfirm" placeholder="Confirm password"/>
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
    background: $primaryLightColor;
    flex-direction: column;

    .information {
      h1 {
        align-self: center;
        margin: 0;
      }

      p {
        font-size: $fontSize;;
        color: $secondaryText;
      }
      display: flex;
      flex-direction: column;
      margin-left: auto;
      margin-right: auto;
      margin-bottom: 0;
      padding-bottom: 10px;
    }

    .register-form {
      display: flex;
      flex-direction: column;
      margin-left: auto;
      margin-right: auto;
      padding: 30px;
      border-radius: 5px;
      background: $primaryLightColor;

      border: 1px solid $lightGrey;
      box-shadow: 10px 5px 5px $lightGrey;
      .register-input {
        display: flex;
        flex-direction: column;
        color: $primaryText;
        margin-bottom: 15px;

        input {
          font-size: $fontSize;;
          max-width: 300px;
          min-height: 25px;
          padding: 10px;
          border-radius: 3px;
          border: 1px solid $dividerColor;
        }

        label {
          font-size: $fontSize;;
        }
      }

      .register-submit {
        min-width: 55.2px;
        max-width: 100px;
        align-self: center;
        button {
          font-size: 20px;
          font-weight: 500;
          background: none;
          color: $accentColor;
        }
      }

      .login-link {
        padding-top: 5px;
        a {
          color: $accentColor;
          text-decoration: none;
          &:hover {
            text-decoration: underline;
          }
        }

      }
    }
  }
</style>