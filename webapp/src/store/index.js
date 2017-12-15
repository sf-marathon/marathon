import Vue from 'vue'
import Vuex from 'vuex'
import actions from './actions'
import mutations from './mutations'

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    pageInfo: {},
    joinInfo: {},
    address: null,
    title: '',
    percent: 0,
    joined: 0,
    joinedArr: null
  },
  actions,
  mutations
})
