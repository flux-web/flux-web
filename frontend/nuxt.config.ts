import config from './config';

export default {
  env: {
    API_EXTERNAL_URL: process.env.API_EXTERNAL_URL,
    API_SERVICE_URL: process.env.API_SERVICE_URL,
    WS_URL: process.env.WS_URL,
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
