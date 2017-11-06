<template>
  <div class='hello'>
    <div class="box content is-medium">
      <div class="columns is-mobile is-centered">
        <div class="column">
          <p>Просто наберите текст на кириллице в поле ниже. Или скопируйте текст на кириллице туда. В общем, там
            должна быть кириллица :)</p>
          <div class="field">
            <p class="control">
              <textarea class="textarea" v-model="cyrillicText" v-on:input="transliteText" rows="10"></textarea>
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
                        readonly>{{latinText}}</textarea>
            </p>
          </div>
        </div>
      </div>
    </div>
    <div class="columns is-mobile is-centered has-text-centered">
      <div class="column">
        <social-sharing url="https://marazm.mi7ter.xyz/"
                        title="Конвертер кириллицы на латиницу для казахского языка"
                        description="Конвертер кириллицы на латиницу для казахского языка. С апострофами."
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
  </div>
</template>

<script>
  let transliteArray = {
    НГ: 'N\'',
    нг: 'n\'',
    А: 'A',
    а: 'a',
    Ә: 'A\'',
    ә: 'a\'',
    Ə: 'A\'',
    ə: 'a\'',
    Б: 'B',
    б: 'b',
    Д: 'D',
    д: 'd',
    Е: 'E',
    е: 'e',
    Ф: 'F',
    ф: 'f',
    Ғ: 'G\'',
    ғ: 'g\'',
    Г: 'G',
    г: 'g',
    Х: 'H',
    х: 'h',
    І: 'I',
    i: 'i',
    И: 'I\'',
    и: 'i\'',
    Й: 'I\'',
    й: 'i\'',
    H: 'H',
    h: 'h',
    Ж: 'J',
    ж: 'j',
    К: 'K',
    к: 'k',
    Л: 'L',
    л: 'l',
    М: 'M',
    м: 'm',
    Н: 'N',
    н: 'n',
    Ң: 'N\'',
    ң: 'n\'',
    О: 'O',
    о: 'o',
    Ө: 'O\'',
    ө: 'o\'',
    П: 'P',
    п: 'p',
    Қ: 'Q',
    қ: 'q',
    Р: 'R',
    р: 'r',
    Ш: 'S\'',
    ш: 's\'',
    С: 'S',
    с: 's',
    Т: 'T',
    т: 't',
    Ұ: 'U',
    ұ: 'u',
    Ү: 'U\'',
    ү: 'u\'',
    В: 'V',
    в: 'v',
    Ы: 'Y',
    ы: 'y',
    У: 'Y\'',
    у: 'y\'',
    З: 'Z',
    з: 'z',
    Ч: 'C\'',
    ч: 'c\'',
    Э: 'E',
    э: 'e',
    Щ: 's\'',
    щ: 's\'',
    ь: '',
    ъ: '',
    Я: 'I\'a',
    я: 'i\'a',
    Ю: 'I\'y\'',
    ю: 'i\'y\'',
    Ц: 'Ts',
    ц: 'ts'
  }
  export default {
    name: 'QazaqstanBestKebabCountry',
    data () {
      return {
        cyrillicText: 'Тура осындай мәселе 1920 жылдары әліпби тақырыбы көтерілгенде айтылған болатын. Жат дыбыстарды кіргіземіз бе, жоқ па, осы мәселе қаралды. Ахмет Байтұрсынұлының әліпбиге қатысты еңбектерін мұқият оқысақ, біз Алаш зиялыларының кейбір кірме дыбыстардың қазақ әліпбиінен орын алуын жақтағанын да көреміз.',
        latinText: null
      }
    },
    methods: {
      transliteText () {
        let text = this.cyrillicText
        for (let key in transliteArray) {
          text = this.replaceAll(text, key, transliteArray[key])
        }
        this.latinText = text
      },
      replaceAll (str, find, replace) {
        return str.replace(new RegExp(find, 'g'), replace)
      }
    },
    mounted () {
      this.transliteText()
    }
  }
</script>

<!-- Add 'scoped' attribute to limit CSS to this component only -->
<style>
  .link {
    cursor: pointer;
  }
</style>
