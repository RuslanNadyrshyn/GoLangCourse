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
          <input v-model="sign_up.email" class="form_control" type="email" />
        </div>
        <div class="form_group">
          <label>Пароль:</label>
          <input
            v-model="sign_up.password"
            class="form_control"
            type="password"
          />
        </div>
        <div class="form_group">
          <label>Имя:</label>
          <input v-model="sign_up.name" class="form_control" type="text" />
        </div>
        <div class="form_group">
          <button class="reg_btn" type="submit" @click="SignUp">
            Регистрация
          </button>
        </div>
      </div>
    </template>
  </div>
</template>

<script>
export default {
  name: "LoginItem",
  data() {
    return {
      login: {
        email: "",
        password: "",
      },
      sign_up: {
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
    ValidateEmail(email) {
      return String(email)
        .toLowerCase()
        .match(
          /^(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/
        );
    },
    Login() {
      if (this.ValidateEmail(this.login.email) && this.login.password.length) {
        this.$store.dispatch("auth/Login", this.login);
      } else alert("Неверный ввод!");
    },
    SignUp() {
      const validateEmail = (email) => {
        return String(email)
          .toLowerCase()
          .match(
            /^(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/
          );
      };
      if (
        this.sign_up.name.length &&
        this.sign_up.name.length < 20 &&
        this.sign_up.email.length &&
        this.sign_up.email.length < 20 &&
        this.sign_up.password.length >= 4 &&
        validateEmail(this.sign_up.email)
      ) {
        this.$store.dispatch("auth/SignUp", this.sign_up);
      } else
        alert(
          "Неверный ввод!\n" +
            "Все поля дожны быть заполнены!\n" +
            "Длина имени и email не должна превышать 20 символов!\n" +
            "Минимальная длина пароля: 4 символа"
        );
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
  width: 90%;
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
