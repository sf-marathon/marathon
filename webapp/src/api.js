import axios from 'axios'

const host = process.env.NODE_ENV === 'production' ? 'http://sfai.ml:8080' : 'http://sfai.ml:8080'

const apiURI = {
  pageInfo: `${host}/ca/group`,
  join: `${host}/ca/join`,
  address: `/static/citydata.json`
}

export default {
  fetchPageInfo () {
    return axios.get(apiURI.pageInfo)
  },
  fetchAddr () {
    return axios.get(apiURI.address)
  },
  join (data) {
    return axios.post(apiURI.join, data)
  },
  fetchJoined (groupId) {
    return axios.get(apiURI.join + '/' + groupId)
  }
}
