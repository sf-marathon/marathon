import api from '../api'

export default {
  FETCH_INDEX_DATA: ({ commit, dispatch }) => {
    return api.fetchPageInfo()
      .then(resp => {
        if (resp.data['success']) {
          const info = resp.data['obj']
          info['ProMktBase']['user_require'] = info['ProMktBase']['user_require'].split('\r\n')
          info['ProMktBase']['user_require'] = info['ProMktBase']['user_require'].map(item => {
            item = item.replace(/#([0-9]*)#/g, '<strong>$1</strong>')
            return '<li>' + item + '</li>'
          })
          commit('SET_INDEX_DATA', info)
          return dispatch('FETCH_JOINED', { groupId: info['Group']['group_id'] })
        } else {
          return Promise.reject(new Error(resp.data['errorMessage']))
        }
      })
  },
  FETCH_JOINED: ({ commit, state }, { groupId }) => {
    return api.fetchJoined(groupId)
     .then(resp => {
       if (resp.data['success']) {
         commit('SET_JOINED', { joined: resp.data['obj'].length, joinedArr: resp.data['obj'] })
         commit('SET_PERCENT')
       } else {
         return Promise.reject(new Error(resp.data['errorMessage']))
       }
     })
  },
  FETCH_ADDRESS: ({ commit }) => {
    return api.fetchAddr()
      .then(resp => {
        commit('SET_ADDRESS_DATA', resp.data)
      })
  },
  SET_JOIN_DATA: ({ commit, state }, formData) => {
    let weight = formData.weight.replace('kg', '').split('-')

    const postData = {
      'group_id': state.pageInfo.Group.group_id,
      'phone': formData.phone,
      'address': formData.address + formData.detail,
      'total_amount': parseInt(formData.number),
      'expect_daily_amount': parseInt(weight[1]),
      'average_weight': parseInt(weight[0])
    }
    return api.join(postData)
      .then(resp => {
        if (resp.data['success']) {
          return Promise.resolve()
        } else {
          return Promise.reject(new Error(resp.data['errorMessage']))
        }
      })
    // commit('SET_JOIN_DATA', formData)
  },
  SET_USER: ({ commit }, { openid }) => {
    commit('SET_USER_DATA', { openid })
  },
  SET_TITLE: ({ commit }, { title }) => {
    commit('SET_TITLE', { title })
  }
}
