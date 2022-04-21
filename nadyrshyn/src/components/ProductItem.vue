<template>
  <div class="product_item">
    <div class="product_item_title">
      <img class="item_title_logo" :src="product.supplier_image" />
      <label class="title_name">{{ product.supplier_name }}</label>
    </div>
    <div class="product_img">
      <img
        class="product_logo"
        :src="product.image"
        @click="$router.push({ path: `/product/${product.id}` })"
      />
      <template v-if="counter > 0">
        <label class="product_counter">{{ counter }}</label>
      </template>
    </div>
    <div class="product_text">
      <label
        class="product_name"
        @click="$router.push({ path: `/product/${product.id}` })"
      >
        {{ product.name }}
      </label>
      <div class="product_price_container">
        <div class="product_type">{{ product.type }}</div>
        <div class="product_price_block">
          <label class="product_price">{{ product.price }}</label>
          <label class="product_price_grn">грн</label>
        </div>
      </div>
    </div>
    <button class="product_btn" v-on:click="addToBasket()">
      добавить в корзину
    </button>
  </div>
</template>

<script>
export default {
  name: "ProductItem",
  data() {
    return {
      counter: 0,
    };
  },
  props: {
    product: {
      type: Object,
    },
  },
  created() {
    for (let i = 0; i < this.$store.state.basket.products.length; i++)
      if (this.$store.state.basket.products[i].id === this.product.id) {
        this.counter = this.$store.state.basket.products[i].counter;
        break;
      }
  },
  methods: {
    addToBasket() {
      this.$store.dispatch("basket/addProduct", this.product);
      for (let i = 0; i < this.$store.state.basket.products.length; i++)
        if (this.$store.state.basket.products[i].id === this.product.id) {
          this.counter = this.$store.state.basket.products[i].counter;
          break;
        }
    },
  },
};
</script>

<style>
.product_img {
  position: relative;
  display: flex;
  max-width: 100%;
  height: 250px;
  align-items: center;
  justify-content: center;
}
</style>

<style scoped>
.product_item {
  display: flex;
  align-items: center;
  background-color: #8f968b;
  justify-content: space-between;
  flex-direction: column;
  color: #222;
  width: 251px;
  height: 420px;
  margin: 10px 13px;
  padding: 10px 10px 0 10px;
  border-radius: 10px;
}
.product_item:hover {
  background-color: #4b4242ff;
  color: #d3c7c7;
}

.product_item_title {
  display: flex;
  flex-direction: row;
  text-transform: uppercase;
  align-items: center;
  justify-content: center;
  width: 90%;
  height: 40px;
}
.product_counter {
  position: absolute;
  background-color: #d3c7c7;
  color: #222222;
  border: #111 solid 1px;
  border-radius: 5px;
  padding: 5px;
  opacity: 0.9;
  right: 10px;
  bottom: 10px;
  font-size: 28px;
}

.item_title_logo {
  display: block;
  border-radius: 5px;
  margin-right: 5px;
  max-height: 100%;
  max-width: 100%;
}

.title_name {
  font-size: 18px;
}

.product_logo {
  max-height: 200px;
  border: #222 solid 1px;
}
.product_logo:hover {
  cursor: pointer;
  transform: scale(1.03);
}

.product_text {
  display: flex;
  flex-direction: column;
  justify-content: space-between;
  width: 100%;
  margin: auto;
  min-height: 100px;
}
.product_name {
  text-transform: uppercase;
  font-size: 18px;
  font-weight: 700;
}
.product_name:hover {
  cursor: pointer;
}
.product_price_container {
  display: flex;
  justify-content: space-between;
  align-items: center;
}
.product_type {
  font-size: 16px;
  text-transform: uppercase;
}
.product_price_block {
  display: flex;
  align-items: end;
  justify-content: end;
  width: max-content;
}
.product_price {
  font-size: 30px;
}

.product_price_grn {
  margin: 5px;
  font-size: 16px;
}

.product_btn {
  margin: 10px auto;
  padding: 10px;
  background-color: #222222;
  color: #c2c5c1;
}
</style>
