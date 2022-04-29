<template>
  <div class="intro">
    <template v-if="$store.state.products.loaded">
      <template v-if="errors.length">
        <div v-for="(error, index) in errors" :key="index">
          {{ error }}
        </div>
      </template>
      <ProductBlock
        :product="$store.state.products.product.Product"
        :supplier="$store.state.products.product.Supplier"
      ></ProductBlock>
    </template>
    <div v-else>Loading...</div>
  </div>
</template>

<script>
import ProductBlock from "@/components/ProductBlock";
import axios from "axios";

export default {
  name: "ProductView",
  components: { ProductBlock },
  data() {
    return {
      product: null,
      loaded: false,
      errors: [],
    };
  },
  created() {
    if (JSON.parse(localStorage.getItem("delivery_basket")) == null)
      localStorage.setItem("delivery_basket", JSON.stringify([]));
    this.$store.dispatch("products/fetchById", this.$route.params.id);
    // axios
    //   .get("http://localhost:8080/get_product", {
    //     params: { id: this.$route.params.id },
    //   })
    //   .then((res) => this.product = res.data)
    //   .catch((err) => console.log(err))
    //   .finally(() => {
    //     this.$store.commit("products/setLoaded", true);
    //   });
  },
};
</script>
<style scoped></style>
