<template>
<yd-flexbox direction="vertical">
<yd-flexbox-item>
<div class="jh-home" v-show="loaded">
  <div class="jh-card">
    <div class="jh-act-name">{{ pageInfo.Group ? pageInfo.Group.group_name : ''}}</div>
    <div class="jh-home__act-summary">
      <span>{{ pageInfo.ProMktBase.weight_min }}-{{ pageInfo.ProMktBase.weight_max }}KG</span>
      每日最低需寄<strong>{{ pageInfo.ProMktBase.daily_min_packages }}</strong>件
    </div>
    <div class="jh-home__banner" style="background-image:url('/static/business-wl3.jpg')">
      <div class="jh-float-label">
        <p>低至¥<span>{{ pageInfo.ProMktBase.base_price }}</span></p>
        <p>{{ pageInfo.ProMktBase.base_weight }}KG</p>
      </div>
    </div>
    <div class="jh-home__progress">
      <div class="jh-home-progress">
        <div class="jh-home-progress__bar"
          v-bind:style="{ width: percentage + 'px' }"
          style="width: 50%;">
        </div>
        <div class="jh-home-progress__info">{{ percentage }}%</div>
      </div>
      <div class="jh-home-progress-info">
        还差<strong>{{ parseInt(pageInfo['ProMktBase']['group_limit']) - parseInt(joined) }}</strong>人即可成团
      </div>
    </div>
    <div class="jh-home__deadline">
      截止日期：<time>{{ deadline }}</time>
    </div>
  </div>
  <div class="jh-card">
    <div class="jh-home__rule">
      <h3 class="jh-card__title"><yd-icon name="warn-outline" size="20px" style="margin-right: 5px; vertical-align: middle;"></yd-icon>使用要求</h3>
      <ul class="jh-home-rule__list">
        <li v-for="(item, index) in pageInfo['ProMktBase']['user_require']" :key="index" v-html="item"></li>
      </ul>
    </div>
  </div>
  <div class="jh-card">
    <div class="jh-home__recent">
      <h3 class="jh-card__title">
        <yd-icon name="ucenter-outline" size="20px" style="margin-right: 5px; vertical-align: middle;">
        </yd-icon>最近参与
      </h3>
      <div class="jh-home-recent">
        <ul class="jh-avatar-list">
          <li class="jh-avatar-list__item"><img src="" alt=""></li>
          <li class="jh-avatar-list__item"><img src="" alt=""></li>
          <li class="jh-avatar-list__item"><img src="" alt=""></li>
        </ul>
        <div class="jh-home-recent__txt">已有<strong>{{ joined }}</strong>人参团</div>
      </div>
    </div>
  </div>
</div>
<div class="jh-methods">
  <div class="btn-join" @click="handleJoin">立即参加</div>
  <div class="btn-share" @click="handleShare"></div>
</div>
</yd-flexbox-item>
</yd-flexbox>
</template>

<script>
/*
Name string `json:"name"`
WeightRequire string `json:"weight_require"`
MinAmount int `json:"min_amount"`
BasePrice float32 `json:"base_price"`
BaseWeight float32 `json:"base_weight"`
PictureUrl string `json:"picture_url"`
Percentage float32 `json:"percentage"`
Lack int `json:"lack"`
Joined int `json:"joined"`
Deadline string `json:"deadline"`
Duration string `json:"duration"`
UseRequire string `json:"use_require"`
*/
import moment from 'moment'

export default {
  data () {
    return {
      loaded: false
    }
  },
  beforeRouteEnter (from, to, next) {
    next(vm => {
      vm.$store.dispatch('SET_TITLE', { title: '集货详情' })
    })
  },
  beforeMount () {
    this.$store.dispatch('FETCH_INDEX_DATA')
      .then(() => {
        this.loaded = true
      })
      .catch(err => {
        console.log(err)
      })
  },
  computed: {
    pageInfo () {
      const info = this.$store.state.pageInfo
      return info
    },
    percentage () {
      return this.$store.state.percent
    },
    joined () {
      return this.$store.state.joined
    },
    joinedArr () {
      return this.$store.state.joinedArr
    },
    deadline () {
      const time = moment(this.pageInfo['Group']['due_time']).format('MM月DD日')
      return time
    }
  },
  methods: {
    handleJoin () {
      this.$router.push({ path: '/join' })
    },
    handleShare () {
      this.$router.push({ path: '/share' })
    }
  }
}
</script>
