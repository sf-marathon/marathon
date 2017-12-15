<template>
<div class="jh-share">
  <div class="jh-act-name">{{ pageInfo.Group ? pageInfo.Group.group_name : ''}}</div>
  <div class="jh-act-summary">
    全国流向
    <strong>{{ pageInfo.ProMktBase.weight_min }}-{{ pageInfo.ProMktBase.weight_max }}KG</strong>
    |
    <span>每日承诺寄件量最低<strong>{{ pageInfo.ProMktBase.daily_min_packages }}</strong>件</span>
  </div>
  <div class="jh-qrcode">
    <img src="/static/qrcode.png" alt="">
  </div>
  <div class="jh-share-txt">
    长按二维码保存图片分享给朋友，或者点击右上角分享页面。
  </div>
</div>
</template>

<script>
export default {
  computed: {
    pageInfo () {
      return this.$store.state.pageInfo
    }
  },
  beforeRouteEnter (from, to, next) {
    next(vm => {
      vm.$store.dispatch('SET_TITLE', { title: '顺丰集货助手' })
    })
  }
}
</script>
