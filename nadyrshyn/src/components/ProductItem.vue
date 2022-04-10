<template>
  <div class="product_item">
    <div class="product_item_title">
      <img class="supplier_mini_logo" :src="supplier_image" />
      <div class="supplier_name">{{ supplier_name }}</div>
    </div>
    <div class="product_img">
      <img
        class="product_logo"
        :src="image"
        @click="$router.push({ path: `/product/${id}` })"
      />
    </div>
    <div class="product_text">
      {{ name }}
      <div class="product_price">
        {{ price }}
        <div class="product_price_grn">грн</div>
      </div>
    </div>
    <button class="product_btn" v-on:click="addToBasket()">
      добавить в корзину
    </button>
  </div>
</template>
<script>
export default {
  name: "ProductItem",
  data() {
    return {
      counter: 1,
    };
  },
  props: {
    id: {
      type: Number,
    },
    name: {
      type: String,
    },
    menuId: {
      type: Number,
    },
    price: {
      type: Number,
    },
    image: {
      type: String,
    },
    type: {
      type: String,
    },
    ingredients: {
      type: Array,
    },
    supplier_id: {
      type: Number,
    },
    supplier_name: {
      type: String,
    },
    supplier_image: {
      type: String,
    },
  },
  methods: {
    addToBasket() {
      let prod = {
        id: this.id,
        name: this.name,
        menuId: this.menuId,
        price: this.price,
        image: this.image,
        type: this.type,
        ingredients: this.ingredients,
        counter: this.counter,
        supplier_id: this.supplier_id,
        supplier_name: this.supplier_name,
        supplier_image: this.supplier_image,
      };
      console.log(prod);
      let a = false;
      for (let i = 0; i < this.$store.state.basket.products.length; i++) {
        if (this.$store.state.basket.products[i].id === this.id) {
          console.log("counter++ for ", this.id);
          this.$store.commit("basket/incCount", this.id);
          a = true;
          break;
        }
      }
      if (a === false) {
        this.$store.commit("basket/addProduct", prod);
      }
    },
  },
};
</script>

<style scoped>
.product_item {
  display: flex;
  align-items: center;
  background-color: #8f968b;
  justify-content: space-between;
  flex-direction: column;
  color: #222;
  width: 251px;
  height: 420px;
  margin: 10px 13px;
  padding: 10px 10px 0 10px;
  border-radius: 10px;
}
.product_item:hover {
  background-color: #4b4242ff;
  color: #d3c7c7;
}

.product_item_title {
  display: flex;
  flex-direction: row;
  align-items: center;
  justify-content: center;
  width: 90%;
}
.supplier_mini_logo {
  display: block;
  border-radius: 5px;
  margin-right: 5px;
  max-height: 40px;
  max-width: 100%;
}
.supplier_name {
  font-size: 18px;
}
.product_img {
  display: flex;
  max-width: 100%;
  height: 250px;
  align-items: center;
  justify-content: center;
}

.product_logo {
  display: block;
  border-radius: 10px;
  border: #222 solid 1px;
  max-width: 100%;
  max-height: 200px;
  transition: transform 0.2s;
}

.product_logo:hover {
  transform: scale(1.02);
  opacity: 0.9;
  cursor: pointer;
}

.product_text {
  display: flex;
  flex-direction: column;
  justify-content: space-between;
  font-size: 18px;
  font-weight: 700;
  width: 100%;
  text-transform: uppercase;
  margin: auto;
  min-height: 100px;
}

.product_price {
  display: flex;
  align-items: end;
  justify-content: end;
  font-size: 30px;
}

.product_price_grn {
  margin: 5px;
  font-size: 16px;
}

.product_btn {
  margin: 10px auto;
  padding: 10px;
  background-color: #222222;
  color: #c2c5c1;
}
</style>
