<template>
  <div>
    <template v-if="status === 'login'">
      <div class="login_block">
        <div class="form_group">
          <label class="form_label">E-mail:</label>
          <input v-model="login.email" class="form_control" type="email" />
        </div>
        <div class="form_group">
          <label>Пароль:</label>
          <input
            v-model="login.password"
            class="form_control"
            type="password"
          />
        </div>
        <div class="form_group">
          <button class="reg_btn" type="submit" @click="Login">Вход</button>
        </div>
      </div>
    </template>
    <template v-else>
      <div class="login_block">
        <div class="form_group">
          <label>E-mail:</label>
          <input v-model="sign_in.email" class="form_control" type="email" />
        </div>
        <div class="form_group">
          <label>Пароль:</label>
          <input
            v-model="sign_in.password"
            class="form_control"
            type="password"
          />
        </div>
        <div class="form_group">
          <label>Имя:</label>
          <input v-model="sign_in.name" class="form_control" type="text" />
        </div>
        <div class="form_group">
          <label>Возраст:</label>
          <input
            v-model.number="sign_in.age"
            class="form_control"
            type="number"
            min="0"
            required
          />
        </div>
        <div class="form_group">
          <button class="reg_btn" type="submit" @click="SignIn">
            Регистрация
          </button>
        </div>
      </div>
    </template>
  </div>
</template>

<script>
import axios from "axios";

export default {
  name: "LoginItem",
  data() {
    return {
      login: {
        email: "",
        password: "",
      },
      sign_in: {
        name: "",
        email: "",
        password: "",
        age: Number,
      },
    };
  },
  props: {
    status: {
      type: String,
    },
  },
  methods: {
    Login() {
      console.log(this.login);
    },
    SignIn() {
      console.log(this.sign_in);
      axios
        .post("http://localhost:8080/sign_in", this.sign_in)
        .then((res) => console.log(res))
        .catch((err) => console.log(err));
    },
  },
};
</script>

<style>
.form_group {
  font-size: 20px;
  display: flex;
  flex-direction: column;
  justify-content: space-between;
  width: 300px;
  margin: 10px;
}
.form_control {
  font-size: 20px;
  border-radius: 5px;
  text-align: center;
}
</style>

<style scoped>
.login_block {
  display: flex;
  flex-direction: column;
  justify-content: space-evenly;
  align-items: center;
  width: 100%;
}
.reg_btn {
  padding: 10px;
  margin: 0 auto;
  background-color: #dddddd;
  min-width: 100px;
}
</style>
