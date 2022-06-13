<template>
  <div class="q-ma-md">
    <div class="row q-mb-md">
      <router-link to="/orderHistory">
        <div class="col-12">
          <q-icon size="2.5rem"  style="color: #00C31E; float: left;" name="arrow_back_ios" />
          <h6 class="q-my-xs" style="color: #00C31E; float: left;">Back</h6>
        </div>
      </router-link> 
    </div>
    <div class="row">
      <div class="col-12">
        <q-card class="my-card">
          <q-card-section>
           <div class="row">
             <div class="col-2">
              <q-icon size="2.5rem"  style="color: #00C31E; float: left;" name="restaurant" />
             </div>
             <div class="col-10">
               <h6 class="q-my-none">Burger Bener</h6>
               <p class="q-mt-sm q-mb-none" style="float: left;">Booking ID</p>
               <p  class="q-mt-sm q-mb-none" style="float: right;">123123131</p>
             </div>
            </div> 
          </q-card-section>
        </q-card>
      </div>
    </div>
    <div class="row">
      <div class="col-12">
        <h6 class="q-mt-md q-mb-none">Drop Off Point</h6>
      </div>
    </div>
    <div class="row">
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
            <div class="row q-mb-xl">
              <div class="col-2">
                <p class="q-my-none">100x</p>
              </div>
              <div class="col-6">
                <p class="q-my-none">Burger Bener 1</p>
              </div>
              <div class="col-4">
                <p class="q-my-none text-right">Rp. 50.000</p>
              </div>
            </div>
            <div class="row">
              <div class="col-6">
                <p class="q-my-none">Subtotal</p>
              </div>
              <div class="col-6">
                <p class="q-my-none text-right">Rp. 500.000</p>
              </div>
            </div>
            <div class="row">
              <div class="col-6">
                <p class="q-my-none">Delivery Fee</p>
              </div>
              <div class="col-6">
                <p class="q-my-none text-right">Rp. 20.000</p>
              </div>
            </div>
            <div class="row">
              <div class="col-6">
                <p class="q-my-none">Order Fee</p>
              </div>
              <div class="col-6">
                <p class="q-my-none text-right">Rp. 20.000</p>
              </div>
            </div>
            <div class="row q-mt-md">
              <div class="col-6">
                <p class="q-my-none">
                  <b>Grand Total</b>
                </p>
              </div>
              <div class="col-6">
                <p class="q-my-none text-right">
                  <b>Rp. 540.000</b>
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
  name:'detailPage'
}
</script>
