// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import App from './App'
import router from './router'
import VueBlu from 'vue-blu'
import 'vue-blu/dist/css/vue-blu.min.css'
import VueAnalytics from 'vue-ua'
import SocialSharing from 'vue-social-sharing'
import VueGitHubButtons from 'vue-github-buttons'
import 'vue-github-buttons/dist/vue-github-buttons.css'

Vue.use(VueGitHubButtons)
Vue.use(VueAnalytics, {
  appName: 'marazm', // Mandatory
  appVersion: '0.1', // Mandatory
  trackingId: 'UA-80979277-2', // Mandatory
  debug: false, // Whether or not display console logs debugs (optional)
  vueRouter: router
})
Vue.use(VueBlu)
Vue.use(SocialSharing)

Vue.config.productionTip = false

/* eslint-disable no-new */
new Vue({
  el: '#app',
  router,
  template: '<App/>',
  components: { App }
})
