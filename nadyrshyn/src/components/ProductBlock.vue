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
            <div class="product_price_grn">$</div>
          </div>
          <button class="product_button" v-on:click="addToBasket()">Add</button>
        </div>
      </div>
    </div>
    <div class="product_description">
      <div class="product_description_title">
        {{ name }}
      </div>
      <div class="product_description_text"></div>
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
      };
      let a = false;
      for (let i = 0; i < this.$store.state.basket.products.length; i++) {
        if (this.$store.state.basket.products[i].id === this.id) {
          console.log("counter++ for ", this.id);
          this.$store.commit("basket/incCount", this.id);
          a = true;
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
.product_block {
  display: flex;
  font: bold large sans-serif;
  flex-direction: row;
  min-height: 500px;
  width-max: 1200px;
  margin: 50px auto;
  background-color: #686e65;
  border-radius: 20px;
}

.product_block_item {
  display: flex;
  flex-direction: column;
  min-height: 480px;
  max-width: 400px;
  margin: 30px;
  border: 2px #222;
  color: #222;
}

.product_logo {
  display: flex;
  max-width: 100%;
  max-height: 300px;
  background-color: #ddd;
  border: #222 solid 2px;
  border-radius: 10%;
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
  margin-left: 10px;
}

.product_button {
  font-size: 20px;
  width: 100px;
  text-transform: uppercase;
  height: max-content;
  padding: 5px;
  background-color: #2d8d0f;
  border-radius: 10px;
  border-style: solid;
  border-color: #222;
  transition: color 0.2s linear;
}

.product_button:hover {
  background-color: #46792f;
  cursor: pointer;
}

.product_description {
  display: flex;
  flex-direction: column;
  background-color: #4a4e47;
  border-radius: 20px;
  border: solid 1px;
  min-height: 480px;
  width: 60%;
  margin: 30px;
  color: #222;
}

.product_description_title {
  text-align: center;
  font-size: 30px;
  margin: 10px;
}

.product_description_text {
  margin: 10px;
  font-size: 20px;
  font-style: italic;
}

.product_description_ingredients {
  display: flex;
  flex-direction: column;
  margin: 10px;
  font-size: 20px;
  font-style: italic;
}
</style>
