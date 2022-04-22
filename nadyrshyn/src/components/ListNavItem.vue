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
      >
        {{ $store.state.suppliers.sortedSuppliers.length }}
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
        {{ $store.state.products.sortedProducts.length }}
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
      if (this.isProduct === false) {
        let suppliers;
        if (this.type === "Открыто") {
          suppliers = this.$store.dispatch("suppliers/sortByWorkingHours");
        } else {
          suppliers = this.$store.dispatch("suppliers/sortByType", this.type);
        }
        suppliers.then((res) => {
          let prod = [];
          for (let i = 0; i < res.length; i++)
            for (let j = 0; j < res[i].menu.length; j++)
              prod.push(res[i].menu[j]);
          this.$store.dispatch("products/sortBySupplier", prod).then(() => {
            this.$store.dispatch(
              "products/sortByType",
              this.$store.state.products.selectedType
            );
          });
        });
      } else {
        let sortedSuppliers =
          this.$store.getters["suppliers/getSortedSuppliers"];
        let sortedProducts = [];
        for (let i = 0; i < sortedSuppliers.length; i++)
          for (let j = 0; j < sortedSuppliers[i].menu.length; j++)
            sortedProducts.push(sortedSuppliers[i].menu[j]);
        this.$store.dispatch("products/sortBySupplier", sortedProducts);
        this.$store.dispatch("products/sortByType", this.type);
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
  background-color: #d3c7c7;
  color: #111;
  border: solid #111 1px;
  border-radius: 5px;
  padding: 2px 4px;
  right: 2px;
  top: -12px;
  font-size: 14px;
}
.counter.hidden {
  opacity: 0;
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
</style>
