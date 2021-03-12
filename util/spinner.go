package util

import (
	"github.com/briandowns/spinner"
	"github.com/kyokomi/emoji/v2"
	"time"
)

func Loading(loadingFunc func(sp *spinner.Spinner) error, prefixEmoji string, loadingText string, successEmoji string, errorEmoji string) (*spinner.Spinner, error) {
	s := spinner.New(spinner.CharSets[14], 100*time.Millisecond)
	s.Prefix = Bold().Sprint(emoji.Sprint(prefixEmoji, loadingText))
	s.Start()
	err := loadingFunc(s)
	s.Stop()
	if err != nil {
		BoldRed().Println(emoji.Sprint(errorEmoji, loadingText, " - ", err))
		return s, err
	} else {
		BoldGreen().Println(emoji.Sprint(successEmoji, loadingText))
		return s, nil
	}
}

func DefaultLoading(loadingFunc func(sp *spinner.Spinner) error, loadingText string, prefixEmoji string) error {
	_, err := Loading(loadingFunc, prefixEmoji, loadingText, ":thumbs_up:", ":red_exclamation_mark:")
	return err
}
