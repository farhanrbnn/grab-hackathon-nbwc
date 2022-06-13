<template>
  <div class="q-pa-md" style="height: 100vh;">
    <div class="row">
      <router-link to="/home">
        <div class="col-12">
          <q-icon size="2.5rem"  style="color: #00C31E; float: left;" name="arrow_back_ios" />
          <h6 class="q-my-xs" style="color: #00C31E; float: left;">Back</h6>
        </div>
      </router-link> 
    </div>
    <div class="row">
      <div class="col-12">
        <h5 class="text-center">Change Wallet</h5>
      </div>
    </div>
    <div class="row">
      <div v-for="(data, idx) in wallets" :key="idx" class="col-12 q-mb-xs">
         <q-card class="my-card">
          <q-card-section>
            <div class="row">
              <div class="col-2">
                <img
                  :alt="data.bankName " 
                  :src="require(`../../assets/${data.img}.svg`)"
                  style="width: 40px; height: 40px; float: left;"
                >
              </div>
              <div class="col-10">
                <h5 class="q-my-none text-center">{{ data.name }}</h5>
              </div>
            </div> 
          </q-card-section>
        </q-card>
      </div>
    </div>
  </div>
</template>

<script>
import { api } from 'src/boot/axios'
import { Cookies } from 'quasar'

export default {
  name: 'changeWallet',
  data () {
    return {
      data: '',
      token: null,
      wallets: []
    }
  },
  mounted () {
    this.getUserToken()
    this.fetchWallets()
  },
  methods: {
    getUserToken () {
      const token = Cookies.get('user_token')
      this.token = token
    },
    fetchWallets () {
      const config = {
        headers: { Authorization: `Bearer ${this.token}` }
      }

      api.get('/wallets', config)
      .then((res) => {
        const payload = res.data
        for (let i = 0; i < payload.length; i++) {
          const splitName = payload[i].name.split(" ")

          payload[i].img = splitName[0]
        }
        this.wallets = payload
      })
      .catch((err) => {
        console.log(err)
      })
    },
    doChangeWallet (status) {
      if (status) {
        this.$router.push('/home')
      }
    }
  }
}
</script>

<style scoped>
.disabledCard {
  opacity: 0.6;
  filter: grayscale(100%);
}
</style>