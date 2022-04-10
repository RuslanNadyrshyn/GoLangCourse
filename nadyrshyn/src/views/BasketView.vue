<template>
  <div>
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
                v-for="product in $store.getters['basket/getBasket']"
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
          </div>
          <div class="basket_total_price_block">
            <div class="total_price">
              Всего:
              <div class="basket_total_price">
                {{ $store.getters["basket/getTotalPrice"] }}
              </div>
              грн
            </div>
            <div class="basket_buttons">
              <button class="clear red" v-on:click="clearBasket()">
                Очистить корзину
              </button>
              <button class="order green" v-on:click="setOrder()">
                Заказать
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>
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
      if (this.$store.getters["basket/getBasket"].length === 0) {
        console.log("Корзина пустая!");
        return;
      }
      localStorage.setItem(
        "order",
        JSON.stringify(this.$store.getters["basket/getBasket"])
      );
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
  max-height: 600px;
  max-width: 1200px;
  margin: 10px auto;
  border: 1px #222 solid;
  background-color: #686e65;
  overflow-y: scroll;
  scroll-behavior: smooth;
}

.clean_basket_text {
  margin: 40px auto;
}

.basket_total_price_block {
  display: flex;
  justify-content: space-between;
  align-items: center;
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

.clear {
  font-size: 14px;
  margin-right: 5px;
  padding: 6px;
}
.order {
  margin-left: 5px;
  padding: 10px;
}
</style>
