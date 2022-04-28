<template>
  <div class="intro">
    <template v-if="$store.state.orders.loaded">
      <template v-if="$store.state.orders.errors.length">
        <div v-for="(error, index) in $store.state.orders.errors" :key="index">
          {{ error }}
        </div>
      </template>
      <div class="section">
        <div class="container">
          <div class="title">
            <label class="title_text">
              Заказ №{{ $store.state.orders.order.id }}
            </label>
          </div>
          <div class="basket_block">
            <div
              v-for="product in this.$store.getters['orders/getOrderProducts']"
              :key="product.id"
            >
              <OrderItem :product="product"></OrderItem>
            </div>
          </div>
          <div class="total_price_block">
            <div class="total_price_container">
              Всего:
              <div class="total_price">
                {{ $store.state.orders.order.price }}
              </div>
              грн
            </div>
          </div>
        </div>
      </div>
    </template>
  </div>
</template>

<script>
import OrderItem from "@/components/OrderItem";

export default {
  name: "LoadedOrderView",
  components: { OrderItem },
  data() {
    return {
      loaded: false,
      errors: [],
    };
  },
  created() {
    this.$store.dispatch("orders/fetchOrder", this.$route.params.id);
  },
};
</script>

<style scoped></style>
