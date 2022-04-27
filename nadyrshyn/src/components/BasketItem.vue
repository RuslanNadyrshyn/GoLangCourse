<template>
  <div class="basket_item">
    <div class="basket_item_image">
      <img
        :src="product.image"
        alt="product"
        @click="$router.push({ path: `/product/${product.id}` })"
      />
    </div>
    <div class="description_block">
      <div class="basket_title">
        <label
          class="title_product"
          @click="$router.push({ path: `/product/${product.id}` })"
        >
          {{ product.name }}
        </label>
        <div class="price_block">
          <div class="price">{{ product.price }}</div>
          <label>грн</label>
        </div>
      </div>
      <div class="basket_container">
        <div class="container_block">
          <div class="basket_item_image hidden">
            <img
              :src="product.image"
              alt="product"
              @click="$router.push({ path: `/product/${product.id}` })"
            />
          </div>
          <div class="basket_item_ingredients">
            <div
              v-for="ingredient in product.ingredients"
              :key="ingredient.id"
              class="basket_item_ingredient"
            >
              <label>{{ ingredient }}</label>
            </div>
          </div>
        </div>
        <div class="block">
          <div class="counter_block">
            <button class="counter_btn green" v-on:click="counter++">
              &uarr;
            </button>
            <input
              v-model.number="counter"
              class="counter_number"
              type="number"
              step="1"
              min="1"
              required
            />
            <button class="counter_btn red" v-on:click="counter--">
              &darr;
            </button>
          </div>
          <div class="price_container">
            <div class="price_block">
              <div class="price">{{ summary.toFixed(2) }}</div>
              <label>грн</label>
            </div>
            <button class="basket_delete_btn red" v-on:click="deleteProd()">
              Удалить
            </button>
          </div>
        </div>
      </div>
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
      else if (newValue === "") this.summary = this.product.price;
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
  display: flex;
  flex-direction: row;
  min-height: 150px;
  margin: 10px;
  background-color: #4a4e47;
  border-radius: 10px;
}

.basket_item_image {
  display: flex;
  align-items: center;
  justify-content: center;
  margin: auto 5px;
  width: 20%;
  height: 130px;
}
.basket_item_image:hover {
  cursor: pointer;
  transform: scale(1.03);
}
.basket_item_image.hidden {
  display: none;
}

.description_block {
  display: flex;
  flex-direction: column;
  width: 80%;
  justify-content: space-around;
}

.basket_title {
  display: flex;
  margin: 5px 5px 0;
  justify-content: space-between;
}

.price_block {
  display: flex;
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

.title_product {
  font-size: 30px;
  margin: auto;
}
.title_product:hover {
  cursor: pointer;
}

.basket_container {
  display: flex;
  margin: 5px;
  height: 100px;
}
.container_block {
  display: flex;
  flex-direction: row;
  width: 80%;
}
.basket_item_ingredients {
  display: flex;
  flex-direction: column;
  width: 100%;
  background-color: #686e65;
  border: 1px solid #333;
  border-radius: 10px;
  overflow-y: scroll;
  scroll-behavior: smooth;
  margin-right: 5px;
}
.basket_item_ingredient {
  font-size: 18px;
}

.block {
  display: flex;
  flex-direction: row;
  justify-content: space-evenly;
  align-items: center;
}

@media (max-width: 810px) {
  .basket_item {
    font-size: 16px;
    min-height: 100px;
  }
  .basket_item_image {
    width: 20%;
    max-height: 90px;
  }

  .price {
    font-size: 20px;
    min-width: 50px;
    padding: 2px;
  }

  .title_product {
    font-size: 26px;
  }

  .basket_container {
    height: 64px;
  }
  .basket_item_ingredient {
    font-size: 14px;
  }
}

@media (max-width: 560px) {
  .basket_item_image {
    display: none;
  }
  .basket_item_image.hidden {
    max-height: 64px;
    display: flex;
    width: 30%;
  }
  .basket_item_ingredients {
    width: 70%;
  }
  .title_product {
    font-size: 20px;
  }
  .description_block {
    width: 100%;
  }
}
</style>

<style scoped>
.counter_block {
  display: flex;
  flex-direction: column;
  align-items: center;
  margin-right: 5px;
}
.counter_btn {
  display: flex;
  justify-content: center;
  align-items: center;
  font-size: 22px;
  width: 77px;
  border-radius: 0;
}

.counter_number {
  font: large medium sans-serif;
  font-size: 24px;
  padding: 5px 0;
  margin: -1px 0;
  background-color: #686e65;
  width: 75px;
  border: 1px solid #222;
}
.price_container {
  display: flex;
  flex-direction: column;
  justify-content: space-between;
  align-items: center;
  height: 100%;
}
.basket_delete_btn {
  width: 100%;
  padding: 10px;
}
@media (max-width: 810px) {
  .counter_number {
    font-size: 16px;
    width: 40px;
    height: 15px;
  }
  .counter_btn {
    font-size: 14px;
    width: 42px;
    height: 20px;
  }
  .basket_delete_btn {
    font-size: 14px;
    width: 100%;
    padding: 5px;
  }
}
</style>
