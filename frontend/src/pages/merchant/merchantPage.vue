<template>
  <div class="q-pa-md">
    <div class="row">
      <router-link to="/dropOff">
        <div class="col-12">
          <q-icon size="2.5rem"  style="color: #00C31E; float: left;" name="arrow_back_ios" />
          <h6 class="q-my-xs" style="color: #00C31E; float: left;">Back</h6>
        </div>
      </router-link> 
    </div>
    <div class="row">
      <div class="col-12">
        <h6 class="q-mb-md text-center">Select Merchant</h6>
      </div>
    </div>
    <div class="row">
      <div v-for="(data, idx) in merchantLists" :key="idx" class="col-12 q-mb-md">
        <q-card @click="goToMerchantProduct(data.is_available, data.id)"  :class="(data.is_available) ? 'myCard':'myCard disabledCard'">
          <q-card-section>
            <div class="row">
              <div class="col-2">
                <q-icon size="2.5rem"  style="color: #00C31E; float: left;" name="restaurant" />
              </div>
              <div class="col-10">
                <div class="row">
                  <div class="col-12">
                    {{ data.name }}
                  </div>
                  <div class="col-12">
                    <q-icon size="1rem"  style="color: #FFC14F; float: left;" name="star_rate" />
                    <p style="float: left;">{{ data.rating }}</p>
                  </div>
                </div>
              </div>
            </div>
          </q-card-section>
        </q-card> 
      </div> 
    </div>
  </div>
</template>

<script>
import { Cookies } from 'quasar'
import { api } from 'src/boot/axios'

export default {
  name: 'merchantPage',
  data () {
    return {
      data: '',
      merchantList: [
        {
          name: 'Nasi Ngadap',
          rating: '4.8',
          isAvailable: true,
        },
        {
          name: 'Burger Bener',
          rating: '5.0',
          isAvailable: true,
        },
        {
          name: 'Warteg Doa Ibu',
          rating: '4.6',
          isAvailable: false,
        }
      ],
      token: null,
      merchantLists: []
    }
  },
  mounted () {
    this.getUserToken()
    this.fetchMerchant()
  },
  methods: {
    goToMerchantProduct (status, merchantId) {
      if (status) {
        this.$router.push(`/merchantProduct/${merchantId}`)
      }
    },
    getUserToken () {
      const token = Cookies.get('user_token')
      this.token = token
    },
    fetchMerchant () {
      const config = {
        headers: { Authorization: `Bearer ${this.token}` }
      }
      api.get('/merchants', config)
      .then((res) => {
        const payload = res.data
        this.merchantLists = payload
        console.log(this.merchantLists)
      })
      .catch((err) => {
        console.log(err)
      })
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