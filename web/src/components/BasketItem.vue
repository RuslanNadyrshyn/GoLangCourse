<template>
  <div class="basket_item">
    <div class="bi_image">
      <img
        :src="product.image"
        alt="product"
        @click="$router.push({ path: `/product/${product.id}` })"
      />
    </div>
    <div class="description_container">
      <div class="description_block">
        <label
          class="title_product"
          @click="$router.push({ path: `/product/${product.id}` })"
        >
          {{ product.name }}
        </label>
        <div class="ingredients_block">
          <a class="bi_image hidden">
            <img
              :src="product.image"
              alt="product"
              @click="$router.push({ path: `/product/${product.id}` })"
            />
          </a>
          <div class="bi_ingredients">
            <div
              v-for="ingredient in product.ingredients"
              :key="ingredient.id"
              class="bi_ingredient"
            >
              <label>{{ ingredient }}</label>
            </div>
          </div>
        </div>
      </div>
      <div class="price_container">
        <div class="price_item">
          <div class="price">{{ summary.toFixed(2) }}</div>
          <label>грн</label>
        </div>
        <div style="display: flex; margin-top: 5px">
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
          <div class="price_block">
            <div class="price_item">
              <div class="price">{{ product.price }}</div>
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
      if (newValue > 0 && Number.isInteger(newValue)) {
        let basket = this.$store.getters["basket/getBasket"];
        for (let i = 0; i < basket.length; i++)
          if (basket[i].id === this.product.id) {
            basket[i].counter = newValue;
            this.summary = this.product.price * this.product.counter;
            this.$store.dispatch("basket/updateProduct", basket[i]);
            this.$store.dispatch("basket/calcTotalPrice");
          }
      } else if (newValue !== "") {
        if (!Number.isInteger(newValue)) {
          this.counter = Math.floor(newValue);
        } else this.counter = 1;
      } else if (newValue === "") this.summary = this.product.price;
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
  margin: 5px;
  background-color: #4a4e47;
  border-radius: 10px;
}

.bi_image {
  display: flex;
  align-items: center;
  justify-content: center;
  margin: auto 5px;
  width: 20%;
  height: 130px;
}
.bi_image:hover {
  cursor: pointer;
  transform: scale(1.03);
}
.bi_image.hidden {
  display: none;
}

.description_container {
  display: flex;
  flex-direction: row;
  width: 80%;
  justify-content: space-around;
  margin: 5px;
}

.description_block {
  display: flex;
  flex-direction: column;
  width: 80%;
  text-align: center;
}
.title_product {
  font-size: 30px;
  margin: auto;
}
.title_product:hover {
  cursor: pointer;
}
.ingredients_block {
  display: flex;
  flex-direction: row;
  height: 100px;
}
.bi_ingredients {
  display: flex;
  flex-direction: column;
  background-color: #686e65;
  border: 1px solid #333;
  border-radius: 10px;
  overflow: scroll;
  scroll-behavior: smooth;
}
.bi_ingredient {
  font-size: 18px;
}

.price_container {
  display: flex;
  flex-direction: column;
  justify-content: space-evenly;
  align-items: center;
}

.price_block {
  display: flex;
  flex-direction: column;
  justify-content: space-between;
  align-items: end;
  margin-left: 5px;
}
.price_item {
  display: flex;
  flex-direction: row;
  justify-content: end;
  align-items: center;
}
.price {
  display: flex;
  justify-content: center;
  align-items: center;
  font-size: 30px;
  min-width: 75px;
  width: max-content;
  max-width: 155px;
  overflow: hidden;
  background-color: #686e65;
  margin-right: 5px;
  padding: 0 5px;
  border-radius: 5px;
  border: 1px solid #333;
}

.bi_ingredients {
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
.bi_ingredient {
  font-size: 18px;
}

@media (max-width: 810px) {
  .basket_item {
    font-size: 16px;
    min-height: 100px;
  }
  .bi_image {
    width: 20%;
    max-height: 90px;
  }

  .price {
    font-size: 20px;
    min-width: 50px;
  }

  .title_product {
    font-size: 26px;
  }

  .ingredients_block {
    height: 64px;
  }
  .bi_ingredient {
    font-size: 14px;
  }
}

@media (max-width: 560px) {
  .bi_image {
    display: none;
  }
  .bi_image.hidden {
    max-height: 64px;
    display: flex;
    width: 30%;
  }
  .bi_ingredients {
    width: 70%;
  }
  .title_product {
    font-size: 20px;
  }
  .description_container {
    width: 100%;
  }
  .price_item {
    font-size: 14px;
  }
}
</style>

<style scoped>
.counter_block {
  display: flex;
  flex-direction: column;
  justify-content: center;
}
.counter_btn {
  display: flex;
  justify-content: center;
  align-items: center;
  font-size: 20px;
  width: 77px;
  border-radius: 0;
}

.counter_number {
  font: large medium sans-serif;
  font-size: 22px;
  padding: 2px 0;
  margin: -1px 0;
  background-color: #686e65;
  width: 75px;
  border: 1px solid #222;
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
