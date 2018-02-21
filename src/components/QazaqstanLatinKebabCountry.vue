<template>
  <div class='hello'>
    <div class="box content is-medium">
      <div class="columns is-mobile is-centered">
        <div class="column">
          <p>Просто наберите текст на кириллице в поле ниже. Или скопируйте текст на кириллице туда. В общем, там
            должна быть кириллица :)</p>
          <div class="field">
            <p class="control">
              <textarea class="textarea" v-model="cyrillicText" rows="10"></textarea>
            </p>
          </div>
        </div>
      </div>
      <div class="columns is-mobile is-centered" v-if="latinText">
        <div class="column">
          <h4 class="title is-4">Na'ti'je:</h4>
          <div class="field">
            <p class="control">
              <textarea class="textarea" v-on:focus="$event.currentTarget.select()" rows="10"
                        readonly>{{ latinText }}</textarea>
            </p>
          </div>
        </div>
      </div>
    </div>
    <div class="columns is-mobile is-centered has-text-centered">
      <div class="column">
        <social-sharing url="https://marazm.mi7ter.xyz/"
                        title="Конвертер кириллицы на латиницу для казахского языка"
                        description="Конвертер кириллицы на латиницу для казахского языка. С диакритиками."
                        quote="Конвертер кириллицы на латиницу для казахского языка."
                        hashtags="qazaq, latin, cyrillic, converter"
                        inline-template>
          <div class="tags"> Поделиться в
            <network network="facebook" class="link tag">
              <i class="fa fa-facebook"></i> Facebook
            </network>
            <network network="twitter" class="link tag">
              <i class="fa fa-twitter"></i> Twitter
            </network>
            <network network="vk" class="link tag">
              <i class="fa fa-vk"></i> VK
            </network>
          </div>
        </social-sharing>
      </div>
    </div>
    <div class="columns">
      <div class="column">
        <gh-btns-star slug="Mi7teR/QazaqstanLatinKebabCountry" show-count></gh-btns-star>
      </div>
    </div>
  </div>
</template>

<script>
  let transliteArray = {
    'сх': 'skh',
    'а': 'a',
    'ә': 'á',
    'ə': 'á',
    'б': 'b',
    'в': 'v',
    'г': 'g',
    'ғ': 'ǵ',
    'д': 'd',
    'е': 'e',
    'ё': 'ıo',
    'ж': 'j',
    'з': 'z',
    'и': 'ı',
    'й': 'ı',
    'к': 'k',
    'қ': 'q',
    'л': 'l',
    'м': 'm',
    'н': 'n',
    'ң': 'ń',
    'о': 'o',
    'ө': 'ó',
    'п': 'p',
    'р': 'r',
    'с': 's',
    'т': 't',
    'у': 'ý',
    'ұ': 'u',
    'ү': 'ú',
    'ф': 'f',
    'х': 'h',
    'һ': 'h',
    'ц': 'ts',
    'ч': 'ch',
    'ш': 'sh',
    'щ': 'sh',
    'ъ': '',
    'ы': 'y',
    'і': 'i',
    'ь': '',
    'э': 'e',
    'ю': 'ıý',
    'я': 'ıá'
  }
  export default {
    name: 'QazaqstanBestKebabCountry',
    data () {
      return {
        cyrillicText: 'Тура осындай мәселе 1920 жылдары әліпби тақырыбы көтерілгенде айтылған болатын. Жат дыбыстарды кіргіземіз бе, жоқ па, осы мәселе қаралды. Ахмет Байтұрсынұлының әліпбиге қатысты еңбектерін мұқият оқысақ, біз Алаш зиялыларының кейбір кірме дыбыстардың қазақ әліпбиінен орын алуын жақтағанын да көреміз.'
      }
    },
    computed: {
      latinText () {
        let text = this.cyrillicText
        return this.latinize(text)
      }
    },
    methods: {
      latinize (text) {
        for (let key in transliteArray) {
          let character = transliteArray[key]
          if (character.length > 0) {
            text = this.replaceAll(text, key.toUpperCase(), character.charAt(0).toUpperCase() + character.slice(1))
          }
          text = this.replaceAll(text, key, character)
        }
        return text
      },
      replaceAll (str, find, replace) {
        return str.replace(new RegExp(find, 'g'), replace)
      }
    }
  }
</script>

<style>
  .link {
    cursor: pointer;
  }
</style>
