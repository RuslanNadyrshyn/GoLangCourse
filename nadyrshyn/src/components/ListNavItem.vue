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
        this.$store.dispatch("products/sortByType", this.type);
      } else if (this.isSupplier === true) {
        console.log("You pressed supplier type: ", this.type);
        if (this.type === "Открыто") {
          this.$store.dispatch("suppliers/sortByWorkingHours");
        } else {
          this.$store.dispatch("suppliers/sortByType", this.type);
        }
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
  border-radius: 10px 10px 0 0;
  text-transform: uppercase;
  transition: color 0.2s linear;
}

.list_nav_item:hover {
  background-color: rgb(245, 208, 195);
  color: #333;
  cursor: pointer;
}
</style>
