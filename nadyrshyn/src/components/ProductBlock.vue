<template>
  <div class="product_block">
    <div class="product_block_item">
      <div class="product_img">
        <img class="product_logo" :src="image" alt="" />
      </div>
      <div class="product_text">
        <div class="product_name">
          {{ name }}
        </div>
        <div class="product_price_block">
          <div class="product_price">
            {{ price }}
            <div class="product_price_grn">грн</div>
          </div>
          <button class="add_to_basket green" v-on:click="addToBasket()">
            добавить в корзину
          </button>
        </div>
      </div>
    </div>
    <div class="product_description">
      <div class="product_description_title">
        {{ name }}
      </div>
      <div class="product_description_text">
        <div class="supplier_img">
          <img class="supplier_logo" :src="supplier_image" />
        </div>
        <div class="supplier_name">{{ supplier_name }}</div>
      </div>
      <div class="product_description_title">Ingredients:</div>
      <div class="product_description_ingredients">
        <div v-for="ingredient in ingredients" :key="ingredient.id">
          <div class="ingredient_item">{{ ingredient }}</div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "ProductBlock",
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
      let inBasket = false;
      for (let i = 0; i < this.$store.state.basket.products.length; i++)
        if (this.$store.state.basket.products[i].id === this.id) {
          this.$store.commit("basket/incCount", this.id);
          inBasket = true;
        }
      if (inBasket === false) this.$store.commit("basket/addProduct", prod);
    },
  },
};
</script>

<style scoped>
.product_block {
  display: flex;
  font: bold large sans-serif;
  flex-direction: row;
  min-height: 500px;
  max-width: 1200px;
  margin: 50px auto;
  background-color: #686e65;
}

.product_block_item {
  display: flex;
  flex-direction: column;
  min-height: 480px;
  width: 350px;
  margin: 30px 0 30px 30px;
  border: 2px #222;
  color: #222;
}
.product_img {
  display: flex;
  max-width: 100%;
  height: 250px;
  align-items: center;
  justify-content: center;
}
.product_logo {
  display: flex;
  max-width: 100%;
  max-height: 250px;
  background-color: #ddd;
  border: #222 solid 1px;
  border-radius: 10px;
  margin: auto;
}

.product_text {
  display: flex;
  font-size: 24px;
  flex-direction: column;
  padding: 7px;
}
.product_name {
  margin: 0 auto;
  text-align: center;
}

.product_price_block {
  display: flex;
  flex-direction: row;
  align-items: center;
  padding: 10px;
  justify-content: space-between;
  color: rgb(59, 31, 14);
}

.product_price {
  display: flex;
  flex-direction: row;
  font-size: 26px;
}

.product_price_grn {
  margin-left: 5px;
}

.add_to_basket {
  width: 120px;
}

.product_description {
  display: flex;
  flex-direction: column;
  background-color: #4a4e47;
  border-radius: 10px;
  border: solid 1px #222;
  min-height: 480px;
  width: 60%;
  margin: 30px;
  color: #c2c5c1;
}

.product_description_title {
  text-align: center;
  font-size: 30px;
  margin: 10px;
}

.product_description_text {
  font-size: 20px;
  margin: 10px;
  display: flex;
  flex-direction: row;
  align-items: center;
  width: 90%;
}
.supplier_logo {
  display: block;
  border-radius: 5px;
  margin-right: 15px;
  max-height: 60px;
  max-width: 100%;
}
.supplier_name {
  font-size: 28px;
}

.product_description_ingredients {
  display: flex;
  flex-direction: column;
  padding: 10px;
  background-color: #8f968b;
  border: solid #222 1px;
  border-radius: 5px;
  margin: 10px;
  font-size: 20px;
  color: #333;
}
</style>
