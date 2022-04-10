<template>
  <div class="body" id="app">
    <nav></nav>
    <router-view />
  </div>
</template>

<style lang="scss">
body {
  margin: 0;
}
button {
  font-size: 16px;
  text-transform: uppercase;
  border-radius: 5px;
  width: fit-content;
  cursor: pointer;
  border: solid 1px #111;
  opacity: 0.9;
  transition: color 0.2s linear, transform 0.1s linear, scale 0.2s;
}
button:hover {
  opacity: 1;
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
</style>
