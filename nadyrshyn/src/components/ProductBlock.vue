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
          <button class="product_button">Add</button>
        </div>
      </div>
    </div>
    <div class="product_description">
      <div class="product_description_title">
        {{ name }}
      </div>
      <div class="product_description_text">Supplier</div>
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
  methods: {
    addToBracket() {
      this.counter++;
    },
    increment() {
      this.counter++;
    },
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
};
</script>

<style scoped>
.product_block {
  display: flex;
  flex-direction: row;
  background-color: #c8dcf4;
  min-height: 500px;
  width: 1200px;
  margin: 50px auto;
  background-color: #686e65;
  border-radius: 20px;
}

.product_block_item {
  display: flex;
  flex-direction: column;
  min-height: 480px;
  width: 30%;
  margin: 10px;
  border: 2px #222;
  color: #222;
}

.product_logo {
  display: flex;
  max-width: 100%;
  min-height: 150px;
  background-color: #ddd;
  border-radius: 20px;
  margin: 10px auto;
}

.product_text {
  display: flex;
  font-size: 24px;
  flex-direction: column;
  padding: 7px;
}
.product_name {
  margin: 0 auto;
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
  background-color: #71bf4d;
  border-radius: 10px;
  transition: color 0.2s linear;
}

.product_button:hover {
  background-color: #46792f;
}

.product_description {
  display: flex;
  flex-direction: column;
  min-height: 480px;
  width: 60%;
  margin: 10px;
  color: #222;
}

.product_description_title {
  text-align: center;
  font-size: 24px;
  margin: 10px;
}

.product_description_text {
  margin: 10px;
  font-size: 20px;
}

.product_description_ingredients {
  display: flex;
  flex-direction: column;
  margin: 10px;
  font-size: 20px;
}
</style>
