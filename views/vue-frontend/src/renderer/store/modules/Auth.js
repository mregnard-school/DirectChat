const state = {
  token: ''
};

const mutations = {
  setToken(state, token) {
    state.token = token;
  },
  removeToken(state) {
    state.token = '';
  }
};

const actions = {};

export default {
  state,
  mutations,
  actions
}
