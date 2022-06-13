<template>
  <div class="q-pa-md" style="height: 100vh;">
    <div class="row">
      <div class="col-12">
        <h4 class="text-center">Start to spread kindness</h4>
      </div>
    </div>
    <div class="row">
      <div class="col-12 flex flex-center">
        <img
          alt="register" 
          src="~assets/register.svg"
          style="width: 200px; height: 200px"
        >
      </div>
    </div>
    <q-form @submit="onSubmit">
    <div class="row q-mt-lg">
      <div class="col-12">
        <q-input color="green" v-model="userData.username" label="Username">
          <template v-slot:prepend>
            <q-icon name="person" />
          </template>
        </q-input>
        </div>
      <div class="col-12">
        <q-input color="green" v-model="userData.email" label="email">
          <template v-slot:prepend>
            <q-icon name="email" />
          </template>
        </q-input>
        </div> 
      <div class="col-12">
        <q-input color="green" v-model="userData.phone" label="Phone Number">
          <template v-slot:prepend>
            <q-icon name="phone_iphone" />
          </template>
        </q-input>
      </div>
      <div class="col-12">
        <q-input color="green" v-model="userData.name" label="Full Name">
          <template v-slot:prepend>
            <q-icon name="person" />
          </template>
        </q-input>
      </div>
      <div class="col-12">
        <q-input color="green" type="password" v-model="userData.password" label="Password">
          <template v-slot:prepend>
            <q-icon name="lock" />
          </template>
        </q-input>
      </div>
    </div>
    <div class="row q-mx-lg q-mb-lg fixed-bottom">
      <div class="col-12">
        <q-btn type="submit" style="background: #00C31E; color: white" class="full-width" label="Submit" />
      </div>
    </div>
    </q-form>
  </div>
</template>

<script>
import { api } from 'src/boot/axios'
import { useQuasar } from 'quasar'


export default {
  name: 'registerPage',
  data () {
    return {
      userData: {
        username: null,
        password: null,
        name: null,
        phone: null,
        email:  null 
      }
    }
  },
  methods: {
    onSubmit () {
      const userData = this.userData

      api.post('user/create', userData)
      .then((res) => {
        alert('success register')
        this.$router.push('/login')
      })
      .catch((err) => {
        console.log(err)
        alert('error:', err.responseText)
      })
    }
  }
}
</script>
