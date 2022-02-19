<template>
  <div class="example">
    Example with id: {{ $route.params.id }}
    <card
      :title="`Title`"
      :click="true"
      :cardValue="213"
      :cardData="{ a: 5, b: 23 }"
      @calc="increment"
    >
    </card>
    <div class="card">{{ v }}</div>
  </div>
</template>

<script>
import Card from "@/components/Card";
export default {
  name: "ExampleItem",
  components: { Card },
  data() {
    return {
      v: 0,
    };
  },
  methods: {
    increment() {
      this.v++;
    },
  },
};
</script>

<style scoped>
.example {
  color: darkgreen;
  font-size: 32px;
}
</style>
