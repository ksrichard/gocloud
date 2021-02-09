package util

import "github.com/fatih/color"

func BoldRed() *color.Color {
	red := color.New(color.FgRed)
	boldRed := red.Add(color.Bold)
	return boldRed
}

func BoldGreen() *color.Color {
	red := color.New(color.FgGreen)
	boldRed := red.Add(color.Bold)
	return boldRed
}

func Bold() *color.Color {
	black := color.New(color.FgBlue)
	boldBlack := black.Add(color.Bold)
	return boldBlack
}
