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
              <label class="user_text">
                Id: {{ $store.state.auth.user.id }}
              </label>
              <label class="user_text">
                Имя: {{ $store.state.auth.user.name }}
              </label>
              <label class="user_text">
                Email: {{ $store.state.auth.user.email }}
              </label>
              <label class="user_text">
                Заказов: {{ $store.state.auth.orders.length }}
              </label>
            </div>
            <template v-if="$store.state.auth.orders.length > 0">
              <div class="orders_container">
                <div class="order_nav">
                  <div class="order_cell">
                    <label class="cell_text">Заказ:</label>
                  </div>
                  <div class="order_cell">
                    <label class="cell_text">Цена:</label>
                  </div>
                  <div class="order_cell wide">
                    <label class="cell_text">Адрес:</label>
                  </div>
                  <div class="order_cell wide">
                    <label class="cell_text">Дата:</label>
                  </div>
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
                    <div class="order_cell">
                      <label class="cell_text">{{ order.id }}</label>
                    </div>
                    <div class="order_cell">
                      <label class="cell_text">{{ order.price }}</label>
                    </div>
                    <div class="order_cell wide">
                      <label class="cell_text small">{{ order.address }}</label>
                    </div>
                    <div class="order_cell wide">
                      <label class="cell_text small">
                        {{ order.created_at }}
                      </label>
                    </div>
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
    this.$store.dispatch("auth/fetchOrders");
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
  width: 700px;
  border: solid #222 2px;
  overflow-x: scroll;
}
.user_text {
  margin: 5px 0 5px 10px;
}
.logout_btn {
  position: absolute;
  top: 5px;
  right: 5px;
  padding: 5px;
}

.orders_container {
  display: flex;
  flex-direction: column;
  height: max-content;
  max-height: 500px;
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
  cursor: default;
  display: flex;
  flex-direction: row;
  justify-content: space-between;
  max-height: min-content;
  font-size: 18px;
}

.order_cell {
  display: flex;
  width: 20%;
  overflow: auto;
  align-items: center;
  border: solid #222 1px;
}

.wide {
  width: 30%;
}
.cell_text {
  font-size: 16px;
  margin: 3px;
}
.cell_text.small {
  font-size: 12px;
}

@media (max-width: 810px) {
  .user_container {
    width: 90%;
  }
  .orders_container {
    width: 90%;
  }
  .cell_text {
    font-size: 12px;
  }
}
</style>
