<template>
  <div>
    <div v-if="$store.state.products.loaded">
      <template v-if="errors.length">
        <div v-for="(error, index) in errors" :key="index">
          {{ error }}
        </div>
      </template>
      <HeaderItem></HeaderItem>
      <div class="intro">
        <div class="section">
          <div class="container">
            <div class="product_list_block">
              <div class="product_list_nav">
                <div v-for="type in productTypes" :key="type">
                  <ProductListNavItem :type="type">
                    {{ productTypes }}
                  </ProductListNavItem>
                </div>
              </div>
              <div class="Product_list">
                <div
                  v-for="product in $store.state.products.products"
                  :key="product.id"
                >
                  <ProductItem
                    :id="product.id"
                    :name="product.name"
                    :image="product.image"
                    :price="product.price"
                    :type="product.type"
                    :ingredients="product.ingredients"
                  >
                    {{ products }}
                  </ProductItem>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
    <div v-else>Loading...</div>
  </div>
</template>

<script>
import HeaderItem from "@/components/Header";
import ProductListNavItem from "@/components/ProductListNavItem";
import ProductItem from "@/components/ProductItem";

export default {
  name: "ProductList",
  components: { ProductItem, ProductListNavItem, HeaderItem },
  data() {
    return {
      suppliers: [],
      products: this.$store.dispatch("products/fetchProducts"),
      productTypes: [],
      selectedProductType: {},
      loaded: false,
      errors: [],
    };
  },
  mounted() {
    this.$store.dispatch("products/fetchProducts");
    // this.$store.dispatch("suppliers/fetchSuppliers");

    let products = this.$store.getters["products/getProducts"];
    this.productTypes.push("ALL");
    for (let i = 0; i < products.length; i++) {
      if (this.productTypes.includes(products[i].type) == false) {
        this.productTypes.push(products[i].type);
      }
    }
    console.log("fetched");
  },
};
</script>

<style scoped>
.intro {
  display: flex;
  flex-direction: column;
  justify-content: center;
  width: 100%;
  background-color: #ffcfb4;
  background-size: cover;
}
.list_title {
  display: flex;
  justify-content: center;
  color: #222;
  padding: 5px 5px 5px 0;
}
.product_list_block {
  display: flex;
  flex-direction: column;
  justify-content: center;
  background-color: #686e65;
  border-radius: 20px;
  margin-top: 10px;
}
.product_list_nav {
  display: flex;
  font-size: 20px;
  font-family: Arial, Helvetica, sans-serif;
  font-weight: 700;
  flex-wrap: wrap;
  background-color: coral;
  border-radius: 14px 10px 0 0;
}
.container {
  width: 100%;
  max-width: 1200px;
  margin: 0 auto;
}
.Product_list {
  display: flex;
  flex-wrap: wrap;
  align-items: center;
  justify-content: space-evenly;
  width: 100%;
}
</style>
