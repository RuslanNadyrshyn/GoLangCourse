<template>
  <div class="Product_item" @click="$router.push({ name: 'HomeView' })">
    <div class="Product_img">
      <img class="Product_logo" :src="image" alt="" />
    </div>
    <div class="Product_text">
      {{ name }}
      <div class="Product_price">
        {{ price }}
      </div>
    </div>
  </div>
</template>
<script>
export default {
  name: "ProductItem",
  props: {
    id: {
      type: Number,
    },
    name: {
      type: String,
    },
    menuId: {
      type: Number,
    },
    price: {
      type: Number,
    },
    image: {
      type: String,
    },
    type: {
      type: String,
    },
    ingredients: {
      type: Array,
    },
  },
};
</script>

<style scoped>
.Product_item {
  display: flex;
  flex-direction: column;
  color: #222;
  text-align: end;
  max-width: 250px;
  height: 350px;
  margin: 10px;
}

.Product_item:hover {
  opacity: 0.9;
  cursor: pointer;
}

.Product_img {
  display: flex;
  max-width: 100%;
  height: 250px;
  align-items: center;
}

.Product_logo {
  display: block;
  border-radius: 20%;
  max-width: 100%;
  transition: transform 0.2s;
}

.Product_logo:hover {
  transform: scale(1.05);
}

.Product_text {
  font-size: 20px;
  font-weight: 700;
  text-transform: uppercase;
  margin-top: 10px;
}
</style>
