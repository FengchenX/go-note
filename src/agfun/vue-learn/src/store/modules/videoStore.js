import { videoApi } from 'api'
import * as TYPE from '../actionType/videoType'

const state = {
	videoList: [],
  total: 0
}

const getters = {
	videoList: state => state.videoList,
  total: state => state.total
}

const actions = {
	getVideos({commit, state, rootState}, params) {
    rootState.requesting = true;
    commit(TYPE.VIDEO_LIST_REQUEST);
    videoApi.list(params).then((response) => {
      rootState.requesting = false
      commit(TYPE.VIDEO_LIST_SUCCESS, response)
    }, (error) => {
      rootState.requesting = false
      commit(TYPE.VIDEO_LIST_FAILURE)
    });
  }
}

const mutations = {
	[TYPE.VIDEO_LIST_REQUEST] (state) {

	},
	[TYPE.VIDEO_LIST_SUCCESS] (state, res) {
		state.videoList = res.data.videos;
		state.total = res.data.total;
	},
	[TYPE.VIDEO_LIST_FAILURE] (state) {

	}
}

export default {
	state,
	getters,
	actions,
	mutations
}
