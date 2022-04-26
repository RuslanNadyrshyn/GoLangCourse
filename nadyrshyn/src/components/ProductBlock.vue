<template>
  <div class="product_block">
    <div class="product_block_item">
      <label class="product_name">{{ product.name }}</label>
      <div class="block_img">
        <img class="block_logo" :src="product.image" alt="product_image" />
        <template v-if="counter > 0">
          <div class="counter_block">
            <label class="counter text">в корзине:</label>
            <label class="counter">{{ counter }}</label>
          </div>
        </template>
      </div>
      <div class="product_price_block">
        <div class="product_price_container">
          {{ product.price }}
          <div class="product_price_grn">грн</div>
        </div>
        <button class="add_to_basket_btn green" v-on:click="addToBasket()">
          добавить в корзину
        </button>
      </div>
    </div>
    <div class="description_block">
      <div class="description_container">
        <label class="description_title">тип:</label>
        <div class="ingredients_container">
          <label>{{ product.type }}</label>
        </div>
      </div>
      <div class="description_container">
        <label class="description_title">Ингредиенты:</label>
        <div class="ingredients_container">
          <div
            v-for="ingredient in product.ingredients"
            :key="ingredient.id"
            class="ingredient"
          >
            <label>{{ ingredient }}</label>
          </div>
        </div>
      </div>
      <div class="description_container">
        <label class="description_title">{{ product.supplier_name }}</label>
        <img class="supplier_logo" :src="product.supplier_image" />
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "ProductBlock",
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

<style scoped>
.product_block {
  display: flex;
  font: bold large sans-serif;
  flex-direction: row;
  min-height: 500px;
  max-width: 1200px;
  margin: 10px auto 0;
  background-color: #686e65;
}
.block_img {
  position: relative;
  display: flex;
  align-items: center;
  justify-content: center;
  margin: 0 auto;
}
.block_logo {
  min-height: 200px;
  max-height: 400px;
}
.counter_block {
  position: absolute;
  display: flex;
  flex-direction: row;
  align-items: end;
  width: max-content;
  height: min-content;
  right: 5px;
  bottom: 5px;
  font-size: 16px;
  text-transform: uppercase;
}
.counter {
  background-color: #444;
  color: #d3c7c7;
  width: max-content;
  border-radius: 5px 5px 5px 0;
  padding: 2px 4px;
  font-size: 24px;
  text-transform: uppercase;
}
.counter.text {
  font-size: 16px;
  border-radius: 5px 0 0 5px;
}
.product_block_item {
  display: flex;
  flex-direction: column;
  width: 50%;
  margin: 20px;
  color: #222;
}

.product_name {
  display: flex;
  font-size: 32px;
  margin: 0 auto;
  text-align: center;
}

.product_price_block {
  display: flex;
  flex-direction: row;
  align-items: center;
  justify-content: space-between;
  font-size: 40px;
  margin-top: 10px;
}

.product_price_container {
  display: flex;
  flex-direction: row;
  background-color: #444;
  border-radius: 5px;
  color: #d3c7c7;
  padding: 5px 10px;
}

.product_price_grn {
  margin-left: 5px;
}

.add_to_basket_btn {
  font-size: 18px;
  padding: 10px;
  border: #333 solid 1px;
  margin-left: 10px;
}

.description_block {
  display: flex;
  flex-direction: column;
  height: max-content;
  width: 50%;
  margin: 20px;
  color: #222;
}

.description_container {
  display: flex;
  flex-direction: column;
  margin: 10px 0;
}

.description_title {
  font-size: 26px;
  margin: 0 auto 10px;
  text-transform: uppercase;
}
.ingredients_container {
  display: flex;
  flex-direction: column;
  align-items: center;
  text-align: center;
  margin: 0 auto;
  padding: 5px 0;
  width: 90%;
  min-height: max-content;
  background-color: #444;
  color: #d3c7c7;
  border-radius: 10px;
  font-size: 26px;
  text-transform: uppercase;
}

.supplier_logo {
  display: block;
  border-radius: 5px;
  margin: auto;
  max-height: 200px;
  max-width: 80%;
}

@media (max-width: 810px) {
  .product_block {
    flex-direction: column;
    align-items: center;
  }
  .product_block_item {
    width: 90%;
    margin: 0;
  }
  .product_name {
    margin: 10px auto;
  }
  .product_price_block {
    font-size: 30px;
  }
  .description_block {
    width: 70%;
    margin: 0;
  }
}
@media (max-width: 560px) {
  .product_price_block {
    font-size: 26px;
  }
  .add_to_basket_btn {
    font-size: 14px;
  }
  .description_title {
    font-size: 22px;
    margin-bottom: 5px;
  }
  .description_container {
    margin: 5px 0;
  }
  .ingredients_container {
    font-size: 20px;
  }
}
</style>
