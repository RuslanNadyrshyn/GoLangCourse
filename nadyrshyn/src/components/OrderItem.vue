<template>
  <div class="basket_item">
    <div class="bi_image">
      <img
        :src="product.image"
        alt="product"
        @click="$router.push({ path: `/product/${product.id}` })"
      />
    </div>
    <div class="description_container">
      <div class="description_block">
        <label
          class="title_product"
          @click="$router.push({ path: `/product/${product.id}` })"
        >
          {{ product.name }}
        </label>
        <div class="ingredients_block">
          <a class="bi_image hidden">
            <img
              :src="product.image"
              alt="product"
              @click="$router.push({ path: `/product/${product.id}` })"
            />
          </a>
          <div class="bi_ingredients">
            <div
              v-for="ingredient in product.ingredients"
              :key="ingredient.id"
              class="bi_ingredient"
            >
              <label>{{ ingredient }}</label>
            </div>
          </div>
        </div>
      </div>
      <div class="price_block">
        <div class="price">{{ product.counter }}</div>
        <div class="price_item">
          <div class="price">{{ summary.toFixed(2) }}</div>
          <label>грн</label>
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
<style scoped></style>
