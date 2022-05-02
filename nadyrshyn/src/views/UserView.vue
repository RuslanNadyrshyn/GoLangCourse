<template>
  <div class="intro">
    <template v-if="$store.getters['auth/getUserID'] == $route.params.id">
      <div class="section">
        <div class="container">
          <template v-if="$store.state.auth.loaded">
            <template v-if="$store.state.auth.errors.length">
              <div
                v-for="(error, index) in $store.state.auth.errors"
                :key="index"
              >
                {{ error }}
              </div>
            </template>
            <div class="title">
              <label class="title_text">
                Пользователь №{{ $store.state.auth.user.id }}
              </label>
            </div>
            <div class="user_container">
              <button class="logout_btn" @click="Logout">Logout</button>
              <label>Id: {{ $store.state.auth.user.id }}</label>
              <label>Name: {{ $store.state.auth.user.name }}</label>
              <label>Email: {{ $store.state.auth.user.email }}</label>
              <label>Orders: {{ $store.state.auth.orders.length }}</label>
            </div>
            <template v-if="$store.state.auth.orders.length > 0">
              <div class="orders_container">
                <div class="order_item order_nav">
                  <div class="order">Заказ:</div>
                  <div class="order">Стоимость:</div>
                  <div class="order wide">Адрес:</div>
                  <div class="order wide">Дата:</div>
                </div>
                <div v-for="order in $store.state.auth.orders" :key="order.id">
                  <a
                    class="order_item"
                    @click="
                      $router.push({
                        path: `/order/${order.id}`,
                      })
                    "
                  >
                    <div class="order">{{ order.id }}</div>
                    <div class="order">{{ order.price }}</div>
                    <div class="order wide">
                      {{ order.address }}
                    </div>
                    <div class="order wide">{{ order.created_at }}</div>
                  </a>
                </div>
              </div>
            </template>
          </template>
          <label v-else class="loading">Загрузка...</label>
        </div>
      </div>
    </template>
    <div v-else>User is not authenticated</div>
  </div>
</template>

<script>
export default {
  name: "UserView",
  created() {
    if (this.$store.getters["auth/getAccess"])
      this.$store.dispatch("auth/fetchOrders", this.$route.params.id);
  },
  methods: {
    Logout() {
      console.log("logout");
      this.$store.dispatch("auth/Logout");
      this.$router.push({ path: `/` });
    },
  },
};
</script>

<style scoped>
.user_container {
  position: relative;
  font-size: 20px;
  background-color: #686e65;
  display: flex;
  flex-direction: column;
  margin: 20px auto;
  padding: 10px;
  width: 680px;
  border: solid #222 2px;
}
.logout_btn {
  position: absolute;
  top: 5px;
  right: 5px;
  padding: 5px;
}

.orders_container {
  font-size: 18px;
  display: flex;
  flex-direction: column;
  height: max-content;
  width: 700px;
  max-width: 1200px;
  margin: 0 auto;
  border: 1px #222 solid;
  background-color: #686e65;
  overflow-y: auto;
  scroll-behavior: smooth;
}

.order_item {
  display: flex;
  flex-direction: row;
  justify-content: space-between;
  max-height: min-content;
}

.order_item:hover {
  cursor: pointer;
  background-color: #333;
}

.order_nav {
  background-color: #444;
  cursor: none;
}

.order {
  display: flex;
  width: 150px;
  overflow: auto;
  align-items: center;
  border: solid #222 1px;
  padding: 3px;
}

.wide {
  width: 200px;
}
</style>
