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
      <div class="basket_title">
        <div class="price_block">
          <div class="price">{{ product.price }}</div>
          <label>грн</label>
        </div>
        <label class="title_product">{{ product.name }}</label>
      </div>
      <div class="basket_container">
        <div class="basket_item_ingredients">
          <div
            v-for="ingredient in product.ingredients"
            :key="ingredient.id"
            class="basket_item_ingredient"
          >
            <label>{{ ingredient }}</label>
          </div>
        </div>
        <div class="block">
          <div class="price">{{ product.counter }}</div>
          <div class="price_block">
            <div class="price">{{ summary.toFixed(2) }}</div>
            <label>грн</label>
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
.block {
  flex-direction: column;
}
</style>
