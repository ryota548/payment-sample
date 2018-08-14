<template>
<div class="hello">
  <ul>
    <li v-for="item in items" :key="item.ID" @click="pageto(item.ID)">
      <item-card :item="item"></item-card>
    </li>
  </ul>
</div>
</template>

<script>
import axios from 'axios'
import ItemCard from './ItemCard'
export default {
  name: 'Home',
  components: {
    ItemCard
  },
  data () {
    return {
      items: []
    }
  },
  methods: {
    // ページ移動
    pageto: function (id) {
      this.$router.push(`/items/${id}`)
    }
  },
  // 商品リストをすべてとってくる
  created () {
    axios.get('http://localhost:8888/api/v1/items').then(res => {
      this.items = res.data
    })
  }
}
</script>
<style scoped>
h1, h2 {
  font-weight: normal;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
</style>
