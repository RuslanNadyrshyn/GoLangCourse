<template>
  <div>
    <div v-if="$store.state.products.loaded">
      <template v-if="errors.length">
        <div v-for="(error, index) in errors" :key="index">
          {{ error }}
        </div>
      </template>
      <HeaderItem></HeaderItem>
      <div class="intro">
        <div class="section">
          <div class="container">
            <div class="basket_title">Корзина</div>
            <div class="basket_block">
              <div v-if="$store.state.basket.products.length === 0">
                <div class="clean_basket_text">Корзина пуста</div>
              </div>
              <div v-else-if="$store.state.basket.products.length !== 0">
                <div
                  v-for="product in $store.state.basket.products"
                  :key="product.id"
                >
                  <BasketItem
                    :id="product.id"
                    :menuId="product.menuId"
                    :name="product.name"
                    :image="product.image"
                    :price="product.price"
                    :type="product.type"
                    :ingredients="product.ingredients"
                    :counter="product.counter"
                  >
                    {{ products }}
                  </BasketItem>
                </div>
              </div>
              <div class="basket_total_price_block">
                <div class="total_price">
                  Всего:
                  <div class="basket_total_price">
                    {{ $store.state.basket.totalPrice }}
                  </div>
                  грн
                </div>
                <button class="basket_btn clear" v-on:click="clearBasket()">
                  Очистить корзину
                </button>
              </div>
              <button class="basket_btn" v-on:click="clearBasket()">
                Заказать
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>
    <div v-else>Loading...</div>
  </div>
</template>

<script>
import HeaderItem from "@/components/HeaderItem";
import BasketItem from "@/components/BasketItem";

export default {
  name: "BasketView",
  components: { BasketItem, HeaderItem },
  data() {
    return {
      products: [],
      price: Number,
      loaded: false,
      errors: [],
    };
  },
  created() {
    this.$store.dispatch("basket/calcTotalPrice");
  },
  methods: {
    clearBasket() {
      this.$store.commit("basket/clearBasket");
      this.$store.dispatch("basket/calcTotalPrice");
    },
  },
};
</script>

<style scoped>
.intro {
  display: flex;
  flex-direction: column;
  width: 100%;
  min-height: 100vh;
  background-color: #333;
  background-size: cover;
}

.container {
  width: 100%;
  max-width: 1200px;
  margin: 0 auto;
}

.basket_title {
  text-transform: uppercase;
  margin: 10px;
  font-size: 32px;
  text-align: center;
}

.basket_block {
  display: flex;
  flex-direction: column;
  text-align: center;
  align-items: normal;
  min-height: 170px;
  max-width: 1200px;
  margin: 20px auto;
  background-color: #686e65;
  border-radius: 20px;
}

.clean_basket_text {
  margin: 40px auto;
}

.basket_total_price_block {
  display: flex;
  justify-content: space-between;
  font-size: 30px;
  flex-direction: row;
  margin: 0 10px;
}

.total_price {
  display: flex;
  flex-direction: row;
}

.basket_total_price {
  margin: 0 10px;
}

.basket_btn {
  font-size: 20px;
  width: min-content;
  cursor: pointer;
  border-radius: 10px;
  padding: 10px;
  margin-bottom: 20px;
  margin-left: 10px;
  border: solid 1px #111;
  background-color: #3e8e41;
  transition: color 0.2s linear, transform 0.1s linear, scale 0.2s;
}

.basket_btn:hover {
  opacity: 0.8;
  cursor: pointer;
  transform: scale(1.05);
}

.basket_btn:active {
  opacity: 1;
  transform: scale(1);
}

.basket_btn.clear {
  width: max-content;
  margin-bottom: 0;
  font-size: 16px;
}
</style>
