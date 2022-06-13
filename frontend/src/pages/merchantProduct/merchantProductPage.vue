<template>
  <div class="q-pa-none">
    <div class="row">
      <div class="row q-mb-md" style="position: absolute !important; top: 10px; left: 10px; z-index: 1;">
          <router-link to="/merchant">
            <div class="col-12">
              <q-icon size="3rem"  style="color: #00C31E; float: left;" name="arrow_circle_left" />
            </div>
          </router-link> 
        </div>
      <div class="col">
        <img
          alt="ovo" 
          src="~assets/warteg.jpeg"
          style="max-width: 100%"
        >
      </div>    
    </div>
  </div>
  <div class="q-pt-none q-px-md">
    <div class="row">
      <div class="col-12">
        <q-card class="my-card">
          <q-card-section>
            <div class="row">
              <div class="col-12">
                <h6 class="q-my-xs text-center">{{ merchant.name }}</h6>
              </div>
            </div>
          </q-card-section>
        </q-card>
      </div>
    </div>
    <div class="row">
      <div class="col-12">
        <h6 class="q-my-md"> 
          Product List
        </h6>
      </div>
    </div>
    <div class="row justify-center q-gutter-xs" style="padding-bottom: 100px;">
      <q-intersection
        v-for="data in products"
        :key="data.id"
        class="example-item"
      >
        <q-card class="q-ma-sm">
          <img :src="require('../../assets/food.jpg')" style="width:11em;">
          <q-card-section>
            <p class="q-my-none">
              <b>{{ data.name }}</b>
            </p>
            <p class="q-my-xs">Rp. {{ formatPrice(data.price) }}</p>
            <div class="row q-mt-md">
              <div class="col-2">
                <q-icon size="2rem" @click="removeProduct(data)"  style="color: #00C31E;" name="remove" />
              </div>
              <div class="col-8">
                <h6 class="q-my-none text-center">{{ qtyCounter(data.id) }}</h6>
              </div>
              <div class="col-2">
                <q-icon size="2rem" @click="addProduct(data)"  style="color: #00C31E;" name="add" />
              </div>
            </div>
          </q-card-section>
        </q-card>
      </q-intersection>
    </div>
    <div v-if="showFooter" class="row q-mx-lg q-mb-lg fixed-bottom" style="margin-top: 200px; background-color: #FFF;" >
      <div class="col-12">
        <footer-product-display @click="goToCheckout()" :localData="localData" />
      </div>
    </div>
  </div>
</template>

<script>
import footerProductDisplay from '../../components/footerProduct/footerProductDisplay.vue'
import { Cookies } from 'quasar'
import { api } from 'src/boot/axios'

export default {
  name: 'merchantProductPage',
  components: {
    footerProductDisplay
  },
  data () {
    return {
     listProduct: [
       {
         productId: 1,
         name: 'Burger 1',
         description: 'lorem ipsum',
         price: 50000,
         images: require('../../assets/burger.jpeg') 
       },
       {
         productId: 2,
         name: 'Burger 2',
         description: 'lorem ipsum',
         price: 20000,
         images: require('../../assets/burger.jpeg') 
       },
       {
         productId: 3,
         name: 'Burger 2',
         description: 'lorem ipsum',
         price: 20000,
         images: require('../../assets/burger.jpeg') 
       },
       {
         productId: 4,
         name: 'Burger 2',
         description: 'lorem ipsum',
         price: 20000,
         images: require('../../assets/burger.jpeg') 
       }
     ],
     showFooter: false,
     cart: null,
     localData: null,
     token: null,
     products: [],
     merchant: {}
    }
  },
  watch: {
    localData (val) {
      this.qtyCounter()
      console.log(this.localData)
    }
  },
  mounted () {
    this.footerStatus()
    this.getLocalData()
    this.getUserToken()
    this.fetchMerchantProduct()
    this.fetchMerchant()
  },
  methods: {
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
        const payload = res.data
        this.merchant = payload
      })
      .catch((err) => {
        console.log(err)
      })
    },
    fetchMerchantProduct () {
      const config = {
        headers: { Authorization: `Bearer ${this.token}` }
      }

      api.get('/products', config)
      .then((res) => {
        const payload = res.data
        this.products = payload
        console.log(res.data)
      })
      .catch((err) => {
        console.log(err)
      })
    },
    qtyCounter (id) {
      const localData = localStorage.getItem('cart')
      const parsedData = JSON.parse(localData)

      if (localData !== null) {
        const product = parsedData.filter(d => d.productId === id)

        if (product.length > 0) { 
          return product[0].quantity
        }
      }

      return 0 
    },
    goToCheckout (id) {
      this.$router.push(`/checkout/${this.$route.params.merchantId}`)
    },
    getLocalData () {
      const localData = localStorage.getItem('cart')
      this.localData = localData
    },
    formatPrice (value) {
      const val = (value/1).toFixed(2).replace('.', ',')
      return val.toString().replace(/\B(?=(\d{3})+(?!\d))/g, ".")
    },
    footerStatus () {
      const localData = localStorage.getItem('cart')
      const parsedData = JSON.parse(localData)

      if (localData !== null) {
        if (parsedData.length > 0) {
        this.showFooter = true 
      } else {
        this.showFooter = false 
      }
      } 
    },
    addProduct (val) {
      this.showFooter = true
      const localData = localStorage.getItem('cart')
      const parsedData = JSON.parse(localData)

      if (localData === null) {
        const data = [{
          productId: val.id,
          quantity: 1,
          name: val.name,
          productPrice: val.price 
        }]

        const stringifyData = JSON.stringify(data)
        localStorage.setItem('cart', stringifyData)

        this.localData = JSON.parse(stringifyData) 

      } else {
        const isAvailable = parsedData.filter(d => d.productId === val.id)

        if (isAvailable.length != 0) { 
          isAvailable[0].quantity = isAvailable[0].quantity + 1

          const stringifyData = JSON.stringify(parsedData)

          localStorage.setItem('cart', stringifyData)

        } else {
          const newProduct = {
            productId: val.id,
            quantity: 1,
            name: val.name,
            productPrice: val.price 
          }

          parsedData.push(newProduct)

          const stringifyData = JSON.stringify(parsedData)

          localStorage.setItem('cart', stringifyData)
        }
        this.localData = parsedData 
      }
    },
    removeProduct (val) {
      const localData = localStorage.getItem('cart')
      const parsedData = JSON.parse(localData)

      const product = parsedData.filter(d => d.productId === val.id)
      
      if (product[0].quantity > 0) {
        product[0].quantity = product[0].quantity - 1

        const stringifyData = JSON.stringify(parsedData)

        localStorage.setItem('cart', stringifyData)
        
        this.localData = product
      }

      if (product[0].quantity === 0) {
        const deleteObject = parsedData.filter(item => item.productId !== productId)
        const stringifyData = JSON.stringify(deleteObject)

        localStorage.setItem('cart', stringifyData)

        // if (parsedData.length === 0) {
          // this.showFooter = false
        // }
      }
      
    }
  }
}
</script>

<style scoped>
.example-item {
  height: '5em';
  width: '5em';
}
</style>