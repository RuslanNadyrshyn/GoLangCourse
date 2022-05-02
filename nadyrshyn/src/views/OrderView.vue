<template>
  <div class="intro">
    <template v-if="$store.state.orders.errors.length">
      <div v-for="(error, index) in $store.state.orders.errors" :key="index">
        {{ error }}
      </div>
    </template>
    <template v-if="loaded">
      <div class="section">
        <div class="container">
          <div class="title">
            <label class="title_text">Заказ</label>
          </div>
          <div class="basket_block">
            <div v-for="product in order.products" :key="product.id">
              <OrderItem :product="product">{{ order.products }}</OrderItem>
            </div>
          </div>
          <div class="total_price_block">
            <div class="total_price_container">
              Всего:
              <div class="total_price">
                {{ order.totalPrice }}
              </div>
              грн
            </div>
            <button
              class="to_basket_btn red"
              v-on:click="$router.push({ path: `/basket` })"
            >
              Вернуться в корзину
            </button>
          </div>
          <div class="user_form">
            <div v-if="!$store.state.auth.Access" class="form_group">
              <label class="form-label">Имя:</label>
              <input
                v-model="order.user.name"
                class="form_control"
                type="text"
              />
            </div>
            <div class="form_group">
              <label class="form-label">Адрес:</label>
              <input v-model="order.address" class="form_control" type="text" />
            </div>
            <button class="order_btn green" v-on:click="setOrder()">
              Оформить заказ
            </button>
          </div>
        </div>
      </div>
    </template>
    <label v-else>Loading...</label>
  </div>
</template>

<script>
import OrderItem from "@/components/OrderItem";

export default {
  name: "OrderView",
  components: { OrderItem },
  data() {
    return {
      order: {
        products: JSON.parse(localStorage.getItem("delivery_basket")),
        totalPrice: 0,
        user: {
          name: "",
        },
        address: "",
      },
      loaded: false,
      errors: [],
      formData: {
        userId: "",
        title: "",
        body: "",
      },
    };
  },
  created() {
    this.$store.dispatch("basket/calcTotalPrice");
    this.order.totalPrice = this.$store.getters["basket/getTotalPrice"];
    if (this.order.products.length) this.loaded = true;
  },
  methods: {
    setOrder() {
      if (this.$store.getters["auth/getAccess"])
        this.order.user = this.$store.getters["auth/getUser"];
      if (
        this.order.address !== "" &&
        this.order.address.length < 50 &&
        this.order.user.name !== "" &&
        this.order.user.name.length < 20
      ) {
        this.$store.dispatch("orders/fetchOrderPOST", this.order).then(() => {
          this.errors = this.$store.getters["orders/getErrors"];
          if (!this.errors.length) {
            this.$store.commit("basket/clearBasket");
            this.$router.push("/");
          } else alert(this.errors);
        });
      } else {
        alert("Введите имя и адрес!");
      }
    },
  },
};
</script>

<style scoped>
.user_form {
  display: flex;
  flex-direction: column;
  justify-content: space-around;
  align-items: center;
  margin: 0 auto 20px;
  height: max-content;
  max-width: 400px;
  background-color: #686e65;
}
.to_basket_btn {
  padding: 10px;
}
.order_btn {
  margin: 10px;
  padding: 10px;
}
@media (max-width: 810px) {
  .to_basket_btn {
    font-size: 14px;
    padding: 5px;
  }
}
@media (max-width: 560px) {
  .user_form {
    max-width: 90vw;
  }
  .to_basket_btn {
    font-size: 10px;
  }
}
</style>
