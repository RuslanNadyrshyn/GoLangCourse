<template>
  <div class="intro">
    <template v-if="$store.getters['auth/getUserID'] == $route.params.id">
      <div class="section">
        <div class="container">
          <template v-if="$store.getters['auth/getLoaded']">
            <template v-if="$store.state.auth.errors.length">
              <div
                v-for="(error, index) in $store.getters['auth/getErrors']"
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
            <table>
              <tr>
                <th class="user_header">Id:</th>
                <td>{{ $store.state.auth.user.id }}</td>
              </tr>
              <tr>
                <th class="user_header">Имя:</th>
                <td>{{ $store.state.auth.user.name }}</td>
              </tr>
              <tr>
                <th class="user_header">Email:</th>
                <td>{{ $store.state.auth.user.email }}</td>
              </tr>
              <tr>
                <th class="user_header">Заказов:</th>
                <td>{{ $store.state.auth.orders.length }}</td>
              </tr>
            </table>
            <template v-if="$store.state.auth.orders.length > 0">
              <table class="orders">
                <tr style="background-color: #444">
                  <th>Заказ:</th>
                  <th>Цена:</th>
                  <th>Адрес:</th>
                  <th>Дата:</th>
                </tr>
                <tr
                  class="cells"
                  v-for="order in $store.getters['auth/getOrders']"
                  :key="order.id"
                  @click="
                    $router.push({
                      path: `/order/${order.id}`,
                    })
                  "
                >
                  <td>{{ order.id }}</td>
                  <td>{{ order.price }}$</td>
                  <td>{{ order.address }}</td>
                  <td>{{ order.created_at }}</td>
                </tr>
              </table>
              <button class="logout_btn" @click="Logout">Logout</button>
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
  mounted() {
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
.logout_btn {
  position: absolute;
  top: 50px;
  right: 10px;
  padding: 5px;
}
table {
  margin: 0 auto;
  width: 700px;
  border: solid #222 2px;
  background-color: #686e65;
  border-spacing: 0;
  border-collapse: collapse;
}
th {
  padding: 5px 0;
}
.user_header {
  width: 100px;
  text-align: left;
  padding: 5px 10px;
}

.orders {
  overflow-y: auto;
  scroll-behavior: smooth;
  text-align: center;
}
.cells {
  border: solid #222 1px;
}
.cells:hover {
  cursor: pointer;
  background-color: #333;
}
td {
  max-width: 70px;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
  padding: 5px 0;
}

@media (max-width: 810px) {
  table {
    width: 90%;
  }
  .orders {
    width: 90%;
  }
}
</style>
