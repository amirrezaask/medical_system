package i18n

import (
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func init() {
	_ = message.SetString(language.English, "hello", "world")
}
