const state = {
  connected: [],
  disconnected: [],
};

const mutations = {
  connectFriend(state, friend) {
    state.disconnected = state.disconnected.filter(amigo => {
      return amigo.id !== friend.id
    });
    
    let fr = state.connected.find(connectedFr =>
        connectedFr.id === friend.id);
    if (!fr) {
      state.connected.push(friend);
    }
  },
  
  disconnectFriend(state, friend) {
    state.connected = state.connected.filter(amigo => amigo.id !== friend.id);
  
  
    let fr = state.disconnected.find(disconnectedFr =>
        disconnectedFr.id === friend.id);
    if (!fr) {
      state.disconnected.push(friend);
    }
    
    
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
      resolve(
          state.connected.filter((amigo) => amigo.id === friend.id).length > 0);
    });
  }
};

export default {
  state,
  mutations,
  actions,
}
