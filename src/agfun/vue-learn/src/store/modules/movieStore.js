import { movieApi } from 'api'
import * as TYPE from '../actionType/movieType'

const state = {
	movieList: []
}

const getters = {
	movieList: state => state.movieList
}

const actions = {
	movieList({commit, state, rootState}) {
		rootState.requesting = true;
		commit(TYPE.MOVIE_LIST_REQUEST);
		movieApi.list().then((response) => {
			rootState.requesting = false
			commit(TYPE.MOVIE_LIST_SUCCESS, response)
		}, (error) => {
			rootState.requesting = false
			commit(TYPE.MOVIE_LIST_FAILURE)
		});
	}
}

const mutations = {
	[TYPE.MOVIE_LIST_REQUEST] (state) {

	},
	[TYPE.MOVIE_LIST_SUCCESS] (state, res) {
		// state.movieList = movieList.data
	
		console.log(res);
		state.movieList = res.Data.rows;
	},
	[TYPE.MOVIE_LIST_FAILURE] (state) {

	}
}

export default {
	state,
	getters,
	actions,
	mutations
}
