<template>
  <div v-if="type == 'pizza'">
    <ProductListNavItem class="product_list_nav" :type="type">
    </ProductListNavItem>
  </div>
</template>

<script>
export default {
  name: "ProductListNav",
  props: {
    type: {
      type: String,
    },
  },
};
</script>

<style scoped>
.product_list_nav {
  display: flex;
  font-size: 20px;
  background-color: coral;

  border-radius: 14px 10px 0 0;
}
</style>
