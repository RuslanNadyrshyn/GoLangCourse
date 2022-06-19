<template>
  <div class="intro">
    <div class="section">
      <div class="container">
        <template v-if="this.loaded">
          <template v-if="$store.state.products.errors.length">
            <div
              v-for="(error, index) in $store.state.products.errors"
              :key="index"
            >
              {{ error }}
            </div>
          </template>
          <ProductBlock
            :product="this.product.Product"
            :supplier="this.product.Supplier"
          ></ProductBlock>
        </template>
        <label v-else class="loading">Загрузка...</label>
      </div>
    </div>
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
      product: [],
      loaded: false,
    };
  },
  mounted() {
    if (JSON.parse(localStorage.getItem("delivery_basket")) == null)
      localStorage.setItem("delivery_basket", JSON.stringify([]));

    this.loaded = false;
    axios
      .get(this.$store.getters["products/getByIdURL"], {
        params: { id: this.$route.params.id },
      })
      .then((res) => {
        this.product = res.data;
      })
      .catch((err) => console.log(err))
      .finally(() => (this.loaded = true));
  },
};
</script>
<style scoped></style>
