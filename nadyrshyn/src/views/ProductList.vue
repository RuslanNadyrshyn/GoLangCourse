<template>
  <div>
    <div v-if="loaded">
      <template v-if="errors.length">
        <div v-for="(error, index) in errors" :key="index">
          {{ error }}
        </div>
      </template>

      <div class="Product_list">
        <div v-for="product in products" :key="product.id">
          <Product
            :id="product.id"
            :name="product.name"
            :image="product.image"
            :price="product.price"
            :type="product.type"
            :ingredients="product.ingredients"
          >
            {{ products }}
          </Product>
        </div>
      </div>
    </div>
    <div v-else>Loading...</div>
  </div>
</template>

<script>
import Product from "@/components/Product";
import axios from "axios";

export default {
  name: "ProductList",
  components: { Product },
  data() {
    return {
      products: [],
      loaded: false,
      errors: [],
    };
  },
  mounted() {
    const url = "http://localhost:8082/get_products";
    axios
      .get(url)
      .then((res) => {
        this.products = res.data;
        console.log(this.products);
      })
      .catch((err) => {
        this.errors.push(err);
      })
      .finally(() => {
        this.loaded = true;
      });
  },
};
</script>

<style scoped>
.Product_list {
  display: flex;
  flex-wrap: wrap;
  align-items: center;
  justify-content: space-evenly;
  width: 100%;
}
</style>
