<template>
  <div class="login">
    <div class="login-form" @keypress.enter="login">
      <div class="login-input">
        <label for="pseudo">Pseudo</label>
        <input v-model.trim="pseudo" autofocus type="text" id="pseudo" placeholder="Enter your pseudo...">
      </div>

      <div class="login-input">
        <label for="password">Password</label>
        <input v-model.trim="password" type="password" id="password" placeholder="Enter password"/>
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
  import {userAuthed} from "@/util";

  export default {
    name: "Login",
    data() {
      return {
        pseudo: 'Billy',
        password: 'azerty',
        error: false,
      }
    },
    methods: {
      login() {
        if (this.pseudo !== '' && this.password !== '') {
          const payload = {
            pseudo: this.pseudo,
            password: this.password,
          };
          http.post('/clients/login', payload)
              .then((response) => {
                let client = response.data;
                userAuthed(client);
              })
              .catch(() => {
                this.error = true;
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
    background: $primaryColor;
    .login-form {
      display: flex;
      flex-direction: column;
      margin: auto;
      padding: 30px;
      border-radius: 5px;
      background: $primaryLightColor;
      .login-input {
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

      .login-submit {
        min-width: 55.2px;
        max-width: 100px;
        align-self: flex-end;
      }
    }
  }
</style>