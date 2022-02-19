<template>
  <div class="home">
    <example
      :title="`Title`"
      :click="true"
      :exampleValue="213"
      :exampleData="{ a: 5, b: 23 }"
    >
      title
    </example>
    <img alt="Vue logo" src="../assets/logo.png" />
    <HelloWorld msg="Welcome to Your Vue.js App" />
  </div>
</template>

<script>
// @ is an alias to /src
import HelloWorld from "@/components/HelloWorld.vue";
import Example from "@/views/example";

export default {
  name: "HomeView",
  components: {
    Example,
    HelloWorld,
  },
};
</script>
