<template>
  <div class="q-pa-md" style="height: 100vh;">
    <div class="row q-mt-xl">
      <div class="col-12 text-center">
          <q-icon size="5rem" class="text-center"  style="color: #00C31E;" name="account_circle" />
          <p class="q-mt-md">{{ userData.name }}</p>
      </div>
    </div>
    <q-form>
      <div class="row">
        <div class="col-12 q-mb-md">
          <p class="q-mb-none"><b>Email</b></p>
          <q-input v-model="userData.email" color="green"  />
        </div>
        <div class="col-12 q-mb-md">
          <p class="q-mb-none"><b>Phone Number</b></p>
          <q-input v-model="userData.phone" color="green"  />
        </div>
      </div>
    </q-form>
    <div class="row">
      <div class="col-12">
        <h6 class="text-center" style="color: #FF5252;">Log Out</h6>
      </div>
    </div>
    <footer-menu />
  </div>
</template>

<script>
import { api } from 'src/boot/axios'
import footerMenu from '../../components/footerMenu/footerMenu.vue'
import { Cookies } from 'quasar'

export default {
  name: 'accountPage',
  components: {
    footerMenu
  },
  data () {
    return {
      userData: {}
    }
  },
  mounted () {
    this.getUserToken()
    this.fetchUserData()
  },
  methods: {
    fetchUserData () {
      const config = {
        headers: { Authorization: `Bearer ${this.token}` }
      }
      api.get('/user', config)
      .then((res) => {
        this.userData = res.data
      })
      .catch((err) => {
        console.log(err)
      })
    },
    getUserToken () {
      const token = Cookies.get('user_token')
      this.token = token
    },
  }
}
</script>
