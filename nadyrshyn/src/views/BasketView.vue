<template>
  <div class="intro">
    <div class="section">
      <div class="container">
        <div class="title">
          <label class="title_text">Корзина</label>
        </div>
        <div class="basket_block">
          <div v-if="$store.state.basket.products.length === 0">
            <div class="clean_basket_text">Корзина пуста</div>
          </div>
          <div v-else>
            <div
              v-for="product in $store.getters['basket/getBasket']"
              :key="product.id"
            >
              <BasketItem :product="product">{{ products }}</BasketItem>
            </div>
          </div>
        </div>
        <div class="total_price_block">
          <div class="total_price_container">
            Всего:
            <div class="total_price">
              {{ $store.getters["basket/getTotalPrice"] }}
            </div>
            грн
          </div>
          <div class="basket_buttons">
            <button class="clear_btn red" v-on:click="clearBasket()">
              Очистить корзину
            </button>
            <button class="order_btn green" v-on:click="setOrder()">
              Заказать
            </button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import BasketItem from "@/components/BasketItem";

export default {
  name: "BasketView",
  components: { BasketItem },
  data() {
    return {
      products: [],
      price: Number,
      loaded: false,
      errors: [],
    };
  },
  created() {
    if (JSON.parse(localStorage.getItem("delivery_basket")) == null)
      localStorage.setItem("delivery_basket", JSON.stringify(this.products));
    this.$store.dispatch("basket/calcTotalPrice");
  },
  methods: {
    clearBasket() {
      this.$store.commit("basket/clearBasket");
      this.$store.dispatch("basket/calcTotalPrice");
    },
    setOrder() {
      let basket = this.$store.getters["basket/getBasket"];
      if (basket.length === 0) {
        alert("Корзина пустая!");
      } else {
        basket.totalPrice = this.$store.getters["basket/getTotalPrice"];
        localStorage.setItem("order", JSON.stringify(basket));
        this.$router.push({ path: `/order` });
      }
    },
  },
};
</script>

<style>
.basket_block {
  display: flex;
  flex-direction: column;
  text-align: center;
  align-items: normal;
  min-height: 150px;
  max-height: 490px;
  max-width: 1200px;
  margin: 0 auto 10px;
  border: 1px #222 solid;
  background-color: #686e65;
  overflow-y: scroll;
  scroll-behavior: smooth;
}
.total_price_block {
  display: flex;
  justify-content: space-between;
  align-items: center;
  flex-direction: row;
  margin: 10px;
}
.total_price_container {
  font-size: 30px;
  display: flex;
  flex-direction: row;
}
.total_price {
  margin: 0 10px;
}

@media (max-width: 810px) {
  .total_price_container {
    font-size: 18px;
  }
  .total_price {
    margin-right: 5px;
  }
}
@media (max-width: 450px) {
  .basket_buttons {
    display: flex;
    flex-direction: column-reverse;
    justify-content: space-evenly;
    align-items: center;
  }
}
</style>

<style scoped>
.clean_basket_text {
  margin: 40px auto;
}
.clear_btn {
  font-size: 14px;
  padding: 5px;
}
.order_btn {
  padding: 10px;
}
@media (max-width: 810px) {
  .order_btn {
    padding: 5px;
  }
}
@media (max-width: 450px) {
  .order_btn {
    margin-bottom: 5px;
    width: 100%;
  }
}
</style>
