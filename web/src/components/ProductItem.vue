<template>
  <div class="product_item">
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
  width: 250px;
  height: 370px;
  margin: 5px auto;
  padding: 10px 10px 0 10px;
  border-radius: 10px;
}
.product_item:hover {
  background-color: #4b4242ff;
  color: #d3c7c7;
}

.product_counter {
  position: absolute;
  background-color: #333;
  color: #c2c5c1;
  border: #111 solid 1px;
  border-radius: 5px;
  padding: 5px;
  right: 10px;
  bottom: 0;
  font-size: 28px;
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
  height: 100px;
}
.product_name {
  display: -webkit-box;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
  text-transform: uppercase;
  font-size: 18px;
  font-weight: 700;
  text-align: center;
  overflow: hidden;
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
  overflow: hidden;
  text-overflow: ellipsis;
  max-width: min-content;
}
.product_type:hover {
  overflow: visible;
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

@media (max-width: 560px) {
  .product_item {
    width: 200px;
    height: 280px;
  }
  .product_img {
    max-height: 150px;
  }
  .product_logo {
    max-height: 150px;
  }
  .product_counter {
    font-size: 16px;
  }
  .product_name {
    font-size: 16px;
  }
  .product_text {
    height: 80px;
  }
  .product_type {
    font-size: 12px;
  }
  .product_btn {
    font-size: 14px;
    margin-top: 0;
    padding: 5px;
  }
}
</style>
