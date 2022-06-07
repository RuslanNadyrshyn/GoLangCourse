<template>
  <div class="intro">
    <div class="section">
      <div class="container">
        <template v-if="$store.state.products.loaded">
          <template v-if="$store.state.products.errors.length">
            <div
              v-for="(error, index) in $store.state.products.errors"
              :key="index"
            >
              {{ error }}
            </div>
          </template>
          <ProductBlock
            :product="$store.state.products.product.Product"
            :supplier="$store.state.products.product.Supplier"
          ></ProductBlock>
        </template>
        <label v-else class="loading">Загрузка...</label>
      </div>
    </div>
  </div>
</template>

<script>
import ProductBlock from "@/components/ProductBlock";

export default {
  name: "ProductView",
  components: { ProductBlock },
  mounted() {
    if (JSON.parse(localStorage.getItem("delivery_basket")) == null)
      localStorage.setItem("delivery_basket", JSON.stringify([]));
    this.$store.dispatch("products/fetchById", this.$route.params.id);
  },
};
</script>
<style scoped></style>
