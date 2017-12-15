<template>
  <div class="jh-popbox" v-show="visible">
    <div class="jh-popbox__content">
      <div class="jh-icon-success"></div>
      <div class="jh-popbox__title">报名成功！</div>
      <div class="jh-popbox__txt">转发给朋友或者朋友圈，即可加快集货进度</div>
      <div class="jh-popbox__btns">
        <div class="btn-pop-share" @click="handleShare">分享</div>
      </div>
      <div class="jh-popbox__close" @click="handleClose"></div>        
    </div>
  </div>
</template>

<script>
export default {
  props: ['visible', 'status'],
  data () {
    return {}
  },
  methods: {
    handleShare () {
      this.$emit('btn-click')
    },
    handleClose () {
      this.$emit('update', false)
    }
  }
}
</script>

