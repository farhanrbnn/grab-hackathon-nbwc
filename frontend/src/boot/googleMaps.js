import { boot } from 'quasar/wrappers'
import  VueGoogleMaps from '@fawmi/vue-google-maps'

// "async" is optional;
// more info on params: https://v2.quasar.dev/quasar-cli/boot-files
export default boot(async ({ app }) => {
  // something to do
  app.use(VueGoogleMaps, { 
    load: {
      key: 'AIzaSyAXeq6g3HL9uaX2X-kphWHhr-MghMf844A',
      libraries: 'places'
    },
  })
})
