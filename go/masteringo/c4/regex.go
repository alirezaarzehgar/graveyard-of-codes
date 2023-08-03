package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	r := regexp.MustCompile("OK")
	str := "hey guys! OK I'm sad. OK I'm not cool. OK I'm ugly. But I'm very kind :)"
	out := r.FindAllString(str, 10)
	fmt.Println(out)

	r = regexp.MustCompile(`^\w+\.pdf$`)
	str = `readme.md
document.pdf
image.png
music.mp4
manual.pdf`
	for _, line := range strings.Fields(str) {
		out := r.FindString(line)
		if out != "" {
			fmt.Println(out)
		}
	}

	str = "0x80d935144700aece3ab8fd3393adc3acab41a32893dc80"
	r = regexp.MustCompile("..")
	out = r.FindAllString(str, 100)
	fmt.Println(out)

	fmt.Printf("%v\n", strings.SplitAfter("123++321++444", "++"))
}
