<template>
  <header>
    <div class="container">
      <div class="header_inner">
        <router-link class="header_title" :to="'/'">Доставка</router-link>
        <nav class="nav">
          <router-link
            v-bind:class="$route.path === '/' ? 'nav_link_active' : 'nav_link'"
            :to="'/'"
          >
            Главная
          </router-link>
          <router-link
            v-bind:class="
              $route.path === '/basket' ? 'nav_link_active' : 'nav_link'
            "
            :to="'/basket'"
          >
            Корзина
            <label
              v-bind:class="
                this.$store.state.basket.products.length > 0
                  ? 'counter'
                  : 'counter hidden'
              "
            >
              {{ this.$store.state.basket.products.length }}
            </label>
          </router-link>
          <router-link
            v-bind:class="
              $route.path === '/login' ? 'nav_link_active' : 'nav_link'
            "
            :to="'/login'"
          >
            Войти
          </router-link>
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
header {
  width: 100%;
  background-color: #4b4242ff;
  font: bold large sans-serif;
}
.header_logo {
  font-size: 30px;
  text-transform: uppercase;
}
.header_title {
  display: flex;
  justify-content: space-between;
  cursor: pointer;
  color: #d3c7c7;
}

.header_inner {
  display: flex;
  font-size: 30px;
  justify-content: space-between;
  align-items: center;
}

.nav {
  display: flex;
  flex-direction: row;
  font-size: 20px;
  text-transform: uppercase;
  text-decoration: none;
}

.nav_link {
  position: relative;
  display: inline-block;
  vertical-align: top;
  padding: 10px;
  color: #d3c7c7;
  text-decoration: none;
  transition: 0.2s linear;
}

.nav_link:hover {
  transform: scale(0.95);
  color: #333;
  cursor: pointer;
  background-color: #8f968b;
}

.nav_link_active {
  position: relative;
  display: inline-block;
  vertical-align: top;
  text-decoration: none;
  padding: 10px;
  color: #dc410f;
  transition: color 0.2s linear;
  background-color: #8f968b;
}

.nav_link_active:hover {
  color: #c8dcf4;
  cursor: pointer;
  background-color: #8f968b;
}

.counter {
  position: absolute;
  background-color: #dc410f;
  color: #111;
  border-radius: 5px;
  padding: 2px 4px;
  right: 3px;
  bottom: -15px;
  font-size: 14px;
}
.counter.hidden {
  opacity: 0;
}
</style>
