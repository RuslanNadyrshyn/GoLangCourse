<template>
  <header class="header">
    <div class="container">
      <div class="header_inner">
        <div class="header_title">Доставка</div>
        <nav class="nav">
          <a class="nav_link" href="/HomeView"> Главная </a>
          <a class="nav_link_active" href="./product-list.html">
            Список товаров
          </a>
          <a class="nav_link" href="./bin.html"> Корзина </a>
          <a class="nav_link" href="./login.html"> Войти </a>
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
}
.header_logo {
  font-family: "Arial Black";
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
  margin: 0 15px;
  color: #333;
  text-decoration: none;
  transition: color 0.2s linear;
}

.nav_link:hover {
  color: #c8dcf4;
}

.nav_link_active {
  display: inline-block;
  vertical-align: top;
  margin: 0 15px;
  color: rgb(185, 77, 14);
  transition: color 0.2s linear;
}

.nav_link_active:hover {
  color: #c8dcf4;
}
</style>
