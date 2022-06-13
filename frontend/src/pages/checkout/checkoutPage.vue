<template>
  <div class="q-ma-md" style="padding-bottom: 200px;">
    <div class="row">
      <router-link to="/merchantProduct">
        <div class="col-12">
          <q-icon size="2.5rem"  style="color: #00C31E; float: left;" name="arrow_back_ios" />
          <h6 class="q-my-xs" style="color: #00C31E; float: left;">Back</h6>
        </div>
      </router-link> 
    </div>
    <div class="row">
      <div class="col-12">
        <h6 class="text-center q-mt-md q-mb-sm">Burger Bener</h6>
      </div>
    </div>
    <div class="row">
      <div class="col-12">
        <h6 class="q-mb-xs q-mt-md">Deliver To</h6>
      </div>
      <div class="col-12">
        <q-card class="my-card">
          <q-card-section>
            <div class="row">
              <div class="col-2">
                <q-icon size="2rem"  style="color: #FF3C3C; float: left;" name="location_on" />
              </div>
              <div class="col-10">
                <p class="q-mb-none q-mt-xs">Jalan Kemanggisan Raya</p>
              </div>
            </div>
          </q-card-section>
        </q-card>
      </div>
    </div>
    <div class="row">
      <div class="col-12">
        <h6 class="q-mt-md q-mb-none">Order Summary</h6>
      </div>
    </div>
    <div class="row">
      <div class="col-12">
        <q-card class="my-card">
          <q-card-section>
            <div v-for="data in cart" :key="data.productId" class="row">
              <div class="col-2">
                <p class="q-my-none">{{ data.quantity }}x</p>
              </div>
              <div class="col-6">
                <p class="q-my-none">{{ data.name }}</p>
              </div>
              <div class="col-4">
                <p class="q-my-none text-right">Rp. {{ formatPrice(data.productPrice) }}</p>
              </div>
            </div>
            <div class="row q-mt-xl">
              <div class="col-6">
                <p class="q-my-none">Subtotal</p>
              </div>
              <div class="col-6">
                <p class="q-my-none text-right">Rp. {{ formatSubtotalPrice }}</p>
              </div>
            </div>
            <div class="row">
              <div class="col-6">
                <p class="q-my-none">Delivery Fee</p>
              </div>
              <div class="col-6">
                <p class="q-my-none text-right">Rp. {{ formatPrice(deliveryFee)}}</p>
              </div>
            </div>
            <div class="row">
              <div class="col-6">
                <p class="q-my-none">PIC Fee</p>
              </div>
              <div class="col-6">
                <p class="q-my-none text-right">Rp. {{ formatPrice(picFee)}}</p>
              </div>
            </div>
            <div class="row">
              <div class="col-6">
                <p class="q-my-none">Admin Fee</p>
              </div>
              <div class="col-6">
                <p class="q-my-none text-right">Rp. {{ formatPrice(adminFee)}}</p>
              </div>
            </div>
          </q-card-section>
        </q-card>
      </div>
    </div>
    <div class="row">
      <div class="col-12">
        <h6 class="q-mt-md q-mb-none">Funding Source</h6>
      </div>
      <div class="col-12">
        <q-card class="my-card">
          <q-card-section>
            <div class="row">
              <div class="col-2">
                <q-icon size="2rem"  style="color: #00C31E; float: left;" name="person" />
              </div>
              <div class="col-10">
                <p class="q-my-none q-mt-sm">{{ fundSource.source }}</p>
              </div>
            </div>
            <div class="row">
              <div class="col-6">
               <p class="q-mt-md q-mb-none">
                  <b>Total Funding</b> 
                 </p> 
              </div>
              <div class="col-6">
               <p class="q-mt-md q-mb-none text-right">
                 <b> Rp. {{ formatPrice(fundSource.amount) }}</b>
               </p> 
              </div>
            </div>
          </q-card-section>
        </q-card>
      </div>
    </div>
    <div class="row">
      <div class="col-12">
        <h6 class="q-mt-md q-mb-none">Payment Method</h6>
      </div>
      <div class="col-12">
        <q-card class="my-card">
          <q-card-section>
            <div class="row">
              <div class="col-2">
                <img
                  :alt="payment.name" 
                  :src="payment.img"
                  style="width: 40px; height: 40px; float: left;"
                >
              </div>
              <div class="col-10">
                <p class="q-my-none q-mt-sm">{{ payment.name }}</p>
              </div>
            </div>
          </q-card-section>
        </q-card>
      </div>
    </div>
  </div>
  <checkout-button :grandTotal="grandTotal" :merchantId="this.$route.params.merchantId" />
