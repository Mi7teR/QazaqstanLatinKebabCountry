package main

import (
	"flag"
	"gopkg.in/telegram-bot-api.v4"
	"log"
	"strings"
)

func main() {
	var token string
	flag.StringVar(&token, "token", "empty", "telegram bot token")
	flag.Parse()
	if token == "empty" {
		panic("token qai'da?")
	}
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
		article := tgbotapi.NewInlineQueryResultArticle(update.InlineQuery.ID, "Добавить апострофов", text)
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
		"нг": "n'",
		"а":  "a",
		"ə":  "a'",
		"б":  "b",
		"д":  "d",
		"е":  "e",
		"ф":  "f",
		"ғ":  "g'",
		"г":  "g",
		"х":  "h",
		"i":  "i",
		"и":  "i'",
		"й":  "i'",
		"h":  "h",
		"ж":  "j",
		"к":  "k",
		"л":  "l",
		"м":  "m",
		"н":  "n",
		"ң":  "n'",
		"о":  "o",
		"ө":  "o'",
		"п":  "p",
		"қ":  "q",
		"р":  "r",
		"ш":  "s'",
		"с":  "s",
		"т":  "t",
		"ұ":  "u",
		"ү":  "u'",
		"в":  "v",
		"ы":  "y",
		"у":  "y'",
		"з":  "z",
		"ч":  "c'",
		"э":  "e",
		"щ":  "s'",
		"ь":  "",
		"ъ":  "",
		"я":  "i'a",
		"ю":  "i'y'",
		"ц":  "ts",
	}
	for s, r := range replace {
		str = strings.Replace(str, s, r, -1)
		if len(r) > 0 { //uppercase
			str = strings.Replace(str, strings.ToUpper(s), strings.ToUpper(string([]rune(r)[0]))+strings.TrimLeft(r, string([]rune(r)[0])), -1)
		}
	}
	return str
}
