<template>
  <div class="body" id="app">
    <HeaderItem></HeaderItem>
    <router-view />
  </div>
</template>

<style lang="scss">
body {
  margin: 0;
}
.intro {
  display: flex;
  flex-direction: column;
  width: 100%;
  min-height: 100vh;
  background-color: #333;
}
.container {
  max-width: 1200px;
  margin: 0 auto;
}
.title {
  display: flex;
  justify-content: center;
  margin: 10px auto;
}
.title_logo {
}
.title_text {
  font: bold 32px sans-serif;
  color: #d3c7c7;
}
button {
  font-size: 16px;
  text-transform: uppercase;
  border-radius: 5px;
  width: fit-content;
  cursor: pointer;
  border: solid 1px #111;
  transition: all 0.2s;
}
button:hover {
  cursor: pointer;
  transform: scale(1.05);
}
button:active {
  opacity: 0.9;
  transform: scale(1);
}
.green {
  background-color: green;
}
.red {
  background-color: red;
}
img {
  display: block;
  border-radius: 10px;
  max-width: 100%;
  max-height: 100%;
  transition: transform 0.2s;
}
</style>
<script>
import HeaderItem from "@/components/HeaderItem";
export default {
  components: { HeaderItem },
};
</script>
