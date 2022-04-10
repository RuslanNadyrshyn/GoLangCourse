<template>
  <div class="list_nav_item" @click="sortType()">
    {{ type }}
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
    isSupplier: {
      type: Boolean,
    },
  },
  methods: {
    sortType() {
      if (this.isProduct === true) {
        let sortedSuppliers =
          this.$store.getters["suppliers/getSortedSuppliers"];
        let sortedProducts = [];
        for (let i = 0; i < sortedSuppliers.length; i++)
          for (let j = 0; j < sortedSuppliers[i].menu.length; j++)
            sortedProducts.push(sortedSuppliers[i].menu[j]);
        this.$store.dispatch("products/sortBySupplier", sortedProducts);
        this.$store.dispatch("products/sortByType", this.type);
      } else if (this.isSupplier === true) {
        console.log("You pressed supplier type: ", this.type);
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
          this.$store.dispatch("products/sortBySupplier", prod);
        });
      }
    },
  },
};
</script>

<style scoped>
.list_nav_item {
  padding: 10px;
  color: #d3c7c7;
  text-decoration: none;
  text-transform: uppercase;
  transition: color 0.2s linear;
}

.list_nav_item:hover {
  background-color: rgb(245, 208, 195);
  color: #333;
  cursor: pointer;
}
</style>
