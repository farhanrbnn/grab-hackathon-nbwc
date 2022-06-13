<template>
  <div class="q-pa-md" style="height: 100vh;">
    <div class="row">
      <div class="col-12">
        <h4 class="text-center">Login</h4>
      </div>
    </div>
    <div class="row">
      <div class="col-12 flex flex-center">
        <img
          alt="login" 
          src="~assets/login.svg"
          style="width: 200px; height: 200px"
        >
      </div>
    </div>
    <q-form @submit="onSubmit">
    <div class="row q-mt-lg">
      <div class="col-12"> 
        <q-input color="green" v-model="userData.phone" label="Phone Number">
          <template v-slot:prepend>
            <q-icon name="phone_iphone" />
          </template>
        </q-input>
      </div>
      <div class="col-12"> 
        <q-input color="green" v-model="userData.password" type="password" label="Password">
          <template v-slot:prepend>
            <q-icon name="lock" />
          </template>
        </q-input>
      </div>
    </div>
    <div class="row q-mx-lg q-mb-lg fixed-bottom">
      <div class="col-12">
        <q-btn type="submit" style="background: #00C31E; color: white"  class="full-width" label="Submit" />
      </div>
    </div> 
    </q-form>
  </div>
</template>

<script>
import { api } from 'src/boot/axios'
import { Cookies } from 'quasar'

export default {
  name: 'loginsPage',
  data () {
    return {
      userData: {
        phone: null,
        password: null 
      }
    }
  },
  methods: {
    onSubmit () {
      const userData = this.userData

      api.post('user/signin', userData)
      .then((res) => {
        const data = res.data
        Cookies.set('user_token', data.token, {
          secure: true,
          expires: data.expires_at
        })
        this.$router.push('/home')
      })
      .catch((err) => {
        console.log(err)
        alert('error')
      })
    }
  }
}
</script>
