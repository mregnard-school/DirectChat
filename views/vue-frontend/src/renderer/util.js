import store from './store';
import router from './router';
import mutStore from './mutableStore';
import {parseIpAndPortFromString} from 'p2p/services/util';
import Client from 'p2p/client/client';
import {wrapper} from "./axios-wrapper";
import { sha256 } from 'js-sha256';

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
        peer.node.connectTo(ipPort.address, ipPort.port);
      })
    }
  });
};

const hashConversation = (friends) => {
  const pseudos = friends.map(friend => friend.pseudo).sort().join(",");
  const hash = sha256(pseudos);
  console.log(hash);
  return hash;
};

export {ip, userAuthed, hashConversation};