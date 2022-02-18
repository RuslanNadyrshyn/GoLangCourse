<template>
  <a class="product_list_nav_type"
     @click="$router.push({name: 'Products'})">
    {{ type }}
  </a>
</template>

<script>
export default {
  name: "ProductListNavItem",
  props: {
    type: {
      type: String,
    },
  },
};
</script>

<style scoped>
.product_list_nav_type {
  padding: 5px 20px;
  background-color: coral;
  color: #333;
  text-decoration: none;
  border-radius: 10px 10px 0 0;
}
</style>