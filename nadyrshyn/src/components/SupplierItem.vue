<template>
  <div class="supplier_item" @click="sortBySupplier()">
    <div class="supplier_img">
      <img class="supplier_logo" :src="supplier.image" alt="supplier_logo" />
    </div>
    <div class="supplier_text">
      {{ supplier.name }}
    </div>
    <div class="working_hours">
      {{ supplier.workingHours.opening }}-{{ supplier.workingHours.closing }}
    </div>
  </div>
</template>

<script>
export default {
  name: "SupplierItem",
  props: {
    supplier: {
      type: Object,
    },
  },
  methods: {
    sortBySupplier() {
      this.$store.dispatch("products/sortBySupplier", this.supplier.menu);
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
  width: 170px;
  height: 180px;
  margin: 5px;
  padding: 10px 5px;
  border-radius: 10px;
}

.supplier_item:hover {
  opacity: 0.9;
  cursor: pointer;
}

.supplier_img {
  display: flex;
  width: 100%;
  height: 110px;
}

.supplier_logo {
  display: block;
  border-radius: 20px;
  max-width: 100%;
  max-height: 100px;
  margin: auto;
  transition: transform 0.2s;
}

.supplier_logo:hover {
  transform: scale(1.05);
}

.supplier_text {
  display: flex;
  flex-direction: column;
  text-align: center;
  justify-content: center;
  font-size: 20px;
  min-height: 50px;
  font-weight: 700;
  text-transform: uppercase;
}

.working_hours {
  font-size: 16px;
}

@media (max-width: 560px) {
  .supplier_item {
    width: 110px;
    height: 120px;
  }
  .supplier_img {
    height: 50px;
  }
  .supplier_logo {
    max-height: 50px;
  }
  .supplier_text {
    font-size: 14px;
  }
}
</style>
