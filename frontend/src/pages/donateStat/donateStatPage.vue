<template>
  <div class="q-pa-md">
    <div class="row">
      <div class="col-12">
        <h6 class="text-center">Donation Statistics</h6>
      </div>
    </div>
    <div class="row">
      <div class="col-12">
        <iframe title="Mockup dashboard for grab - Page 1" width="350" height="800" src="https://app.powerbi.com/view?r=eyJrIjoiODAxYjc5MDItZDFhZi00OTE2LTllY2UtNzI5MzI2MGVhZTE0IiwidCI6IjlmM2VmYzI2LWVlYTMtNDVlMC05YTMzLWJkZDk2YmJjN2Y2NyIsImMiOjEwfQ%3D%3D" frameborder="0" allowFullScreen="true"></iframe>
      </div>
    </div>
  </div>
  <footer-menu />
</template>

<script>
import footerMenu from '../../components/footerMenu/footerMenu.vue'

export default {
  name: 'donateStat',
  components: {
    footerMenu
  }
}
</script>
