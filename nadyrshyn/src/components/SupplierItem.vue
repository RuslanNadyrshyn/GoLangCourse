<template>
  <div class="supplier_item" @click="sortBySupplier()">
    <div class="supplier_img">
      <img class="supplier_logo" :src="image" alt="" />
    </div>
    <div class="supplier_text">
      {{ name }}
    </div>
    <div class="working_hours">
      {{ workingHours.opening }}-{{ workingHours.closing }}
    </div>
  </div>
</template>

<script>
export default {
  name: "SupplierItem",
  props: {
    id: {
      type: Number,
    },
    name: {
      type: String,
    },
    image: {
      type: String,
    },
    type: {
      type: String,
    },
    workingHours: {
      type: Object,
    },
    menu: {
      type: Array,
    },
  },
  methods: {
    sortBySupplier() {
      this.$store.dispatch("products/sortBySupplier", this.menu);
    },
  },
};
</script>

<style scoped>
.supplier_item {
  display: flex;
  background-color: #8f968b;
  flex-direction: column;
  color: #222;
  text-align: center;
  max-width: 200px;
  height: 180px;
  margin: 10px;
  padding: 10px;
  border-radius: 20px;
}

.supplier_item:hover {
  opacity: 0.9;
  cursor: pointer;
}

.supplier_img {
  display: flex;
  width: 100%;
  height: 250px;
  align-items: center;
}

.supplier_logo {
  display: block;
  border-radius: 20px;
  max-width: 100px;
  margin: auto;
  transition: transform 0.2s;
}

.supplier_logo:hover {
  transform: scale(1.05);
}

.supplier_text {
  display: flex;
  flex-direction: column;
  justify-content: center;
  font-size: 20px;
  min-height: 50px;
  font-weight: 700;
  text-transform: uppercase;
}

.working_hours {
  font-size: 16px;
}
</style>
