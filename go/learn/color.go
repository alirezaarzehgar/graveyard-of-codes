package main

import "fmt"

type Color uint

const (
	Black Color = iota + 30
	Red
	Green
	Yellow
	Blue
	Magenta
	Cyan
	LightGray

	BgBlack Color = iota + 40
	BgRed
	BgGreen
	BgYellow
	BgBlue
	BgMagenta
	BgCyan
	BgLightGray

	Gray Color = iota + 90
	LightRed
	LightGreen
	LightYellow
	LightBlue
	LightMagenta
	LightCyan
	White

	BgGray Color = iota + 100
	BgLightRed
	BgLightGreen
	BgLightYellow
	BgLightBlue
	BgLightMagenta
	BgLightCyan
	BgWhite
)

func (c Color) Println(strs ...string) {
	for _, v := range strs {
		fmt.Printf("\033[%dm%s\033[0m", c, v)
	}
	fmt.Println()
}

func (c Color) Sprint(i ...any) string {
	out := ""
	for _, v := range i {
		out += fmt.Sprintf("\033[%dm%s\033[0m", c, v)
	}
	return out
}

func (c Color) Sprintln(i ...any) string {
	out := c.Sprint(i...)
	out += "\n"
	return out
}

func main() {
	fmt.Print(Red.Sprintln("Hello world"))
}
