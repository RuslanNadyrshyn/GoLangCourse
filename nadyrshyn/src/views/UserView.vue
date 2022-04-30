<template>
  <div class="intro">
    <div class="section">
      <div class="container">
        <template v-if="loaded">
          <template v-if="errors.length">
            <div v-for="(error, index) in errors" :key="index">
              {{ error }}
            </div>
          </template>
          <div class="title">
            <label class="title_text">Пользователь №{{ user.id }}</label>
          </div>
          <div class="user_container">
            <label>Id: {{ user.id }}</label>
            <label>Name: {{ user.name }}</label>
            <label>Email: {{ user.email }}</label>
            <label>Orders: {{ orders.length }}</label>
          </div>
          <div class="orders_container">
            <div class="order_item order_nav">
              <div class="order">Заказ:</div>
              <div class="order">Стоимость:</div>
              <div class="order wide">Адрес:</div>
              <div class="order wide">Дата:</div>
            </div>
            <div
              v-for="order in orders"
              :key="order.id"
            >
              <a
                class="order_item"
                @click="$router.push({ path: `/order/${order.id}` })">
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
        <div v-else>Loading...</div>
      </div>
    </div>
  </div>
</template>

<script>
import axios from "axios";

export default {
  name: "UserView",
  data() {
    return {
      user: null,
      orders: [],
      loaded: false,
      errors: [],
    };
  },
  created() {
    axios
      .get("http://localhost:8080/user", {
        params: {
          id: this.$route.params.id,
        }
      })
      .then((res) => {
        console.log("user", res.data);
        this.user = res.data;
      })
      .catch((err) => console.log(err))
      .finally(() => {
        console.log("finally");
        this.loaded = true;
      });

    axios
      .get("http://localhost:8080/orders", {
        params: {
          userId: this.$route.params.id,
        }
      })
      .then((res) => {
        console.log("orders", res.data);
        this.orders = res.data;
      })
      .catch((err) => console.log(err))
      .finally(() => {
        console.log("finally");
      });

    //  this.$store.dispatch("GetUserById", this.$route.params.id);
  },
}
</script>

<style scoped>
.user_container {
  font-size: 20px;
  background-color: #686e65;
  display: flex;
  flex-direction: column;
  margin: 20px auto;
  width: 500px;
}

.orders_container {
  font-size: 18px;
  display: flex;
  flex-direction: column;
  height: max-content;
  width: max-content;
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
  padding: 2px;
}
.wide {
  width: 200px;
}
</style>