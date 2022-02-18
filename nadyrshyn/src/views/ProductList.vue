<template>
  <div>
    <div v-if="loaded">
      <template v-if="errors.length">
        <div v-for="(error, index) in errors" :key="index">
          {{ error }}
        </div>
      </template>
      <Header></Header>
      <div class="intro">
        <div class="section">
          <div class="container">
            <div class="product_list_block">
              <div v-for="product in products" :key="product.id">
                <ProductListNav :type="product.type"></ProductListNav>
              </div>
              <div class="Product_list">
                <div v-for="product in products" :key="product.id">
                  <Product
                    :id="product.id"
                    :name="product.name"
                    :image="product.image"
                    :price="product.price"
                    :type="product.type"
                    :ingredients="product.ingredients"
                  >
                    {{ products }}
                  </Product>
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
import Product from "@/components/Product";
import axios from "axios";
import Header from "@/components/Header";
import ProductListNav from "@/components/ProductListNav";

export default {
  name: "ProductList",
  components: { ProductListNav, Header, Product },
  data() {
    return {
      suppliers: [],
      products: [],
      loaded: false,
      errors: [],
    };
  },
  mounted() {
    // const suppliers_url = "http://localhost:8082/get_supppliers";
    // axios
    //     .get(suppliers_url)
    //     .then((res) => {
    //       this.suppliers = res.data;
    //       console.log(this.suppliers);
    //     })
    //     .catch((err) => {
    //       this.errors.push(err);
    //     })
    //     .finally(() => {
    //       this.loaded = true;
    //     });
    const url = "http://localhost:8082/get_products";
    axios
      .get(url)
      .then((res) => {
        this.products = res.data;
        console.log(this.products);
      })
      .catch((err) => {
        this.errors.push(err);
      })
      .finally(() => {
        this.loaded = true;
      });
  },
};
</script>

<style scoped>
.intro {
  display: flex;
  flex-direction: column;
  justify-content: center;
  width: 100%;
  background-color: #ffcfb4;
  background-size: cover;
}
.product_list_block {
  display: flex;
  flex-direction: column;
  justify-content: center;
  background-color: #686e65;
  border-radius: 20px;
  margin-top: 10px;
}
.product_list_nav {
  display: flex;
  font-size: 20px;
  background-color: coral;
  min-height: 50px;
  border-radius: 14px 10px 0 0;
}
.container {
  width: 100%;
  max-width: 1200px;
  margin: 0 auto;
}
.Product_list {
  display: flex;
  flex-wrap: wrap;
  align-items: center;
  justify-content: space-evenly;
  width: 100%;
}
</style>
