import Vue from 'vue';
import VueNativeSock from 'vue-native-websocket';
import store from '../store/index';
import config from '../config';

Vue.use(VueNativeSock, 'ws://' + config.wsUrl, {
  store: store(),
  reconnection: true,
  reconnectionAttempts: 5,
  reconnectionDelay: 3000,
  format: 'json',
});