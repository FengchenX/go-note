import { movieApi } from 'api'
import * as TYPE from '../actionType/movieType'

const state = {
	movieList: [],
  movie: {},
  msg: '',
}

const getters = {
	movieList: state => state.movieList,
  movie: state => state.movie,
  msg: state => state.msg
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
	},
  addMovie({commit, state, rootState}, movie) {
	  rootState.requesting = true;
	  commit(TYPE.MOVIE_POST_REQUEST);
    movieApi.addMovie(movie).then((res)=>{
      rootState.requesting = false;
      commit(TYPE.MOVIE_POST_SUCCESS, res);
    }, (error)=>{
      rootState.requesting = false;
      commit(TYPE.MOVIE_POST_FAILURE);
    })
  }
}

const mutations = {
	[TYPE.MOVIE_LIST_REQUEST] (state) {

	},
	[TYPE.MOVIE_LIST_SUCCESS] (state, res) {
		// state.movieList = movieList.data
	
		console.log(res);
		state.movieList = res.data.rows;
	},
	[TYPE.MOVIE_LIST_FAILURE] (state) {

	},
  [TYPE.MOVIE_POST_REQUEST] (state) {

  },
  [TYPE.MOVIE_POST_SUCCESS] (state, res) {
	  state.movie = res.data;
    state.msg = res.msg;
  },
  [TYPE.MOVIE_POST_FAILURE] (state) {

  },
}

export default {
	state,
	getters,
	actions,
	mutations
}
