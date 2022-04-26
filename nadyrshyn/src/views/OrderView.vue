<template>
  <div>
    <div class="intro">
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
            <div class="form_group">
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
    </div>
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
        products: this.$store.getters["basket/getBasket"],
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
  },
  methods: {
    setOrder() {
      if (this.order.address !== "" && this.order.user.name !== "") {
        this.$store.dispatch("orders/fetchOrderPOST", this.order);
        this.$store.commit("basket/clearBasket");
        this.$router.push("/");
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
}
</style>
