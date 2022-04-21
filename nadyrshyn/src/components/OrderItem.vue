<template>
  <div class="basket_item">
    <div class="basket_item_image">
      <img
        :src="product.image"
        alt="product"
        @click="$router.push({ path: `/product/${product.id}` })"
      />
    </div>
    <div class="price_block">
      <label class="price_block_title">Цена:</label>
      <div class="price_container">
        <div class="price">{{ product.price }}</div>
        <label>грн</label>
      </div>
    </div>
    <div class="description_block">
      <label class="basket_item_title">
        {{ product.name }}
      </label>
      <div class="basket_item_description">
        <div class="basket_item_ingredients">
          <div v-for="ingredient in product.ingredients" :key="ingredient.id">
            <label class="basket_item_ingredient">{{ ingredient }}</label>
          </div>
        </div>
      </div>
    </div>
    <div class="price_block">
      <label class="price_block_title">Кол-во:</label>
      <div class="price">{{ product.counter }}</div>
    </div>
    <div class="price_block">
      <label class="price_block_title">Общая цена:</label>
      <div class="price_container">
        <div class="price">{{ summary.toFixed(2) }}</div>
        <label>грн</label>
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
