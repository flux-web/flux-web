import Vue from 'vue';
import App from './App.vue';
import router from './router';
import store from './store/store';
import './registerServiceWorker';
import VueGoodTablePlugin from 'vue-good-table';
import config from './config';
import axios from 'axios';

axios.defaults.baseURL = config.apiBaseUrl;

// import the styles
import 'vue-good-table/dist/vue-good-table.css';
Vue.use(VueGoodTablePlugin);

require('./assets/scss/_main.scss');

Vue.config.productionTip = false;

new Vue({
  router,
  store,
  render: (h) => h(App),
}).$mount('#app');
