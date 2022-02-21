<template>
  <div class="Supplier_item">
    <div class="Supplier_img">
      <img class="Supplier_logo" :src="image" alt="" />
    </div>
    <div class="Supplier_text">
      {{ name }}
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
      type: Array,
    },
  },
};
</script>

<style scoped>
.Supplier_item {
  display: flex;
  flex-direction: column;
  color: #222;
  text-align: center;
  max-width: 300px;
  height: 150px;
  margin: 10px;
}
.Supplier_item:hover {
  opacity: 0.9;
  cursor: pointer;
}
.Supplier_img {
  display: flex;
  max-width: 100%;
  height: 250px;
  align-items: center;
}

.Supplier_logo {
  display: block;
  border-radius: 20%;
  max-width: 100px;
  margin: auto;
  transition: transform 0.2s;
}

.Supplier_logo:hover {
  transform: scale(1.05);
}
.Supplier_text {
  font-size: 20px;
  font-weight: 700;
  text-transform: uppercase;
}
</style>
