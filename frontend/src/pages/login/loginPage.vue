<template>
  <div class="q-pa-md" style="height: 100vh;">
    <div class="row">
      <div class="col-12 flex flex-center">
        <img
          alt="Quasar logo" 
          src="~assets/gift.svg"
          style="width: 400px; height: 400px"
        >
      </div>
    </div>
    <div class="row">
      <div class="col-12">
        <h4 class="text-center q-mb-sm q-mt-sm">Soombang</h4>
      </div>
      <div class="col-12"> 
        <h6 class="text-center q-mb-sm q-mt-sm">Charity platform for those who are generous</h6>
      </div>
    </div>
    <div style="margin-top: 100px;" class="row q-my-lg q-mx-md fixed-bottom">
      <div class="col-12">
        <q-btn style="background: #00C31E; color: white" @click="goToLogin()" class="full-width" label="Login" />
      </div>
      <div class="col-12 q-mt-md">
        <q-btn outline style="color: #00C31E;" @click="goToRegister()" class="full-width" label="Sign Up to start spread kindness" />
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'loginPage',
  data () {
    return {
      test: ''
    }
  },
  methods: {
    goToRegister () {
      this.$router.push('register')
    },
    goToLogin () {
      this.$router.push('login')
    }
  }
}
</script>

<style scoped>
.footer {
    margin-right: 10px;
    margin-left: 10px;
    position: fixed;
    bottom: 0;
    width: 100%;
}
</style>
