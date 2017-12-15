export default {
  SET_INDEX_DATA: (state, data) => {
    state.pageInfo = data
  },
  SET_ADDRESS_DATA: (state, addr) => {
    state.address = addr
  },
  SET_JOIN_DATA: (state, formData) => {
    state.joinInfo = formData
  },
  SET_USER_DATA: (state, { openid }) => {
    state.openid = openid
  },
  SET_TITLE: (state, { title }) => {
    state.title = title
  },
  SET_JOINED: (state, { joined, joinedArr }) => {
    state.joined = joined
    state.joinedArr = joinedArr
  },
  SET_PERCENT: (state) => {
    state.percent = Math.round((parseInt(state.joined * 100) / state.pageInfo['ProMktBase']['group_limit']))
  }
}
