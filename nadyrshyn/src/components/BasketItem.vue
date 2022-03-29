<template>
  <div class="basket_item">
    <div class="basket_item_image">
      <img class="product_logo" :src="image" alt="" />
    </div>
    <div class="description_block">
      <div class="basket_item_title">
        {{ name }}
      </div>
      <div class="basket_item_description">
        <div class="basket_item_ingredients">
          <div v-for="ingredient in ingredients" :key="ingredient.id">
            <div class="basket_item_ingredient">{{ ingredient }}</div>
          </div>
        </div>
      </div>
    </div>
    <div class="counter_container">
      <div class="basket_item_counter">
        <button class="counter_btn green" v-on:click="increment">+</button>
        <div class="counter_number">{{ counter }}</div>
        <button class="counter_btn red" v-on:click="decrement">-</button>
      </div>
      <div class="basket_item_price_block">
        <div class="basket_item_price_container">
          <div class="basket_item_price">{{ price }}</div>
          грн
        </div>
        <button class="basket_item_remove_btn" v-on:click="deleteProd()">
          Удалить
        </button>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "BasketItem",
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
    counter: {
      type: Number,
    },
  },
  methods: {
    increment() {
      this.$store.commit("basket/incCount", this.id);
      this.$store.dispatch("basket/calcTotalPrice");
    },
    decrement() {
      this.$store.commit("basket/decCount", this.id);
      this.$store.dispatch("basket/calcTotalPrice");
    },
    deleteProd() {
      this.$store.dispatch("basket/deleteProduct", this.id);
      this.$store.dispatch("basket/calcTotalPrice");
    },
  },
};
</script>

<style scoped>
.basket_item {
  display: flex;
  flex-direction: row;
  justify-content: space-between;
  align-items: center;
  min-height: 100px;
  margin: 5px auto;
  padding: 15px;
  background-color: #4a4e47;
  border-radius: 20px;
}

.basket_item_image {
  border-radius: 20px;
  cursor: pointer;
  width: 20%;
}

.product_logo {
  display: flex;
  max-width: 130px;
  min-height: 100px;
  background-color: #ddd;
  border-radius: 20px;
  margin: 10px auto;
}
.description_block {
  display: flex;
  flex-direction: column;
  height: 130px;
  width: 50%;
}

.basket_item_description {
  display: flex;
  flex-direction: column;
  min-height: 80px;
  background-color: #686e65;
  border: 1px solid #333;
  border-radius: 10px;
}

.basket_item_title {
  text-align: center;
  font-size: 24px;
  margin-bottom: 5px;
}

.basket_item_ingredients {
  font-size: 18px;
  display: flex;
  flex-wrap: wrap;
}

.basket_item_ingredient {
  margin: 0 10px;
}

.counter_container {
  display: flex;
  flex-direction: row;
  justify-content: space-between;
  width: 25%;
}

.basket_item_counter {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  width: 40%;
  height: 150px;
}

.counter_btn {
  display: flex;
  justify-content: center;
  font-size: 24px;
  width: 40px;
  height: 35px;
  border: solid 1px #111;
  background-color: #c2c5c1;
  border-radius: 10px;
  margin: 5px 0;
  transition: 0.2s;
}

.counter_btn.green {
  background-color: #2d8d0f;
}

.counter_btn.red {
  background-color: #df2727;
}

.counter_btn:hover {
  opacity: 0.8;
  cursor: pointer;
}

.counter_number {
  display: flex;
  justify-content: center;
  align-items: center;
  background-color: #686e65;
  min-width: 40px;
  height: 40px;
  font-size: 30px;
  border: 1px solid #333;
  border-radius: 10px;
}

.basket_item_price_block {
  display: flex;
  height: 150px;
  width: 60%;
  flex-direction: column;
  justify-content: space-evenly;
}

.basket_item_price_container {
  display: flex;
  flex-direction: row;
  justify-content: center;
  align-items: center;
}

.basket_item_price {
  display: flex;
  justify-content: center;
  align-items: center;
  font-size: 30px;
  min-width: 80px;
  background-color: #686e65;
  margin: 0 5px;
  padding: 5px;
  border-radius: 10px;
  border: 1px solid #333;
}

.basket_item_remove_btn {
  font-size: 20px;
  border-radius: 10px;
  background-color: #df4545;
  border: solid 1px #111;
  padding: 5px;
  cursor: pointer;
  transition: 0.2s;
}

.basket_item_remove_btn:hover {
  background-color: #df2727;
}
</style>
