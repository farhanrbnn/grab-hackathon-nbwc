import { boot } from 'quasar/wrappers'
import axios from 'axios'


// "async" is optional;
// more info on params: https://v2.quasar.dev/quasar-cli/boot-files
const api = axios.create({ 
  headers: {
    'Content-Type': 'application/json; charset=utf-8'
  },
  baseURL: 'http://ec2-52-221-244-108.ap-southeast-1.compute.amazonaws.com:9090/api/v1'
 })

export default boot(({ app }) => {
  // something to do
  app.config.globalProperties.$axios = axios
  app.config.globalProperties.$api = api
})

export { axios, api }


