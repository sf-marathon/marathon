<template>
  <div id="app">
    <yd-navbar
      slot="navbar"
      :title="title"
      color="#ffffff"
      bgcolor="#ff6532">
      <router-link to="/" slot="left">
        <yd-navbar-back-icon color="#ffffff" class="btn-nav-back" v-show="showBack"></yd-navbar-back-icon>
      </router-link>
    </yd-navbar>
    <router-view/>
  </div>
</template>

<script>
import uuidv4 from 'uuid/v4'
export default {
  name: 'app',
  beforeMount () {
    if (window.localStorage.openid) {
      this.$store.dispatch('SET_USER', { openid: window.localStorage.openid })
    } else {
      const openid = uuidv4()
      this.$store.dispatch('SET_USER', { openid })
      window.localStorage.openid = openid
    }
  },
  computed: {
    title () {
      return this.$store.state.title
    },
    showBack () {
      return this.$route.path !== '/'
    }
  }
}
</script>

<style>

</style>
