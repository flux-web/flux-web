export default {
  env: {
    NUXT_ENV_COOL: process.env.NUXT_ENV_COOL
  },
  head: {
    title: "flux-web",
    meta: [
      { charset: "utf-8" },
      { name: "viewport", content: "width=device-width, initial-scale=1" },
      { hid: "description", name: "description", content: "Flux web" }
    ],
    link: [
      { rel: "icon", type: "image/x-icon", href: "/favicon.ico" }
    ]
  },
  loading: { color: "#3B8070" },
  css: [
    'vue-good-table/dist/vue-good-table.css',
    './assets/scss/_main.scss'
  ],
  mode: 'universal',
  build: { },
  buildModules: ["@nuxt/typescript-build"],
  modules: [
    "@nuxtjs/axios",
    ['nuxt-env', {
      keys: [
        'NODE_ENV', 
        'BASE_URL', 
        'API_EXTERNAL_URL', 
        'WS_URL', 
        'READ_ONLY',
      ]
    }]
  ],
  plugins: [ 
      '~/plugins/vue-good-table',
      { src: '~/plugins/vue-native-websocket', ssr: false },
      { src: '~/plugins/localStorage', ssr: false },
      '~/plugins/axios'
    ],
  axios: {
  },
}
