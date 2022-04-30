<template>
  <div class="intro">
    <template v-if="$store.state.suppliers.loaded">
      <template v-if="$store.state.suppliers.errors.length">
        <div
          v-for="(error, index) in $store.state.suppliers.errors"
          :key="index"
        >
          {{ error }}
        </div>
      </template>
      <template v-if="$store.state.products.loaded">
        <div class="section">
          <div class="container">
            <div class="title">
              <label class="title_text">Доставка</label>
            </div>
            <div class="list_block">
              <div class="list_nav">
                <ListNavItem :type="'все'" :isProduct="false"></ListNavItem>
                <ListNavItem :type="'Открыто'" :isProduct="false">
                </ListNavItem>
                <div
                  v-for="type in $store.getters['suppliers/getSuppliersTypes']"
                  :key="type"
                >
                  <ListNavItem :type="type" :isProduct="false"></ListNavItem>
                </div>
              </div>
              <div class="list">
                <div
                  v-for="supplier in $store.getters[
                    'suppliers/getSortedSuppliers'
                  ]"
                  :key="supplier.id"
                >
                  <SupplierItem :supplier="supplier"></SupplierItem>
                </div>
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
                <div
                  v-for="product in $store.getters[
                    'products/getSortedProducts'
                  ]"
                  :key="product.id"
                >
                  <ProductItem :product="product"></ProductItem>
                </div>
              </div>
            </div>
          </div>
        </div>
      </template>
      <div v-else>Loading products...</div>
    </template>
    <div v-else>Loading suppliers...</div>
  </div>
</template>

<script>
import ListNavItem from "@/components/ListNavItem";
import ProductItem from "@/components/ProductItem";
import SupplierItem from "@/components/SupplierItem";

export default {
  name: "ProductList",
  components: { SupplierItem, ProductItem, ListNavItem },
  data() {
    return {
      errors: [],
    };
  },
  created() {
    if (JSON.parse(localStorage.getItem("delivery_basket")) == null)
      localStorage.setItem("delivery_basket", JSON.stringify([]));
    if (this.$store.getters["suppliers/getSuppliers"].length === 0) {
      this.$store.dispatch("suppliers/fetchSuppliers");
      setTimeout(() => {
        let suppliers = this.$store.getters["suppliers/getSuppliers"];
        let prod = [];
        for (let i = 0; i < suppliers.length; i++)
          for (let j = 0; j < suppliers[i].menu.length; j++)
            prod.push(suppliers[i].menu[j]);

        this.$store.dispatch("products/fetchP", prod);
      }, 1000);
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
  display: flex;
  flex-wrap: wrap;
  justify-content: space-around;
  width: 100%;
  height: 210px;
  overflow-y: scroll;
  scroll-behavior: smooth;
  margin: auto;
}
.list.product_list {
  height: 880px;
}

@media (max-width: 560px) {
  .list {
    flex-wrap: nowrap;
    height: 150px;
  }
  .list.product_list {
    flex-wrap: wrap;
    height: 600px;
  }
}
</style>
