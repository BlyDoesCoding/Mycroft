package main

import (
	"strings"


	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"

	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Mycroft")

	var characters = []string{
		"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M",
		"N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z",
		"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m",
		"n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z",
		"Ä", "Ö", "Ü", "ä", "ö", "ü", "ß",
		"0", "1", "2", "3", "4", "5", "6", "7", "8", "9",
		"!", "\"", "#", "$", "%", "&", "'", "(", ")", "*", "+", ",", "-", ".", "/", ":", ";", "<", "=", ">", "?", "@", "[", "\\",
		"]", "^", "_", "`", "{", "|", "}", "~",
	}
	charToPosition := make(map[string]int)

	for i, char := range characters {
		charToPosition[char] = i + 1
	}



	page_1 := container.New(
		layout.NewVBoxLayout(),
		layout.NewSpacer(),

		widget.NewButton("Translate Text to Jelly", func() {
			tmp := ""
			for _, char := range w.Clipboard().Content() {
				for key, value := range charToPosition {
					if string(char) == " " {

						tmp = tmp + "/"
						break
					}
					if string(char) == key {
						for i := 0; i < value; i++ {
							tmp = tmp + "🪼"
						}
						tmp = tmp + " "
						break
					}

				}
			}

			w.Clipboard().SetContent(tmp)

			tmp = ""
		}),
		layout.NewSpacer(),

		widget.NewButton("Translate Jelly to Text", func() {
			tmp := ""

			letters := strings.Split(w.Clipboard().Content(), " ")

			for _, letter := range letters {

				for key, value := range charToPosition {
					if strings.Count(letter, "🪼") == value {
						if strings.Contains(letter, "/") {
							tmp = tmp + " "
						}
						tmp = tmp + key

					}

				}

			}

			w.Clipboard().SetContent(tmp)

			tmp = ""
		}),
		layout.NewSpacer(),
	)

	w.SetContent(page_1)

	w.ShowAndRun()
}
