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
            <div class="basket_title">Basket</div>
            <div class="basket_block">
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
                >
                  {{ products }}
                  <button class="product_button" v-on:click="counter++">Add</button>
                  <div>{{ counter }}</div>
                </BasketItem>
              </div>
            </div>
            <div class="basket_total_price_block">Всего:
              <div class="basket_total_price">
                {{ $store.state.basket.price }}
              </div>
              грн
            </div>
            <button class="basket_order_btn">Заказать</button>
          </div>
        </div>
      </div>
    </div>
    <div v-else>Loading...</div>
  </div>
</template>

<script>
import HeaderItem from "@/components/Header";
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
};
</script>

<style scoped>
.basket_block {
  display: flex;
  flex-direction: column;
  align-items: center;
  min-height: 500px;
  max-width: 1200px;
  padding: 30px;
  background-color: #686e65;
  border-radius: 20px;
}

.basket_total_price_block {
  font-size: 30px;
  display: flex;
  flex-direction: row;
}

.basket_total_price {
  margin: 0 10px;
}

.basket_order_btn {
  font-size: 26px;
  width: min-content;
  cursor: pointer;
  border-radius: 6px;
  margin: 5px;
  border: solid 1px #111;
  background-color: #3e8e41;
  transition: 0.2s;
}
.basket_order_btn:hover {
  background-color: #404a40;
  color: white;
}
</style>
