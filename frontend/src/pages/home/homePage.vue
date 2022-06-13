<template>
  <div class="q-pa-md" style="padding-bottom: 100px;">
    <div class="row">
      <div class="col-6">
        <p class="text-left q-mt-xs">Hi, {{ userData.name }}</p>
      </div>
      <div class="col-6">
        <router-link to="/account">
          <q-icon size="2rem"  style="color: #00C31E; float: right;" name="account_circle" />
        </router-link>
      </div>
    </div>
    <div class="row">
      <div class="col-12">
        <h6 class="q-my-xs"> Your Location</h6>
        <div class="row q-mb-md">
          <div class="col-12 inline">
            <q-icon size="2rem"  style="color: #FF3C3C; float: left;" name="location_on" />
            <a href="/#/userLocation" class="q-my-xs" style="float: left; color: black;">{{ address }}</a>
          </div> 
        </div>
      </div>
    </div>
    <div class="row">
      <div class="col-12">
        <q-card class="my-card">
          <q-card-section>
            <div class="row">
              <div class="col-12">
                <div class="row">
                  <div class="col-8 q-mb-md">
                    <p class="q-mb-none"> Since using Soombang, you have donated as much </p>
                    <h6 class="q-my-xs" style="color: #00C31E">Rp. 2.000.000</h6>
                  </div>
                  <div class="col-4"> 
                    <img
                    alt="money" 
                    src="~assets/money.svg"
                    style="width: 100%; height: 100%;"
                    >
                  </div>
                </div>
              </div>
              <div class="col-2">
                <a href="/#/donateStat" style="color: #00C31E;">
                  Detail
                  <q-icon class="q-mt-xs" style="color: #00C31E; float: right;" name="chevron_right" />
                </a>
              </div>
            </div>
          </q-card-section>
        </q-card>
      </div>
    </div>
    <div class="row q-mt-md">
      <div class="col-12">
        <q-card class="my-card">
          <q-card-section>
            <div class="row">
              <div class="col-8">
                <p class="q-mt-none">
                  <b>Saldo Anda</b>   
                </p>
                <img
                  alt="ovo" 
                  src="~assets/OVO.svg"
                  style="width: 40px; height: 40px; float: left;"
                >
                <h6 class="q-my-none q-ml-s" style="color: #00C31E; float: left;"> Rp. {{ userWallet.effective_amount }}</h6>
              </div>
              <div class="col-4">
                <router-link to="/changeWallet">
                  <q-icon class="q-my-md" size="3rem" style="color: #00C31E; float: right;" name="apps" />
                </router-link>
              </div>
            </div>  
          </q-card-section>
        </q-card>
      </div>
    </div>
    <div class="row q-mt-md">
      <div class="col-12">
        <q-btn style="background: #00C31E; color: white" @click="goToSourceFund()" icon="add_circle" size="lg" class="full-width" label="Donate Now" />
      </div>
    </div>
    <div class="row q-mt-md">
      <div class="col-6">
        <h6 class="q-my-md">Last Transaction</h6>
      </div>
      <div class="col-6">
        <a href="/#/orderHistory"  class="q-my-md" style="float: right; color:#00C31E;">See More</a>
      </div>
      <div v-if="listTransaction.length === 0" class="col-12"> 
        <h6 class="text-center">You don't have any transactions yet</h6>
      </div> 
      <div @click="goToHistoryDetail()"  v-for="(data, idx) in listTransaction" :key="idx" class="col-12 q-mb-xs">
        <q-card class="my-card">
          <q-card-section>
            <div class="row">
              <div class="col-2">
                <q-icon size="2rem" style="color: #00C31E; float: left;" name="receipt_long" />
              </div>
              <div class="col-6">
                <p>{{ data.donateTo }}</p>
                <p class="q-mb-none">{{ data.date }}</p>
              </div>
              <div class="col-4">
                <p class="text-right">
                  <b> {{ data.total }} </b>
                </p>
              </div>
            </div>
          </q-card-section>
        </q-card>
      </div> 
    </div>
  </div>
  <footer-menu style="margin-top: 200px;" />
</template>

<script>
import axios from 'axios'
import { Cookies } from 'quasar'
import footerMenu from '../../components/footerMenu/footerMenu.vue'
import { api } from 'src/boot/axios'

export default {
  name: 'homePage',
  components: {
    footerMenu
  }, 
  data () {
    return {
      lastTransactions: [
        {
          donateTo: 'donation to panti asuhan A',
          date: '25 Mei 2022',
          total: 'Rp. 250.000'
        },
        {
          donateTo: 'donation to panti asuhan B',
          date: '28 Mei 2022',
          total: 'Rp. 500.000'
        },
        {
          donateTo: 'donation to panti asuhan B',
          date: '28 Mei 2022',
          total: 'Rp. 500.000'
        }
      ],
      address: null,
      position: null,
      token: null,
      userData: {},
      userWallet: {},
      walletOrigin: [],
      listTransaction: []
    }
  },
  created () {
    this.setUserGeo()
  },
  mounted () {
    this.getUserToken()
    this.fetchUserData()
    this.fetchUserWallet()
    this.fetchUserTransaction()
  },
  methods: {
    fetchUserTransaction () {
      const config = {
        headers: { Authorization: `Bearer ${this.token}` }
      }

      api.get('/transactions', config)
      .then((res) => {
        const payload = this.res
        
        if (payload !== undefined) {
          this.listTransaction = payload.data 
        }

        this.listTransaction = [] 
        console.log(this.listTransaction)
      })
      .catch((err) => {
        console.log(err)
      })
    },
    async fetchUserWallet () {
      const config = {
        headers: { Authorization: `Bearer ${this.token}` }
      }
      await api.get('/user-wallets', config)
      .then((res) => {
        const stringifyData = JSON.stringify(res.data[0])

        const payload = res.data[0]
        payload.effective_amount = this.formatPrice(payload.effective_amount)

        this.userWallet = payload

      })
      .catch((err) => {
        console.log(err)
      }) 
    },
    formatPrice (value) {
      const val = (value/1).toFixed(2).replace('.', ',')
      return val.toString().replace(/\B(?=(\d{3})+(?!\d))/g, ".")
    },
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
    goToHistoryDetail () {
      this.$router.push('/orderHistoryDetail')
    },
    getAddress () {
      axios.get('https://maps.googleapis.com/maps/api/geocode/json?latlng=' + this.position.lat +','+ this.position.lng +'&key=AIzaSyAXeq6g3HL9uaX2X-kphWHhr-MghMf844A')
      .then((res) => {
        const address = res.data.results[0].formatted_address
        const split = address.split(",")
        this.address = split[0] + ' ' + split[1]

        Cookies.set('geolocation', this.position)
        Cookies.set('user_address', this.address)
      })
    },
    goToSourceFund () {
      this.$router.push('/sourceFund')
    },
    async setUserGeo () {
      await navigator.geolocation.getCurrentPosition((position) => {
        this.position = {
          lat: position.coords.latitude,
          lng: position.coords.longitude, 
        }
        this.getAddress()
      })
    }
  }
}
</script>

<style scoped>

</style>
