<template>
  <div class="basket_item">
    <div class="basket_item_image">
      <img class="product_logo" :src="image" alt="" />
    </div>
    <div class="basket_item_description">
      <div class="basket_item_title">
        {{ name }}
      </div>
      <div class="basket_item_ingredients">
        <div v-for="ingredient in ingredients" :key="ingredient.id">
          <div class="basket_item_ingredient">{{ ingredient }}</div>
        </div>
      </div>
    </div>
    <div class="basket_item_counter">
      <button class="counter_btn" v-on:click="counter++">+</button>
      <div class="counter_number"></div>
      <button class="counter_btn" v-on:onclick="counter--">-</button>
    </div>
    <div class="basket_item_price_block">
      <div class="basket_item_price">{{ price }}</div>
      <button class="basket_item_remove_btn">Remove</button>
    </div>
  </div>
</template>

<script>
export default {
  name: "BasketItem",
  data() {
    return {
      counter: 0,
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
};
</script>

<style scoped>
.basket_item {
  display: flex;
  flex-direction: row;
  justify-content: space-around;
  align-items: center;
  min-height: 100px;
  background-color: #4a4e47;
  width: 100%;
  margin-bottom: 10px;
  border-radius: 20px;
}

.basket_item_image {
  border-radius: 20px;
  cursor: pointer;
  width: 150px;
}

.product_logo {
  display: flex;
  max-width: 100%;
  min-height: 150px;
  background-color: #ddd;
  border-radius: 20px;
  margin: 10px auto;
}

.basket_item_description {
  display: flex;
  flex-direction: column;
  height: 110px;
  width: 60%;
  padding: 10px;
  background-color: #686e65;
  border: 1px solid #333;
  border-radius: 20px;
}

.basket_item_title {
  margin: 5px;
  cursor: pointer;
  text-align: center;
}

.basket_item_ingredients {
  display: flex;
  flex-wrap: wrap;
}

.basket_item_ingredient {
  margin: 0 10px;
}

.basket_item_counter {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  width: 7%;
  height: 150px;
}

.counter_btn {
  text-align: center;
  font-size: 24px;
  width: 40px;
  height: 35px;
  border: solid 1px #111;
  background-color: #c2c5c1;
  border-radius: 10px;
  margin: 5px;
  transition: 0.2s;
}

.counter_btn:hover {
  opacity: 0.8;
  cursor: pointer;
}

.counter_number {
  background-color: #686e65;
  text-align: center;
  width: 90%;
  font-size: 30px;
  border: 1px solid #333;
  border-radius: 10px;
}

.basket_item_price_block {
  display: flex;
  height: 150px;
  flex-direction: column;
  justify-content: space-around;
}

.basket_item_price {
  font-size: 30px;
  text-align: center;
  background-color: #686e65;
  width: 100%;
  border-radius: 10px;
  border: 1px solid #333;
}

.basket_item_remove_btn {
  font-size: 24px;
  border-radius: 5px;
  background-color: #df4545;
  border: solid 1px #111;
  cursor: pointer;
  transition: 0.2s;
}

.basket_item_remove_btn:hover {
  background-color: #df2727;
}
</style>
