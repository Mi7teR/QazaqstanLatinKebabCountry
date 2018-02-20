package main

import (
  "flag"
  "gopkg.in/telegram-bot-api.v4"
  "log"
  "strings"
  "os"
  "os/signal"
)

func main() {
  var token string
  flag.StringVar(&token, "token", "empty", "telegram bot token")
  flag.Parse()
  if token == "empty" {
    panic("token qai'da?")
  }
  signals := make(chan os.Signal, 1)
  signal.Notify(signals)
  bot, err := tgbotapi.NewBotAPI(token)
  if err != nil {
    log.Panic(err)
  }

  u := tgbotapi.NewUpdate(0)
  u.Timeout = 60

  updates, err := bot.GetUpdatesChan(u)

  for update := range updates {
    if update.Message != nil {
      var messageText string
      switch update.Message.Text {
      case "/start":
        messageText = `Бот конвертирует кириллицу в латиницу по стандартам нового казахского алфавита на основе латиницы.
Бота не нужно никуда добавлять или писать ему что-нибудь здесь - просто наберите в начале сообщения ` + "\n`@ApostrofBot что-нибудь`" + `
и бот автоматически переведет ` + "`что-нибудь`" + `
в латиницу казахского языка, вам останется только нажать на кнопку с переводом.
Или просто напишите сообщение в этот чат`
      default:
        messageText = latinize(update.Message.Text)
      }
      message := tgbotapi.NewMessage(update.Message.Chat.ID, messageText)
      message.ParseMode = "Markdown"
      bot.Send(message)
      continue
    }

    if update.InlineQuery == nil {
      continue
    }

    text := latinize(update.InlineQuery.Query)
    article := tgbotapi.NewInlineQueryResultArticle(update.InlineQuery.ID, "Добавить Европы", text)
    article.Description = text

    inlineConf := tgbotapi.InlineConfig{
      InlineQueryID: update.InlineQuery.ID,
      IsPersonal:    true,
      CacheTime:     0,
      Results:       []interface{}{article},
    }

    if _, err := bot.AnswerInlineQuery(inlineConf); err != nil {
      log.Println(err)
    }
  }
}

func latinize(str string) string {
  replace := map[string]string{
    "а": "a",
    "ә": "á",
    "ə": "á́́́́",
    "б": "b",
    "в": "v",
    "г": "g",
    "ғ": "ǵ",
    "д": "d",
    "е": "e",
    "ё": "ıo",
    "ж": "j",
    "з": "z",
    "и": "ı",
    "й": "ı",
    "к": "k",
    "қ": "q",
    "л": "l",
    "м": "m",
    "н": "n",
    "ң": "ń",
    "о": "o",
    "ө": "ó",
    "п": "p",
    "р": "r",
    "с": "s",
    "т": "t",
    "у": "ý",
    "ұ": "u",
    "ү": "ú",
    "ф": "f",
    "х": "h",
    "һ": "h",
    "ц": "ts",
    "ч": "ch",
    "ш": "sh",
    "щ": "sh",
    "ъ": "",
    "ы": "y",
    "і": "i",
    "ь": "",
    "э": "e",
    "ю": "ıý",
    "я": "ıá",
  }
  for s, r := range replace {
    str = strings.Replace(str, s, r, -1)
    if len(r) > 0 { //uppercase
      str = strings.Replace(str, strings.ToUpper(s), strings.ToUpper(string([]rune(r)[0]))+strings.TrimLeft(r, string([]rune(r)[0])), -1)
    }
  }
  return str
}
