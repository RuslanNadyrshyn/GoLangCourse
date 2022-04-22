<template>
  <div class="intro">
    <template v-if="$store.state.products.loaded">
      <template v-if="errors.length">
        <div v-for="(error, index) in errors" :key="index">
          {{ error }}
        </div>
      </template>
      <div v-for="product in selectedProduct" :key="product.id">
        <ProductBlock :product="product"></ProductBlock>
      </div>
    </template>
    <div v-else>Loading...</div>
  </div>
</template>

<script>
import ProductBlock from "@/components/ProductBlock";

export default {
  name: "ProductView",
  components: { ProductBlock },
  data() {
    return {
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
    if (JSON.parse(localStorage.getItem("delivery_basket")) == null)
      localStorage.setItem("delivery_basket", JSON.stringify([]));
    this.$store.dispatch("suppliers/fetchSuppliers");
    setTimeout(() => {
      let suppliers = this.$store.state.suppliers.suppliers;
      let prod = [];
      for (let i = 0; i < suppliers.length; i++)
        for (let j = 0; j < suppliers[i].menu.length; j++)
          prod.push(suppliers[i].menu[j]);

      this.$store.dispatch("products/fetchP", prod);
    }, 1000);
  },
};
</script>
<style scoped></style>
