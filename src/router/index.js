import Vue from 'vue'
import Router from 'vue-router'
import QazaqstanBestKebabCountry from '@/components/QazaqstanBestKebabCountry'

Vue.use(Router)

const router = new Router({
  mode: 'history',
  routes: [
    {
      path: '/',
      name: 'KebabStan',
      component: QazaqstanBestKebabCountry,
      meta: {
        title: 'Конвертер кириллицы на латиницу с апострофами'
      }
    }
  ]
})

export default router
