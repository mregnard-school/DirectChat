import store from './store';
import router from './router';
import mutStore from './mutableStore';
import {parseIpAndPortFromString} from 'p2p/services/util';
import Client from 'p2p/client/client';

const ifs = require('os').networkInterfaces();
const result = Object.keys(ifs)
    .map(x => ifs[x].filter(x => x.family === 'IPv4' && !x.internal)[0])
    .filter(x => x)[0].address;


const ip = {
  address: result
};

const userAuthed = (client) => {
  client.ips = JSON.parse(JSON.stringify(client.Ips || [])); //Because GORM SUCKS A BIG FAT BAG OF DICKS
  client.friends = JSON.parse(JSON.stringify(client.Friends || []));
  let peer = new Client(client);
  store.commit('setToken', client.token);
  mutStore.push({
    peer: peer,
  });
  peerCreated(peer);
  router.push('/home');
};

const peerCreated = (peer) => {
  if(!peer.client.ips[0]) {
    peer.client.ips[0].address = '127.0.0.1:5000';
  }
  const port = parseIpAndPortFromString(peer.client.ips[0].address).port;
  peer.runServer(port);
  
  peer.client.friends.forEach(friend => {
    if (friend.ips.length !== 0) {
      friend.ips.forEach(ipPort => {
        const parsed = parseIpAndPortFromString(ipPort);
        peer.node.connectTo(parsed.ip, parsed.port);
      })
    }
  });
};

export {ip, userAuthed};