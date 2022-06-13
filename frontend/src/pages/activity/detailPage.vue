<template>
  <div class="q-pa-none"> 
    <div class="row">
        <div class="row q-mb-md" style="position: absolute !important; top: 10px; left: 10px; z-index: 1;">
          <router-link to="/activityList">
            <div class="col-12">
              <q-icon size="3rem"  style="color: #00C31E; float: left;" name="arrow_circle_left" />
            </div>
          </router-link> 
        </div>
      <div class="col-12" style="position: relative;">
        <GMapMap
          style = 'height: 65vh'
          :center="center"
          :zoom="15"
          :options='{ fullscreenControl: false, mapTypeControl: false }' 
          > 
           <GMapCluster :zoomOnClick="true">
              <GMapMarker
                :position="position"
                :clickable="true"
                :draggable="true"
          />
        </GMapCluster>
        </GMapMap>
      </div>
    </div> 
  </div>
  <div class="q-pa-md">
    <div class="row">
      <div class="col-12">
        <h6 class="text-center q-my-sm" style="color: #00C31E;">{{ status }}</h6>
      </div>
    </div>
    <div class="row">
      <div class="col-12">
        <h6 class="q-mt-md q-mb-sm">Merchant</h6>
      </div>
      <div class="col-12">
        <q-card class="my-card">
          <q-card-section>
            <div class="row">
              <div class="col-2">
                <q-icon size="2rem"  style="color: #00C31E; float: left;" name="restaurant" />
              </div>
              <div class="col-10">
                <p>{{ merchant.name }}</p>
                <p class="q-mb-none q-mt-md">{{ merchant.address }}</p>
              </div>
            </div>
          </q-card-section>
        </q-card>
      </div>
    </div>
    <div class="row">
      <div class="col-12">
        <h6 class="q-mt-md q-mb-sm">Drop Off Location</h6>
      </div>
      <div class="col-12">
        <q-card class="my-card">
          <q-card-section>
            <div class="row">
              <div class="col-2">
                <q-icon size="2rem"  style="color: #FF3C3C; float: left;" name="location_on" />
              </div>
              <div class="col-10">
                <p class="q-mb-none q-mt-xs">{{ dropOff.name }}</p>
              </div>
            </div>
          </q-card-section>
        </q-card>
      </div>
    </div>
    <div v-if="this.driver !== null" class="row">
      <div class="col-12">
        <h6 class="q-mt-md q-mb-sm">Driver</h6>
      </div>
      <div class="col-12">
        <q-card class="my-card">
          <q-card-section>
            <div class="row">
              <div class="col-2">
                <q-icon size="2rem"  style="color: #00C31E; float: left;" name="face" />
              </div>
              <div class="col-10">
                <p>Ahmad Supratman</p>
                <p class="q-mb-none">B 123 ABC</p>
              </div>
            </div>
          </q-card-section>
        </q-card>
      </div>
    </div>
    <div class="row">
      <div class="col-12">
        <h6 class="q-mt-md q-mb-sm">Drop Off Person in Charge</h6>
      </div>
      <div class="col-12">
        <q-card class="my-card">
          <q-card-section>
            <div class="row">
              <div class="col-2">
                <q-icon size="2rem"  style="color: #00C31E; float: left;" name="face" />
              </div>
              <div class="col-10">
                <p class="q-mt-sm">{{ dropOff.pic }}</p>
              </div>
            </div>
          </q-card-section>
        </q-card>
      </div>
    </div>
    <div class="row q-mt-md">
      <div class="col-12"> 
        <q-btn style="background: #FF3C3C; color: white"  class="full-width" label="Cancel Order" />
      </div>
    </div>
  </div>
</template>

<script>
import { Cookies } from 'quasar'
import { api } from 'src/boot/axios'

export default {
  name:'detailPage',
  data () {
    return {
      token: null,
      order: {},
      dropOff: {},
      center: {},
      position: {},
      merchant: {},
      summary: {},
      driver: {},
      status: null,
      orderId: null 
    }
  },
  mounted () {
    this.getUserToken()
    this.fetchTransaction()
    // this.fetchOrder()
  },
  methods: {
    getUserToken () {
      const token = Cookies.get('user_token')
      this.token = token
    },
    async fetchTransaction () {
      const config = {
        headers: { Authorization: `Bearer ${this.token}` }
      }

      await api.get(`/transaction/${this.$route.params.orderId}`, config)
      .then((res) => {
        this.status = res.data.status
        this.orderId = res.data.order_id

      })
      .catch((err) => {
        console.log(err)
      })


      await api.get(`/order/${this.orderId}`, config)
      .then((res) => {
        const payload = res.data

        this.position = {
          lat: payload.drop_off_location.coordinate.latitude,
          lng: payload.drop_off_location.coordinate.longitude 
        }

        this.center = {
          lat: payload.drop_off_location.coordinate.latitude,
          lng: payload.drop_off_location.coordinate.longitude 
        }

        this.merchant = payload.manifest.merchant

        this.dropOff.name = payload.drop_off_location.name
        this.dropOff.pic = payload.drop_off_location.pic

        this.summary.adminFee = payload.manifest.admin_fee
        this.summary.deliveryFee = payload.manifest.delivery_fee
        this.summary.totalPrice = payload.manifest.total_price
        this.summary.picFee = payload.manifest.pic_fee

        this.driver = payload.driver

      })
      .catch((err) => {
        console.log(err)
      })
    }
    // fetchOrder () {
    //   const config = {
    //     headers: { Authorization: `Bearer ${this.token}` }
    //   }

    // }
  }
  
}
</script>
