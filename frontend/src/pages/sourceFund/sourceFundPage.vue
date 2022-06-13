<template>
  <div class="q-pa-md"> 
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
        <h6 class="text-center">Select Source of Fund</h6>
      </div>
    </div>
    <q-form @submit="onSubmit">
    <div class="row">
      <div class="col-12">
          <q-select color="green" v-model="fund" :options="options"  emit-value map-options label="source of fund" />
      </div>
    </div>
    <div class="row">
      <div class="col-12">
        <q-input v-model="amount" color="green" label="Amount you want to donate" />
      </div>
    </div>
    <div v-if="fund == 'kitabisa' " class="row q-mt-xl"> 
      <div class="col-12">
        <q-file color="green" v-model="image" label="image" />
      </div> 
    </div>
    <div class="row">
      <div class="col-12 q-mt-lg">
        <q-btn style="background: #00C31E; color: white" label="Submit" type="submit" class="full-width" />
      </div>
    </div>
    </q-form>
  </div>
</template>

<script>
import { ref } from 'vue'
import { api } from 'src/boot/axios'
import { Cookies } from 'quasar'

export default {
  name: 'sourceFund',
  data () {
    return {
      fund: ref(null),
      image: ref(null),
      options: [
        {
          label: 'KitaBisa',
          value: 'kitabisa'
        },
        {
          label: 'Personal',
          value: 'personal'
        } 
      ],
      funding_source: {},
      amount: null
    }
  },
  mounted () {
    this.getUserToken()
  },
  methods:{
    onSubmit () {
      const config = {
        headers: { Authorization: `Bearer ${this.token}` }
      }

      api.get('/user-wallets', config)
      .then((res) => {
        const payload = res.data[0]
        const fundingSource = {}

        fundingSource.user_wallet_id = payload.id
        fundingSource.source = this.fund 
        fundingSource.amount = parseInt(this.amount)

        const stringifyData = JSON.stringify(fundingSource)
        localStorage.setItem('funding_source', stringifyData)
        this.$router.push('/dropOff')
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
