<template>
  <header class="header">
    <div class="container">
      <div class="header_inner">
        <div class="header_title" @click="$router.push({ path: `/products` })">
          Доставка
        </div>
        <nav class="nav">
          <div
            class="nav_link_active"
            @click="$router.push({ path: `/products` })"
          >
            Главная
          </div>
          <div class="nav_link" @click="$router.push({ path: `/basket` })">
            Корзина
          </div>
          <div class="nav_link" @click="$router.push({ path: `/login` })">
            Войти
          </div>
        </nav>
      </div>
    </div>
  </header>
</template>

<script>
export default {
  name: "HeaderItem",
};
</script>

<style scoped>
.header {
  background-color: #71bf4d;
  font: bold large sans-serif;
}
.header_logo {
  font-size: 30px;
  text-transform: uppercase;
}
.header_inner {
  display: flex;
  font-size: 30px;
  justify-content: space-between;
  align-items: center;
}
.container {
  width: 100%;
  max-width: 1200px;
  margin: 0 auto;
}

.nav {
  font-size: 16px;
  text-transform: uppercase;
}

.nav_link {
  display: inline-block;
  vertical-align: top;
  padding: 10px;
  color: #333;
  text-decoration: none;
  transition: color 0.2s linear;
}

.nav_link:hover {
  color: #c8dcf4;
  cursor: pointer;
  background-color: #8f968b;
}

.nav_link_active {
  display: inline-block;
  vertical-align: top;
  text-decoration: none;
  padding: 10px;
  color: rgb(185, 77, 14);
  transition: color 0.2s linear;
  background-color: #8f968b;
}
.nav_link_active:hover {
  color: #c8dcf4;
  cursor: pointer;
  background-color: #8f968b;
}
</style>
