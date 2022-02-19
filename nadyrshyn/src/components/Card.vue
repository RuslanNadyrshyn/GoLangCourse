<template>
  <div>
    <div
      :title="'title' + $router.getRoutes()[0].name"
      :class="{ example: true, clickable: click }"
      :style="{ 'font-size': '40px' }"
      @click="increment"
    >
      {{ "EXAMPLE" }} Clickable: {{ click }}
    </div>
  </div>
</template>

<script>
export default {
  name: "CardItem",
  props: {
    title: {
      type: String,
      required: true,
    },
    cardValue: {
      type: Number,
      default: 0,
    },
    cardData: {
      type: Object,
    },
    click: {
      type: Boolean,
      default: true,
    },
  },
  data() {
    return {
      name: "cardData",
      data: { a: 1, b: 2 },
      clickable: false,
      v: 0,
    };
  },
  methods: {
    increment() {
      this.$emit("calc");
    },
  },
};
</script>

<style scoped>
.example {
  color: brown;
  font-size: 50px;
}

.clickable {
  cursor: pointer;
}
</style>
