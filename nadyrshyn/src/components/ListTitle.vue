<template>
  <div class="list_title_item">
    <div class="list_title_img">
      <img class="list_logo" :src="image" , alt="" />
    </div>
    <div class="list_title_text">{{ text }}</div>
  </div>
</template>

<script>
export default {
  name: "ListTitle",
  props: {
    image: {
      type: String,
    },
    text: {
      type: String,
    },
  },
};
</script>

<style scoped>
.list_title_item {
  display: flex;
  flex-direction: row;
  align-items: center;
  justify-content: left;
  color: #222;
  text-align: center;
  padding: 5px 5px 5px 0;
}

.list_title_img {
  display: block;
  width: 100px;
  margin-right: 10px;
}
.list_logo {
  display: block;
  border-radius: 20%;
  max-width: 100%;
  transition: transform 0.2s;
}
.list_title_text {
  width: max-content;
  font-size: 42px;
}
</style>
