package tint

// Stylize applies a style to a string.
func Stylize(str string, style Style) string {
	return style.Render(str)
}

// Colorize applies a given color code to a string.
func Colorize(str string, code int) string {
	return Stylize(str, Style{Color: code})
}

// ColorizeBackground applies a given color code to a strings background.
func ColorizeBackground(str string, code int) string {
	return Stylize(str, Style{Background: code})
}
