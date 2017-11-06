package main

import (
	"flag"
	"gopkg.in/telegram-bot-api.v2"
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
		if update.InlineQuery.Query == "" {
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
		"НГ": "N'",
		"нг": "n'",
		"А":  "A",
		"а":  "a",
		"Ә":  "A'",
		"ә":  "a'",
		"Ə":  "A'",
		"ə":  "a'",
		"Б":  "B",
		"б":  "b",
		"Д":  "D",
		"д":  "d",
		"Е":  "E",
		"е":  "e",
		"Ф":  "F",
		"ф":  "f",
		"Ғ":  "G'",
		"ғ":  "g'",
		"Г":  "G",
		"г":  "g",
		"Х":  "H",
		"х":  "h",
		"І":  "I",
		"i":  "i",
		"И":  "I'",
		"и":  "i'",
		"Й":  "I'",
		"й":  "i'",
		"H":  "H",
		"h":  "h",
		"Ж":  "J",
		"ж":  "j",
		"К":  "K",
		"к":  "k",
		"Л":  "L",
		"л":  "l",
		"М":  "M",
		"м":  "m",
		"Н":  "N",
		"н":  "n",
		"Ң":  "N'",
		"ң":  "n'",
		"О":  "O",
		"о":  "o",
		"Ө":  "O'",
		"ө":  "o'",
		"П":  "P",
		"п":  "p",
		"Қ":  "Q",
		"қ":  "q",
		"Р":  "R",
		"р":  "r",
		"Ш":  "S'",
		"ш":  "s'",
		"С":  "S",
		"с":  "s",
		"Т":  "T",
		"т":  "t",
		"Ұ":  "U",
		"ұ":  "u",
		"Ү":  "U'",
		"ү":  "u'",
		"В":  "V",
		"в":  "v",
		"Ы":  "Y",
		"ы":  "y",
		"У":  "Y'",
		"у":  "y'",
		"З":  "Z",
		"з":  "z",
		"Ч":  "C'",
		"ч":  "c'",
		"Э":  "E",
		"э":  "e",
		"Щ":  "s'",
		"щ":  "s'",
		"ь":  "",
		"ъ":  "",
		"Я":  "I'a",
		"я":  "i'a",
		"Ю":  "I'y'",
		"ю":  "i'y'",
		"Ц":  "Ts",
		"ц":  "ts",
	}
	for s, r := range replace {
		str = strings.Replace(str, s, r, -1)
	}
	return str
}
