import Vue from 'vue';
import VueNativeSock from 'vue-native-websocket';
import config from '../config';

export default ({ store }: { store: any }) => {
  Vue.use(VueNativeSock, 'ws://' + config.wsUrl, {
    store,
    reconnection: true,
    reconnectionAttempts: 5,
    reconnectionDelay: 3000,
    format: 'json',
  });
}

