<template>
  <div class="row fixed-bottom shadow-2" style="background-color: #FFF;">
    <div class="col-6 text-center" @click="goToPage('home')">
      <q-icon class="q-mt-sm" size="2rem" style="color: #00C31E;" name="home" />
      <p class="q-my-none">Home</p>
    </div>
    <div class="col-6 text-center" @click="goToPage('activityList')">
      <q-icon class="q-mt-sm" size="2rem"  style="color: #00C31E;" name="schedule" />
      <p class="q-my-none" style="color:black; text-decoration: none !important;">Activity</p>
    </div>
  </div>
</template>

<script>
export default {
  name: 'footerMenu',
  methods: {
    goToPage (page) {
      this.$router.push('/' + page)
    }
  }
}
</script>
