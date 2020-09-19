import createPersistedState from 'vuex-persistedstate'

export default ({ store }: { store: any }) => {
    const w = (window as any)

    w.onNuxtReady(() => {
        createPersistedState()(store)
    })
}