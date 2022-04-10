<template>
  <div>
    <div v-if="$store.state.products.loaded">
      <template v-if="errors.length">
        <div v-for="(error, index) in errors" :key="index">
          {{ error }}
        </div>
      </template>
      <div class="intro">
        <HeaderItem></HeaderItem>
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
                :suplier_id="product.supplier_id"
                :supplier_name="product.supplier_name"
                :supplier_image="product.supplier_image"
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
    if (JSON.parse(localStorage.getItem("delivery_basket")) == null)
      localStorage.setItem("delivery_basket", JSON.stringify(this.products));
    this.$store.dispatch("suppliers/fetchSuppliers");
    setTimeout(() => {
      let suppliers = this.$store.state.suppliers.suppliers;
      let prod = [];
      for (let i = 0; i < suppliers.length; i++)
        for (let j = 0; j < suppliers[i].menu.length; j++)
          prod.push(suppliers[i].menu[j]);

      this.$store.dispatch("products/fetchP", prod);
    }, 2000);
  },
};
</script>
<style scoped>
.intro {
  display: flex;
  flex-direction: column;
  justify-content: center;
  width: 100%;
  min-height: 100vh;
  background-color: #333;
}
.container {
  width: 100%;
  max-width: 1200px;
  margin: 0 auto;
}
</style>
