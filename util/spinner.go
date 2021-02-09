package util

import (
	"github.com/briandowns/spinner"
	"github.com/kyokomi/emoji/v2"
	"time"
)

func Loading(loadingFunc func() error, prefixEmoji string, loadingText string, successEmoji string, errorEmoji string) error {
	s := spinner.New(spinner.CharSets[14], 100*time.Millisecond)
	s.Prefix = Bold().Sprint(emoji.Sprint(prefixEmoji, loadingText))
	s.Start()
	err := loadingFunc()
	s.Stop()
	if err != nil {
		BoldRed().Println(emoji.Sprint(errorEmoji, loadingText, " - ", err))
		return err
	} else {
		BoldGreen().Println(emoji.Sprint(successEmoji, loadingText))
		return nil
	}
}

func DefaultLoading(loadingFunc func() error, loadingText string, prefixEmoji string) error {
	return Loading(loadingFunc, prefixEmoji, loadingText, ":thumbs_up:", ":red_exclamation_mark:")
}