</template>

<script>
import { format } from 'quasar'
import checkoutButton from '../../components/checkoutButton/checkoutButton.vue'
import { api } from 'src/boot/axios'
import { Cookies } from 'quasar'

export default {
  name: 'checkoutPage',
  components: {
    checkoutButton
  },
  data () {
    return {
      payment: {
        name: 'OVO Cash',
        img: require('../../assets/OVO.svg') 
      },
      // fundingSource: {
      //   name: 'KitaBisa',
      //   img: require('../../assets/kitabisa.png')
      // },
      cart: null,
      formatSubtotalPrice: null,
      subtotalPrice: null,
      deliveryFee: 20000,
      adminFee: 900,
      picFee: 10000,
      grandTotal: null,
      token: null,
      fundSource: {},
      merchant: {},
      products: []
    }
  },
  mounted () {
    this.getUserToken()
    this.getCartData()
    this.calculateSubtotal()
    this.calculateGrandTotal()
    this.fetchProduct()
    this.fetchMerchant()
    this.fundingSource()
  },
  methods: {
    fundingSource () {
      const fundData = localStorage.getItem('funding_source')
      const parsedData = JSON.parse(fundData)

      this.fundSource = parsedData

    },
    getUserToken () {
      const token = Cookies.get('user_token')
      this.token = token
    },
    fetchMerchant () {
      const config = {
        headers: { Authorization: `Bearer ${this.token}` }
      }

      api.get(`/merchant/${this.$route.params.merchantId}`, config)
      .then((res) => {
        this.merchant = res.data
        const stringifyData = JSON.stringify(this.merchant)

        localStorage.setItem('merchant', stringifyData )
      })
      .catch((err) => {
        console.log(err)
      })
    },
    fetchProduct () { 
      const config = {
        headers: { Authorization: `Bearer ${this.token}` }
      }
      for (let i = 0; i < this.cart.length; i++) {
        api.get(`/product/${this.cart[i].productId}`, config)
        .then((res) => {
          const payload = res.data
          const productObj = {}

          productObj.product = payload
          productObj.quantity = this.cart[i].quantity

          this.products.push(productObj)

          const stringifyData = JSON.stringify(this.products)

          localStorage.setItem('products', stringifyData)
        })
        .catch((err) => {
          console.log(err)
        })
      }
    },
    calculateGrandTotal () {
      const sum = this.subtotalPrice + this.deliveryFee + this.adminFee + this.picFee
      this.grandTotal = sum
    },
    calculateSubtotal () {   
      let sum = 0

      for (let i = 0; i < this.cart.length; i++) {
        const quantity = this.cart[i].quantity
        const price = this.cart[i].productPrice
        const calculatePrice = quantity * price

        sum += calculatePrice 
      }
      this.subtotalPrice = sum 

      const formatSubtotal = this.formatPrice(sum)
      this.formatSubtotalPrice = formatSubtotal
    },
    formatPrice (value) {
      const val = (value/1).toFixed(2).replace('.', ',')
      return val.toString().replace(/\B(?=(\d{3})+(?!\d))/g, ".")
    },
    getCartData () {
      const localData = localStorage.getItem('cart')
      const parsedData = JSON.parse(localData)

      this.cart = parsedData
    }
  }
}
</script>
