<template>
  <div class="login">
    <div class="information">
      <h1>
        DirectChat
      </h1>

      <p>
        Chat with your friends in a purely decentralized way !
      </p>
    </div>

    <div class="login-form" @keypress.enter="login">
      <div class="login-input">
        <input v-model.trim="pseudo" autofocus type="text" id="pseudo" placeholder="Pseudo">
      </div>

      <div class="login-input">
        <input v-model.trim="password" type="password" id="password" placeholder="Password"/>
      </div>

      <div class="error" v-if="error">
        Could not sign in, please check credentials
      </div>

      <div class="login-submit">
        <button @click="login">
          Log in
        </button>
      </div>

      <div class="register-link">
        <router-link to="/register">
          Don't have an account ? Create one now !
        </router-link>
      </div>

    </div>
  </div>
</template>

<script>
  import {http} from '@/axios-wrapper';
  import {ip, userAuthed} from "@/util";

  export default {
    name: "Login",
    data() {
      return {
        pseudo: '',
        password: '',
        error: false,
      }
    },
    methods: {
      login() {
        if (this.pseudo !== '' && this.password !== '') {
          const payload = {
            pseudo: this.pseudo,
            password: this.password,
            ips: [ip],
          };
          http.post('/clients/login', payload)
              .then((response) => {
                let client = response.data;
                userAuthed(client);
              })
              .catch((error) => {
                this.error = true;
                console.log(error);
              })
        }
      },
    },
    watch: {
      pseudo: function() {
        this.error = false;
      },
      password: function() {
        this.error = false;
      }
    }
  }
</script>

<style lang="scss">
  @import '~styles/global';

  .login {

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

    .login-form {
      display: flex;
      flex-direction: column;
      margin-left: auto;
      margin-right: auto;
      padding: 30px;
      border-radius: 5px;
      background: $primaryLightColor;
      border: 1px solid $lightGrey;
      box-shadow: 10px 5px 5px $lightGrey;
      .login-input {
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

      .login-submit {
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

      .register-link {
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