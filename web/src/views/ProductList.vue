<template>
  <div class="intro">
    <template v-if="$store.state.suppliers.errors.length">
      <div v-for="(error, index) in $store.state.suppliers.errors" :key="index">
        {{ error }}
      </div>
    </template>
    <template v-if="$store.state.products.errors.length">
      <div v-for="(error, index) in $store.state.products.errors" :key="index">
        {{ error }}
      </div>
    </template>
    <div class="section">
      <div class="container">
        <div class="title">
          <label class="title_text">Доставка</label>
        </div>
        <div class="list_block">
          <div class="list_nav">
            <ListNavItem :type="'все'" :isProduct="false"></ListNavItem>
            <ListNavItem :type="'Открыто'" :isProduct="false"></ListNavItem>
            <div
              v-for="type in $store.getters['suppliers/getSuppliersTypes']"
              :key="type"
            >
              <ListNavItem :type="type" :isProduct="false"></ListNavItem>
            </div>
          </div>
          <div class="list">
            <template v-if="$store.state.suppliers.loaded">
              <div
                v-for="supplier in $store.getters['suppliers/getSuppliers']"
                :key="supplier.id"
              >
                <SupplierItem :supplier="supplier"></SupplierItem>
              </div>
            </template>
            <label v-else class="loading">Загрузка...</label>
          </div>
        </div>
        <div class="list_block">
          <div class="list_nav">
            <ListNavItem :type="'все'" :isProduct="true"></ListNavItem>
            <div
              v-for="type in $store.getters['products/getProductsTypes']"
              :key="type"
            >
              <ListNavItem :type="type" :isProduct="true"></ListNavItem>
            </div>
          </div>
          <div class="list product_list">
            <template v-if="$store.state.products.loaded">
              <div
                v-for="product in $store.getters['products/getProducts']"
                :key="product.id"
              >
                <ProductItem :product="product"></ProductItem>
              </div>
            </template>
            <label v-else class="loading">Загрузка...</label>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import ListNavItem from "@/components/ListNavItem";
import ProductItem from "@/components/ProductItem";
import SupplierItem from "@/components/SupplierItem";

export default {
  name: "ProductList",
  components: { SupplierItem, ProductItem, ListNavItem },
  mounted() {
    if (JSON.parse(localStorage.getItem("delivery_basket")) == null)
      localStorage.setItem("delivery_basket", JSON.stringify([]));
    if (this.$store.getters["suppliers/getSuppliers"].length === 0) {
      this.$store.dispatch("suppliers/fetchSuppliers");
      this.$store.dispatch("products/fetchProducts");
    }
  },
};
</script>

<style scoped>
.list_block {
  display: flex;
  flex-direction: column;
  justify-content: center;
  background-color: #686e65;
  margin-bottom: 40px;
}

.list_nav {
  position: relative;
  display: flex;
  font-size: 20px;
  font-family: Arial, Helvetica, sans-serif;
  font-weight: 700;
  background-color: #4b4242ff;
  overflow-y: scroll;
  scroll-behavior: smooth;
}

.list {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(200px, 1fr));
  width: 100%;
  height: 210px;
  overflow-y: scroll;
  scroll-behavior: smooth;
  margin: auto;
}

.list.product_list {
  height: 880px;
  grid-template-columns: repeat(auto-fill, minmax(280px, 1fr));
}

@media (max-width: 560px) {
  .list {
    height: 150px;
    grid-template-columns: repeat(auto-fill, minmax(130px, 1fr));
  }
  .list.product_list {
    flex-wrap: wrap;
    height: 600px;
    grid-template-columns: repeat(auto-fill, minmax(230px, 1fr));
  }
}
</style>
