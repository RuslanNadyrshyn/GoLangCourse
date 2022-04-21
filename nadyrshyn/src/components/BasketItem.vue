<template>
  <div class="basket_item">
    <div class="basket_item_image">
      <img
        :src="product.image"
        alt="product"
        @click="$router.push({ path: `/product/${product.id}` })"
      />
    </div>
    <div class="price_block">
      <label class="price_block_title">Цена:</label>
      <div class="price_container">
        <div class="price">{{ product.price }}</div>
        <label>грн</label>
      </div>
    </div>
    <div class="description_block">
      <label class="basket_item_title">
        {{ product.name }}
      </label>
      <div class="basket_item_description">
        <div class="basket_item_ingredients">
          <div v-for="ingredient in product.ingredients" :key="ingredient.id">
            <label class="basket_item_ingredient">{{ ingredient }}</label>
          </div>
        </div>
      </div>
    </div>
    <div class="counter_block">
      <button class="counter_btn green" v-on:click="counter++">&uarr;</button>
      <input
        v-model.number="counter"
        class="counter_number"
        type="number"
        step="1"
        min="1"
        required
      />
      <button class="counter_btn red" v-on:click="counter--">&darr;</button>
    </div>
    <div class="price_block total">
      <label class="price_block_title">Общая цена:</label>
      <div class="price_container">
        <div class="price">{{ summary.toFixed(2) }}</div>
        <label>грн</label>
      </div>
      <button class="basket_delete_btn red" v-on:click="deleteProd()">
        Удалить
      </button>
    </div>
  </div>
</template>

<script>
export default {
  name: "BasketItem",
  data() {
    return {
      counter: this.product.counter,
      summary: this.product.price * this.product.counter,
    };
  },
  props: {
    product: {
      type: Object,
    },
  },
  watch: {
    counter(newValue) {
      if (newValue > 0) {
        let basket = this.$store.getters["basket/getBasket"];
        for (let i = 0; i < basket.length; i++)
          if (basket[i].id === this.product.id) {
            basket[i].counter = newValue;
            this.$store.dispatch("basket/updateProduct", basket[i]);
            this.summary = this.product.price * this.product.counter;
            this.$store.dispatch("basket/calcTotalPrice");
          }
      } else if (newValue !== "") this.counter = 1;
    },
  },
  methods: {
    deleteProd() {
      this.counter = 0;
      this.$store.dispatch("basket/deleteProduct", this.product.id);
    },
  },
};
</script>

<style>
input[type="number"] {
  text-align: center;
  -moz-appearance: textfield;
  -moz-box-sizing: content-box;
}
.basket_item {
  font: large medium sans-serif;
  display: flex;
  flex-direction: row;
  justify-content: space-evenly;
  align-items: center;
  min-height: 150px;
  margin: 10px;
  background-color: #4a4e47;
  border-radius: 10px;
}

.basket_item_image {
  display: flex;
  align-items: center;
  justify-content: center;
  margin: 0 10px;
  width: 170px;
  height: 130px;
}
.basket_item_image:hover {
  cursor: pointer;
  transform: scale(1.03);
}

.price_block {
  height: 90px;
  display: flex;
  margin: 5px;
  flex-direction: column;
  justify-content: space-around;
}
.price_block.total {
  height: 130px;
}

.price_block_title {
  font-size: 16px;
}

.price_container {
  display: flex;
  text-transform: uppercase;
  flex-direction: row;
  justify-content: center;
  align-items: center;
}

.price {
  display: flex;
  justify-content: center;
  align-items: center;
  font-size: 30px;
  min-width: 80px;
  max-width: max-content;
  background-color: #686e65;
  padding: 5px;
  margin-right: 5px;
  border-radius: 5px;
  border: 1px solid #333;
}

.description_block {
  display: flex;
  flex-direction: column;
  min-height: 130px;
  width: 450px;
  margin: 0 5px;
}

.basket_item_description {
  display: flex;
  flex-direction: column;
  min-height: 68px;
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
  padding: 10px 0;
}

.basket_item_ingredient {
  margin: 10px;
}

.counter_block {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  width: 100px;
  height: 150px;
}
</style>

<style scoped>
.counter_btn {
  display: flex;
  justify-content: center;
  text-align: center;
  font-size: 22px;
  width: 77px;
  border-radius: 0;
}

.counter_number {
  font: large medium sans-serif;
  font-size: 30px;
  padding: 5px 0;
  margin: -1px 5px;
  background-color: #686e65;
  width: 75px;
  height: 40px;
  border: 1px solid #222;
}

.basket_delete_btn {
  width: 120px;
  padding: 10px;
  margin: 0 auto;
}
</style>
