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
            <div class="list_block">
              <div class="list_nav">
                <ListNavItem :type="'all'" :isSupplier="true"></ListNavItem>
                <div
                  v-for="type in $store.state.suppliers.suppliersTypes"
                  :key="type"
                >
                  <ListNavItem :type="type" :isSupplier="true">
                    {{ suppliersTypes }}
                  </ListNavItem>
                </div>
              </div>
              <div class="list">
                <div
                  v-for="supplier in $store.state.suppliers.suppliers"
                  :key="supplier.id"
                >
                  <SupplierItem
                    :id="supplier.id"
                    :name="supplier.name"
                    :type="supplier.type"
                    :image="supplier.image"
                    :opening="supplier.workingHours[0]"
                    :closing="supplier.workingHours[1]"
                  >
                    {{ suppliers }}
                  </SupplierItem>
                </div>
              </div>
            </div>
            <div class="list_block">
              <div class="list_nav">
                <ListNavItem :type="'all'" :isProduct="true"></ListNavItem>
                <div
                  v-for="type in $store.state.products.productsTypes"
                  :key="type"
                >
                  <ListNavItem :type="type" :isProduct="true">
                    {{ productsTypes }}
                  </ListNavItem>
                </div>
              </div>
              <div class="list">
                <div
                  v-for="product in $store.getters[
                    'products/getSortedProducts'
                  ]"
                  :key="product.id"
                >
                  <ProductItem
                    :id="product.id"
                    :menuId="product.menuId"
                    :name="product.name"
                    :image="product.image"
                    :price="product.price"
                    :type="product.type"
                    :ingredients="product.ingredients"
                    :supplier_id="product.supplier_id"
                    :supplier_name="product.supplier_name"
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
import HeaderItem from "@/components/HeaderItem";
import ListNavItem from "@/components/ListNavItem";
import ProductItem from "@/components/ProductItem";
import SupplierItem from "@/components/SupplierItem";

export default {
  name: "ProductList",
  components: { SupplierItem, ProductItem, ListNavItem, HeaderItem },
  data() {
    return {
      suppliers: [],
      products: [],
      productsTypes: [],
      suppliersTypes: [],
      loaded: false,
      errors: [],
    };
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
  background-color: #333;
  background-size: cover;
}
.list_block {
  display: flex;
  flex-direction: column;
  justify-content: center;
  background-color: #686e65;
  border-radius: 20px;
  margin: 20px auto;
}
.list_nav {
  display: flex;
  font-size: 20px;
  font-family: Arial, Helvetica, sans-serif;
  font-weight: 700;
  flex-wrap: wrap;
  background-color: #4b4242ff;
  border-radius: 14px 10px 0 0;
}
.container {
  width: 100%;
  max-width: 1200px;
  margin: 0 auto;
}
.list {
  display: flex;
  flex-wrap: wrap;
  align-items: center;
  justify-content: space-evenly;
  width: 100%;
}
</style>
