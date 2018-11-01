const parseIp = require('p2p/services/util').parseIpAndPortFromString;

const state = {
  peer: {},
};

const mutations = {
  createPeer(state, peer) {
    state.peer = peer
  },
  setOnNewConnection(state, onNewConnection) {
    state.peer.setOnNewConnection(onNewConnection);
  },
  runServer(state) {
    const port = parseIp(state.peer.client.ips[0]).port;
    state.peer.runServer(port);
  }
};

const getters = {};

export default {
  state,
  mutations,
  getters,
}
