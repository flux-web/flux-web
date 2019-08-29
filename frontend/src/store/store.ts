import Vue from 'vue';
import Vuex, { StoreOptions } from 'vuex';
import createPersistedState from 'vuex-persistedstate';
import { root } from './root/index';
import { RootState } from './types/RootState';
import {actions} from '@/store/root/actions';
import {mutations} from '@/store/root/mutations';
import {getters} from '@/store/root/getters';
import { workloads } from './workloads/index';
import { namespaces } from './namespaces/index';

Vue.use(Vuex);

const store: StoreOptions<RootState> = {
    state: root,
    getters,
    actions,
    mutations,
    modules: {
        workloads,
        namespaces,
    },
    plugins: [createPersistedState()],
};

export default new Vuex.Store<RootState>(store);
