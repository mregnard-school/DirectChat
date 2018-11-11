import store from './store';
import router from './router';
import mutStore from './mutableStore';
import {parseIpAndPortFromString} from 'p2p/services/util';
import Client from 'p2p/client/client';
import {wrapper} from "./axios-wrapper";

const ifs = require('os').networkInterfaces();
const result = Object.keys(ifs)
    .map(x => ifs[x].filter(x => x.family === 'IPv4' && !x.internal)[0])
    .filter(x => x)[0].address;


const ip = {
  address: result
};

const userAuthed = (client) => {
  let peer = new Client(client);
  store.commit('setToken', client.token);
  wrapper.setToken();
  
  mutStore.push({
    peer: peer,
  });
  peerCreated(peer);
  router.push('/home');
};


const peerCreated = (peer) => {
  const port = peer.client.ips[0].port;
  peer.runServer(port);
  if(!peer.client.friends) {
    peer.client.friends = [];
  }
  peer.client.friends.forEach(friend => {
    if(!friend.ips) {
      friend.ips = [];
    }
    if (friend.ips.length !== 0) {
      friend.ips.forEach(ipPort => {
        peer.node.connectTo('127.0.0.1', ipPort.port);
      })
    }
  });
};

export {ip, userAuthed};