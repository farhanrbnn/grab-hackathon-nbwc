<template>
      <q-card class="my-card" style="background: #00C31E; color: #fff">
        <q-card-section>
          <div class="row"> 
            <div class="col-12">
              <div class="row">
                <div class="col-2">
                  <q-icon size="2rem"  style="color: #fff; float: left;" name="shopping_basket" />
                </div>
                <div class="col-4 q-mt-xs "><b>{{ quantity }} item</b></div>
                <div class="col-6 q-mt-xs text-right"><b>Rp. {{ subTotalPrice }}</b></div>
              </div>
            </div>
          </div>
        </q-card-section>
      </q-card>
</template>

<script>
export default {
  name: 'footerProductDisplay', 
  props: {
    localData: {
      type: null 
    }
  },
  data () {
    return {
     subTotalPrice: 0,
     quantity: 0
    }
  },
  mounted () {
    this.calculateSubtotal()
  }, 
  watch: {
    localData (val) {
      this.calculateSubtotal()
    }
  },
  methods: {
    calculateSubtotal () {
      const localData = localStorage.getItem('cart')
      const parsed = JSON.parse(localData)

      let sum = 0
      let sumQuantity = 0

      for (let i = 0; i < parsed.length; i++) {
        const quantity = parsed[i].quantity
        const price = parsed[i].productPrice
        const calculatePrice = quantity * price

        sum += calculatePrice 

        sumQuantity += parsed[i].quantity
      }

      const formated = this.formatPrice(sum)

      this.subTotalPrice = formated
      this.quantity = sumQuantity
    },
    formatPrice (value) {
      const val = (value/1).toFixed(2).replace('.', ',')
      return val.toString().replace(/\B(?=(\d{3})+(?!\d))/g, ".")
    }
  }
}
</script>
