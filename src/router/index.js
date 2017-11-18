import Vue from 'vue'
import Router from 'vue-router'
import QazaqstanLatinKebabCountry from '@/components/QazaqstanLatinKebabCountry'

Vue.use(Router)

const router = new Router({
  mode: 'history',
  routes: [
    {
      path: '/',
      name: 'KebabStan',
      component: QazaqstanLatinKebabCountry,
      meta: {
        title: 'Конвертер кириллицы на латиницу с апострофами'
      }
    }
  ]
})

export default router
