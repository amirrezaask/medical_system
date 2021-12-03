package i18n

import (
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func Trans(lang string, key string, format ...interface{}) string {
	t := language.MustParse(lang)
	return message.NewPrinter(t).Sprintf(key, format...)
}
