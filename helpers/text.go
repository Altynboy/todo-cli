package helpers

func TrancuteOrWrap(text string, width int) string {
	if len(text) <= width {
		return text
	}

	var wrapped string
	for len(text) > width {
		wrapped += text[:width] + "\n"
		text = text[width:]
	}
	wrapped += text

	return wrapped
}
