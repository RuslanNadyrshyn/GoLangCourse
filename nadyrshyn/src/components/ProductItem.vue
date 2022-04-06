<template>
  <div class="product_item">
    <div class="product_img">
      <img
        class="product_logo"
        :src="image"
        alt=""
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
      };
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
      console.log(this.$store.getters["basket/getBasket"]);
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
  max-width: 250px;
  height: 420px;
  margin: 10px;
  padding: 15px;
  border-radius: 20px;
}

.product_item:hover {
  opacity: 0.9;
  background-color: #4b4242ff;
  color: #d3c7c7;
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
  border-radius: 30px;
  border: #222 solid 1px;
  max-width: 100%;
  transition: transform 0.2s;
}

.product_logo:hover {
  transform: scale(1.05);
  opacity: 0.9;
  cursor: pointer;
}

.product_text {
  font-size: 20px;
  font-weight: 700;
  width: 100%;
  text-transform: uppercase;
  margin-top: 10px;
}

.product_price {
  display: flex;
  align-items: end;
  font-size: 30px;
}

.product_price_grn {
  margin: 5px;
  font-size: 16px;
}

.product_btn {
  font-size: 16px;
  width: 80%;
  text-transform: uppercase;
  height: max-content;
  padding: 10px;
  background-color: #2d8d0f;
  border-radius: 10px;
  border-style: solid;
  border-color: #222;
  transition: color 0.2s linear, transform 0.1s linear, scale 0.1s;
}

.product_btn:hover {
  background-color: #46792f;
  cursor: pointer;
  transform: scale(1.05);
}

.product_btn:active {
  background-color: #3e8e41;
  transform: scale(1);
}
</style>
