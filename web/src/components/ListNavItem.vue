<template>
  <div class="list_nav_item">
    <template v-if="isProduct === false">
      <div
        v-bind:class="
          type === this.$store.state.suppliers.selectedType
            ? 'list_nav active'
            : 'list_nav'
        "
        @click="sortType()"
      >
        {{ type }}
      </div>
      <label
        v-bind:class="
          type === this.$store.state.suppliers.selectedType
            ? 'counter'
            : 'counter hidden'
        "
        @click="sortType()"
      >
        {{ $store.state.suppliers.suppliers.length }}
      </label>
    </template>

    <template v-else-if="isProduct === true">
      <div
        v-bind:class="
          type === this.$store.state.products.selectedType
            ? 'list_nav active'
            : 'list_nav'
        "
        @click="sortType()"
      >
        {{ type }}
      </div>
      <label
        v-bind:class="
          type === this.$store.state.products.selectedType
            ? 'counter'
            : 'counter hidden'
        "
      >
        {{ this.$store.state.products.products.length }}
      </label>
    </template>
  </div>
</template>

<script>
export default {
  name: "ListNavItem",
  props: {
    type: {
      type: String,
    },
    isProduct: {
      type: Boolean,
    },
  },
  methods: {
    sortType() {
      if (this.isProduct) {
        if (this.type !== this.$store.getters["products/getSelectedType"]) {
          let params = {
            supId: this.$store.getters["suppliers/getSelectedSupplier"],
            supType: this.$store.getters["suppliers/getSelectedType"],
            prodType: this.type,
          };
          this.$store.dispatch("products/getByParams", params);
        }
      } else {
        let params = {
          supId: 0,
          supType: this.type,
          prodType: "все",
        };
        this.$store.dispatch("suppliers/getByType", this.type);
        this.$store.dispatch("products/getByParams", params);
      }
    },
  },
};
</script>

<style scoped>
.list_nav_item {
  position: relative;
}

.counter {
  position: absolute;
  color: #111;
  border-radius: 5px;
  right: 2px;
  top: 0;
  font-size: 12px;
}
.counter:hover {
  cursor: pointer;
}
.counter.hidden {
  display: none;
}

.list_nav {
  padding: 10px;
  color: #d3c7c7;
  text-decoration: none;
  text-transform: uppercase;
  transition: color 0.2s linear;
}

.list_nav:hover {
  background-color: rgb(245, 208, 195);
  color: #333;
  cursor: pointer;
}
.list_nav.active {
  background-color: #8f968b;
  color: #333;
}
.list_nav.active:hover {
  background-color: rgb(245, 208, 195);
  color: #333;
  cursor: pointer;
}

@media (max-width: 560px) {
  .counter {
    right: 0;
    top: -4px;
    font-size: 9px;
    padding: 2px;
    background: none;
    border: none;
  }
  .list_nav {
    font-size: 16px;
    padding: 7px 10px;
  }
}
</style>
