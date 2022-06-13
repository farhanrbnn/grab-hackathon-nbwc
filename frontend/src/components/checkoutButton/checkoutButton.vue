<template>
  <div class="row q-px-sm q-py-md fixed-bottom shadow-2" style="background-color: #FFF;">
    <div class="col-12 q-mb-sm">
      <div class="row">
        <div class="col-6">
          <h6 class="q-my-md">Grand Total</h6>
        </div>
        <div class="col-6">
          <h6 class="q-my-md text-right"> Rp. {{ formatPrice(grandTotal) }} </h6>
        </div>
      </div>
    </div>
    <div class="col-12">
      <q-btn @click="createOrder()" style="background: #00C31E; color: white" class="full-width" label="Place Order" />
    </div>
  </div> 
</template>

<script>
import { api } from 'src/boot/axios'
import { Cookies } from 'quasar'

export default {
  name: 'checkoutButton',
  props: {
    grandTotal: {
      type: Number
    },
    merchantId: {
      type: String
    }
  },
  data () {
    return {
      total: null,
      token: null  
    }
  },
  mounted () {
    this.getUserToken()
  },
  methods: {
    formatPrice (value) {
      const val = (value/1).toFixed(2).replace('.', ',')
      return val.toString().replace(/\B(?=(\d{3})+(?!\d))/g, ".")
    }, 
    getUserToken () {
      const token = Cookies.get('user_token')
      this.token = token
    },
    createOrder () {
      const config = {
        headers: { Authorization: `Bearer ${this.token}` }
      }
      const fundingSource = localStorage.getItem('funding_source')
      const dropOffLocation = localStorage.getItem('drop_off_location')
      const manifest = localStorage.getItem('merchant')
      const products = localStorage.getItem('products')

      const parsedFund = JSON.parse(fundingSource)
      const parsedDropOff = JSON.parse(dropOffLocation)
      const parsedManifest = JSON.parse(manifest)
      const parsedProducts = JSON.parse(products)

      const checkout = {
        manifest: {},
        products: []
      }

      checkout.funding_source = parsedFund
      checkout.manifest.merchant = parsedManifest
      checkout.drop_off_location = parsedDropOff
      checkout.products = parsedProducts

      console.log(checkout)
      api.post('/order/create', checkout, config)
      .then((res) => {
        const payload = res.data
        this.$router.push(`/activityDetail/${payload.transaction.id}`)

        localStorage.removeItem('funding_source')
        localStorage.removeItem('drop_off_location')
        localStorage.removeItem('merchant')
        localStorage.removeItem('products')

      })
      .catch((err) => {
        console.log(err)
      })
    }
  }
}
</script>
