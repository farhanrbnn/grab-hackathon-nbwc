
const routes = [
  {
    path: '/',
    component: () => import('layouts/MainLayout.vue'),
    children: [
      { 
        path: '', component: () => import('pages/login/loginPage.vue')
      },
      { 
        path: 'register', component: () => import('pages/register/registerPage.vue')
      },
      { 
        path: 'login', component: () => import('pages/logins/loginsPage.vue')
      },
      { 
        path: 'home', component: () => import('pages/home/homePage.vue')
      },
      { 
        path: 'changeWallet', component: () => import('pages/changeWallet/changeWalletPage.vue')
      },
      { 
        path: 'account', component: () => import('pages/account/accountPage.vue')
      },
      { 
        path: 'order', component: () => import('pages/order/orderPage.vue')
      },
      { 
        path: 'sourceFund', component: () => import('pages/sourceFund/sourceFundPage.vue')
      },
      { 
        path: 'dropOff', component: () => import('pages/dropOff/dropOffPage.vue')
      },
      { 
        path: 'merchant', component: () => import('pages/merchant/merchantPage.vue')
      },
      { 
        path: 'merchantProduct/:merchantId', component: () => import('pages/merchantProduct/merchantProductPage.vue')
      },
      { 
        path: 'orderHistory', component: () => import('pages/orderHistory/orderHistoryPage.vue')
      },
      { 
        path: 'orderHistoryDetail', component: () => import('pages/orderHistory/detailPage.vue')
      },
      { 
        path: 'checkout/:merchantId', component: () => import('pages/checkout/checkoutPage.vue')
      },
      { 
        path: 'userLocation', component: () => import('pages/userLocation/userLocationPage.vue')
      },
      { 
        path: 'activityList', component: () => import('pages/activity/activityPage.vue')
      },
      { 
        path: 'activityDetail/:orderId', component: () => import('pages/activity/detailPage.vue')
      },
      { 
        path: 'donateStat', component: () => import('pages/donateStat/donateStatPage.vue')
      },
      { 
        path: 'activityDetail', component: () => import('pages/activity/detailPage.vue')
      },
      { 
        path: 'mockDriver', component: () => import('pages/mockDriver/mainPage.vue')
      }
    ]
  },

  // Always leave this as last one,
  // but you can also remove it
  {
    path: '/:catchAll(.*)*',
    component: () => import('pages/ErrorNotFound.vue')
  }
]

export default routes
