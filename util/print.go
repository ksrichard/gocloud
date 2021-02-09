package util

import (
	"errors"
	"github.com/kyokomi/emoji/v2"
)

func EmojiPrintln(emojiPrefix string, text string) {
	emoji.Println(
		emojiPrefix,
		Bold().Sprint(text),
	)
}

func BoldError(err string) error {
	return errors.New(BoldRed().Sprint(err))
}
