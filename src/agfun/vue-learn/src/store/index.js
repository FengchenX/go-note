import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)


import bannerStore from './modules/bannerStore'
import rankStore from './modules/rankStore'
import promoteStore from './modules/promoteStore'
import liveStore from './modules/liveStore'
import contentStore from './modules/contentStore'

/***********************************************************/
import topStore from './modules/topStore'
import movieStore from './modules/movieStore'
import videoStore from './modules/videoStore'
/***********************************************************/

const state = {
	requesting: false,
	error: {}
}

const getters = {
	requesting: state => state.requesting,
	error: state => state.error
}

export default new Vuex.Store({
	state,
	getters,
	modules: {
		bannerStore,
		rankStore,
		promoteStore,
		liveStore,
		contentStore,
		/******************************************/
		topStore,
    movieStore,
    videoStore
		/******************************************/
	}
})
