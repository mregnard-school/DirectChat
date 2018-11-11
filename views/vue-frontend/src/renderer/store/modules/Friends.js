const state = {
  connected: [],
  disconnected: [],
};

const mutations = {
  connectFriend(state, friend) {
    state.disconnected = state.disconnected.filter(amigo => {
      return amigo.id !== friend.id
    });
    
    state.connected.push(friend);
  },
  
  disconnectFriend(state, friend) {
    state.connected = state.connected.filter(amigo => amigo.id !== friend.id);
    state.disconnected.push(friend);
  },
  removeFriends(state) {
    state.connected = [];
    state.disconnected = [];
  }
};

const getters = {};

const actions = {
  isConnected: ({state}, friend) => {
    return new Promise((resolve) => {
      resolve(state.connected.filter((amigo) => amigo.id === friend.id).length > 0);
    });
  }
};

export default {
  state,
  mutations,
  actions,
}
