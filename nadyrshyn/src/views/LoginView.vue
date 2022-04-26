<template>
  <div class="intro">
    <div class="title">
      <label class="title_text">{{ title }}</label>
    </div>
    <div class="login_form">
      <div v-if="status === 'login'">
        <div class="login_choice">
          <div class="login_choice_btn active">Войти</div>
          <div
            class="login_choice_btn"
            @click="
              status = 'signIn';
              title = 'Регистрация';
            "
          >
            Регистрация
          </div>
        </div>
      </div>
      <div v-else>
        <div class="login_choice">
          <div
            class="login_choice_btn"
            @click="
              status = 'login';
              title = 'Войти';
            "
          >
            Войти
          </div>
          <div class="login_choice_btn active">Регистрация</div>
        </div>
      </div>
      <LoginItem :status="status"></LoginItem>
    </div>
  </div>
</template>

<script>
import LoginItem from "@/components/LoginItem";

export default {
  name: "LoginView",
  components: { LoginItem },
  data() {
    return {
      status: "login",
      title: "Войти",
    };
  },
};
</script>

<style scoped>
.login_form {
  background-color: #686e65;
  display: flex;
  flex-direction: column;
  margin: 0 auto;
  width: 500px;
  min-height: max-content;
}
.login_choice {
  color: #d3c7c7;
  display: flex;
  flex-direction: row;
  width: 100%;
  background-color: #4b4242ff;
  margin: 0;
  height: 50px;
}
.login_choice_btn {
  font-size: 24px;
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100%;
  width: 50%;
}

.login_choice_btn.active {
  background-color: rgb(245, 208, 195);
  color: #222;
}

.login_choice_btn:hover {
  background-color: #eee;
  color: #333;
  cursor: pointer;
}

@media (max-width: 560px) {
  .login_form {
    max-width: 90vw;
  }
  .login_choice_btn {
    font-size: 20px;
  }
}
</style>
