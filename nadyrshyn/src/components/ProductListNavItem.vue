<template>
  <div class="product_list_nav_item">
    {{ type }}
  </div>
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
.product_list_nav_item {
  padding: 10px;
  background-color: coral;
  color: #333;
  text-decoration: none;
  border-radius: 10px 10px 0 0;
  margin: 0px;
  text-transform: uppercase;
}
.product_list_nav_item:hover {
  background-color: rgb(245, 208, 195);
  cursor: pointer;
}
</style>
