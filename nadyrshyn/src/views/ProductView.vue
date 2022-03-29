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
            <div v-for="product in selectedProduct" :key="product.id">
              <ProductBlock
                :id="product.id"
                :menu-id="product.menuId"
                :name="product.name"
                :image="product.image"
                :price="product.price"
                :type="product.type"
                :ingredients="product.ingredients"
              >
                {{ product }}
              </ProductBlock>
            </div>
          </div>
        </div>
      </div>
    </div>
    <div v-else>Loading...</div>
  </div>
</template>

<script>
import HeaderItem from "@/components/HeaderItem";
import ProductBlock from "@/components/ProductBlock";

export default {
  name: "ProductView",
  components: { HeaderItem, ProductBlock },
  data() {
    return {
      product: [],
      loaded: false,
      errors: [],
    };
  },
  computed: {
    selectedProduct() {
      return this.$store.state.products.products.filter(
        (product) => product.id == this.$route.params.id
      );
    },
  },
  created() {
    this.$store.dispatch("products/fetchProducts");
    this.$store.dispatch("suppliers/fetchSuppliers");
  },
};
</script>
<style scoped>
.intro {
  display: flex;
  flex-direction: column;
  justify-content: center;
  width: 100%;
  min-height: 500px;
  background-color: #ffcfb4;
  background-size: cover;
}
.container {
  width: 100%;
  max-width: 1200px;
  margin: 0 auto;
}
</style>
