package main

import (
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/pterm/pterm"
)

func main() {
	spin, _ := pterm.DefaultSpinner.
		WithText("I'm so fucked up. Please w8").
		WithSequence("x", "-", "+", "^", "1").
		Start()
	time.Sleep(time.Second * 10)
	if rand.Intn(2) == 0 {
		spin.Fail()
	} else {
		spin.Success()
	}
	spin.Stop()

	fakeList := strings.Split("qemu-ui-sdl-8.0.2-1  qemu-ui-spice-app-8.0.2-1  qemu-ui-spice-core-8.0.2-1  qemu-user-8.0.2-1  qemu-vhost-user-gpu-8.0.2-1  qt5-base-5.15.10+kde+r129-3  qt5-declarative-5.15.10+kde+r26-1  qt5-graphicaleffects-5.15.10-1  qt5-imageformats-5.15.10+kde+r9-1  qt5-location-5.15.10+kde+r4-1  qt5-multimedia-5.15.10+kde+r3-1  qt5-quickcontrols-5.15.10-1  qt5-quickcontrols2-5.15.10+kde+r6-1  qt5-speech-5.15.10+kde+r1-1  qt5-svg-5.15.10+kde+r8-1  qt5-tools-5.15.10+kde+r3-1  qt5-translations-5.15.10-1  qt5-wayland-5.15.10+kde+r51-1", "  ")
	bar, _ := pterm.DefaultProgressbar.
		WithTotal(len(fakeList)).
		Start()
	for i := 0; i < len(fakeList); i++ {
		bar.UpdateTitle("Downloading " + fakeList[i])
		pterm.Success.Println(bar.Title)
		bar.Increment()
		time.Sleep(time.Second / 5)
	}
	bar.Stop()

	area, _ := pterm.DefaultArea.
		WithRemoveWhenDone().
		Start()
	for i := 0; i < 3; i++ {
		str, _ := pterm.DefaultBigText.
			WithLetters(pterm.NewLettersFromString(time.Now().Format("15:04:05"))).
			Srender()
		str = pterm.DefaultCenter.Sprint(str)
		area.Update(str)
		time.Sleep(time.Second)
	}
	area.Stop()

	area.WithFullscreen().WithCenter().Start()
	for i := 0; i < 3; i++ {
		area.Update(strconv.Itoa(i))
		time.Sleep(time.Second)
	}
	area.Stop()
	area.Clear()

	str, _ := pterm.DefaultBigText.
		WithLetters(pterm.NewLettersFromString("Hello, world")).
		Srender()
	pterm.Print(str)

	bar, _ = pterm.DefaultProgressbar.
		WithTotal(100).
		WithTitle("hi hi").
		WithBarFiller(" ").
		WithBarCharacter("#").
		WithRemoveWhenDone().
		Start()
	for i := 0; i < 100; i++ {
		bar.Increment()
		time.Sleep(time.Second / 40)
	}
	bar.Stop()
}
