<template>
  <div class="product_block">
    <div class="product_block_item">
      <label class="product_name">{{ product.name }}</label>
      <div class="block_img">
        <img :src="product.image" alt="product_image" />
      </div>
      <div class="product_text">
        <div class="product_price_block">
          <div class="product_price_container">
            {{ product.price }}
            <div class="product_price_grn">грн</div>
          </div>
          <button class="add_to_basket green" v-on:click="addToBasket()">
            добавить в корзину
          </button>
        </div>
      </div>
    </div>
    <div class="description_block">
      <div class="description_container">
        <label class="description_title">Ingredients:</label>
        <div class="ingredients_container">
          <div v-for="ingredient in product.ingredients" :key="ingredient.id">
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
  props: {
    product: {
      type: Object,
    },
  },
  methods: {
    addToBasket() {
      this.$store.dispatch("basket/addProduct", this.product);
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
  display: flex;
  align-items: center;
  justify-content: center;
  max-height: 400px;
  margin: auto;
}

.product_block_item {
  display: flex;
  flex-direction: column;
  width: 50%;
  margin: 20px;
  color: #222;
}
.product_text {
  display: flex;
  font-size: 32px;
  flex-direction: column;
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
  padding: 10px;
  justify-content: space-between;
  font-size: 40px;
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

.add_to_basket {
  font-size: 22px;
  padding: 15px 20px;
  border: #333 solid 1px;
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
  font-size: 30px;
  margin: 0 auto 10px;
}
.ingredients_container {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  margin: 0 auto;
  width: 90%;
  min-height: 150px;
  background-color: #444;
  color: #d3c7c7;
  border-radius: 10px;
  font-size: 26px;
}

.supplier_logo {
  display: block;
  border-radius: 5px;
  margin: auto;
  max-height: 200px;
  max-width: 80%;
}
</style>
