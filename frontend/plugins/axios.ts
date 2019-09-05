import axios from 'axios'

export default ({ app }: { app: any }) => {
    axios.defaults.baseURL = app.$env.API_EXTERNAL_URL;
}