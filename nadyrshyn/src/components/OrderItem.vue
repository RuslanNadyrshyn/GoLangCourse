<template>
  <div class="basket_item">
    <div class="basket_item_image">
      <img
        :src="product.image"
        alt="product"
        @click="$router.push({ path: `/product/${product.id}` })"
      />
    </div>
    <div class="description_block">
      <div class="basket_item_title">
        <div class="price_block">
          <div class="price_container">
            <div class="price">{{ product.price }}</div>
            <label>грн</label>
          </div>
        </div>
        <label class="title_product">{{ product.name }}</label>
      </div>
      <div class="basket_item_title">
        <div class="basket_item_description">
          <div class="basket_item_ingredients">
            <div v-for="ingredient in product.ingredients" :key="ingredient.id">
              <label class="basket_item_ingredient">{{ ingredient }}</label>
            </div>
          </div>
        </div>
        <div class="block">
          <div class="price_block">
            <div class="price">{{ product.counter }}</div>
          </div>
          <div class="price_block">
            <div class="price_container">
              <div class="price">{{ summary.toFixed(2) }}</div>
              <label>грн</label>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "OrderItem",
  data() {
    return {
      summary: this.product.price * this.product.counter,
    };
  },
  props: {
    product: {
      type: Object,
    },
  },
};
</script>

<style scoped>
.title_product {
  font-size: 30px;
}
.description_block {
  width: 90%;
  justify-content: space-around;
}
.basket_item_title {
  margin: 0;
}
.basket_item_description {
  width: 80%;
}
.block {
  margin-left: 5px;
  display: flex;
  flex-direction: row;
  justify-content: center;
}
@media (max-width: 800px) {
  .basket_item {
    min-height: 120px;
  }
  .block {
    flex-direction: column;
    min-height: 70px;
  }
  .price_block {
    margin: 0;
  }
  .price {
    font-size: 24px;
    min-width: 40px;
    padding: 2px;
  }
}
</style>
