import { topApi } from 'api'
import * as TYPE from '../actionType/topType'

const state = {
	toplist: []
}

const getters = {
	toplist: state => state.toplist
}

const actions = {
	toplist({commit, state, rootState}) {
		console.log('111111111111111111111111111111111');
		rootState.requesting = true
		commit(TYPE.TOP_LIST_REQUEST)
		topApi.list().then((response) => {
			rootState.requesting = false
			commit(TYPE.TOP_LIST_SUCCESS, response)
		}, (error) => {
			rootState.requesting = false
			commit(TYPE.TOP_LIST_FAILURE)
		})
	}
}

const mutations = {
	[TYPE.TOP_LIST_REQUEST] (state) {

	},
	[TYPE.TOP_LIST_SUCCESS] (state, toplist) {
		// state.toplist = toplist.data
		console.log(toplist);
		state.toplist = toplist.Data.data
	},
	[TYPE.TOP_LIST_FAILURE] (state) {

	}
}

export default {
	state,
	getters,
	actions,
	mutations
}