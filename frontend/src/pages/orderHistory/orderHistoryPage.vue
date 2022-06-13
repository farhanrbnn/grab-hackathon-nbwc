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
        <h6 class="text-center">Order History</h6>
      </div>
    </div>
    <div class="row">
      <div v-for="data in orderHistory" :key="data" class="col-12 q-mb-xs">
        <q-card class="my-card">
          <q-card-section>
            <div class="row">
              <div class="col-2">
                <q-icon size="2rem" style="color: #00C31E; float: left;" name="receipt_long" />
              </div>
              <div class="col-6">
                <p>{{ data.donateTo }}</p>
                <p class="q-mb-none">{{ data.date }}</p>
              </div>
              <div class="col-4">
                <p class="text-right">
                  <b> {{ data.total }} </b>
                </p>
              </div>
            </div> 
          </q-card-section>
        </q-card>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'orderHistory',
  data () {
    return {
     orderHistory: [
        {
          donateTo: 'donation to panti asuhan ABC',
          date: '25 Mei 2022',
          total: 'Rp. 250.000'
        },
        {
          donateTo: 'donation to panti asuhan BCA',
          date: '28 Mei 2022',
          total: 'Rp. 500.000'
        },
        {
          donateTo: 'donation to panti asuhan BCA',
          date: '28 Mei 2022',
          total: 'Rp. 500.000'
        },
        {
          donateTo: 'donation to panti asuhan BCA',
          date: '28 Mei 2022',
          total: 'Rp. 500.000'
        }
      ]
    }
  }
}
</script>
