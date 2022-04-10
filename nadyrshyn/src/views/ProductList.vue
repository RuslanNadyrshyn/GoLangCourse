<template>
  <div>
    <div v-if="$store.state.suppliers.loaded">
      <template v-if="errors.length">
        <div v-for="(error, index) in errors" :key="index">
          {{ error }}
        </div>
      </template>
      <div v-if="$store.state.products.loaded">
        <HeaderItem></HeaderItem>
        <div class="intro">
          <div class="section">
            <div class="container">
              <div class="title">
                <div class="title_text">Доставка</div>
              </div>
              <div class="list_counter">
                Заведений найдено:
                {{ $store.state.suppliers.sortedSuppliers.length }}
              </div>
              <div class="list_block">
                <div class="list_nav">
                  <ListNavItem :type="'all'" :isSupplier="true"></ListNavItem>
                  <div
                    v-for="type in $store.getters[
                      'suppliers/getSuppliersTypes'
                    ]"
                    :key="type"
                  >
                    <ListNavItem :type="type" :isSupplier="true">
                      {{ suppliersTypes }}
                    </ListNavItem>
                  </div>
                  <ListNavItem :type="'Открыто'" :isSupplier="true">
                  </ListNavItem>
                </div>
                <div class="list">
                  <div
                    v-for="supplier in $store.getters[
                      'suppliers/getSortedSuppliers'
                    ]"
                    :key="supplier.id"
                  >
                    <SupplierItem
                      :id="supplier.id"
                      :name="supplier.name"
                      :type="supplier.type"
                      :image="supplier.image"
                      :workingHours="supplier.workingHours"
                      :menu="supplier.menu"
                    >
                      {{ suppliers }}
                    </SupplierItem>
                  </div>
                </div>
              </div>
              <div class="list_counter">
                Товаров найдено:
                {{ $store.state.products.sortedProducts.length }}
              </div>
              <div class="list_block">
                <div class="list_nav">
                  <ListNavItem :type="'all'" :isProduct="true"></ListNavItem>
                  <div
                    v-for="type in $store.getters['products/getProductsTypes']"
                    :key="type"
                  >
                    <ListNavItem :type="type" :isProduct="true">
                      {{ productsTypes }}
                    </ListNavItem>
                  </div>
                </div>
                <div class="list product_list">
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
                      :supplier_image="product.supplier_image"
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
      <div v-else>Loading</div>
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
    }, 1000);
  },
};
</script>

<style scoped>
.intro {
  display: flex;
  flex-direction: column;
  width: 100%;
  min-height: 100vh;
  background-color: #333;
  background-size: cover;
  font: bold large sans-serif;
}

.title {
  display: flex;
  justify-content: center;
  margin: 10px auto;
}
.title_logo {
}
.title_text {
  font-size: 32px;
  color: #d3c7c7;
}

.list_counter {
  font-size: 20px;
  padding: 5px;
  background-color: #4b4242ff;
  color: #d3c7c7;
  width: max-content;
}

.list_block {
  display: flex;
  flex-direction: column;
  justify-content: center;
  background-color: #686e65;
  margin-bottom: 40px;
}

.list_nav {
  display: flex;
  font-size: 20px;
  font-family: Arial, Helvetica, sans-serif;
  font-weight: 700;
  flex-wrap: wrap;
  background-color: #4b4242ff;
}
.container {
  width: 100%;
  max-width: 1200px;
  margin: 0 auto;
}
.list {
  display: flex;
  flex-wrap: wrap;
  justify-content: start;
  width: 100%;
  height: 220px;
  overflow-y: scroll;
  scroll-behavior: smooth;
  margin: 10px auto;
}
.list.product_list {
  min-height: 450px;
  max-height: 600px;
}
</style>
