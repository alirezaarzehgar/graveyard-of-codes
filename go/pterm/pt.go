package main

import (
	"strings"
	"time"

	"github.com/pterm/pterm"
)

func main() {
	pterm.DefaultHeader.WithBackgroundStyle(pterm.NewStyle(pterm.BgRed)).
		WithFullWidth().Printfln("Somebody give me ohhh")

	pterm.EnableDebugMessages()
	pterm.DefaultHeader.Println("header")

	pterm.Debug.Println("ddd")
	pterm.Error.Println("Hello, World")
	pterm.Warning.Println("Hi")
	pterm.Error.WithShowLineNumber().Println("hello")
	pterm.Success.Println("ffff")

	pterm.DefaultHeader.WithFullWidth().Println("somebody giveme hoyyya")

	pterm.DefaultTable.WithHasHeader().WithData(pterm.TableData{
		{"Firstname", "Lastname", "Email"},
		{"Alireza", "Arzehgar", "alirezaarzehgar82@gmail.com"},
		{"Sina", "Soleymani", "sinasoleymani78@gmail.com"},
		{"Zahra", "Hoseini", "zahrahoseini@gmail.com"},
		{"Ninaz", "Asaran", "ninazasaran42@gmail.com"},
		{"Mehdi", "Vahedi", "mehdivahedi80@gmail.com"},
	}).Render()

	fakeInstallList := strings.Split(`at-spi2-core-2.48.3-1  avahi-1:0.8+r22+gfd482a7-1  babl-0.1.106-1  bash-5.1.016-4  binutils-2.40-6  bmake-20230601-1  btrfs-progs-6.3.2-1  busybox-1.36.1-1  c-ares-1.19.1-1  ca-certificates-mozilla-3.90-1  chromium-114.0.5735.133-1  code-1.79.2-1  containerd-1.7.2-1  curl-8.1.2-1  cython-0.29.35-2  dav1d-1.2.1-1  dbus-1.14.8-1  dbus-glib-0.112-3  dbus-python-1.3.2-2  deepin-qt-dbus-factory-6.0.0-3  deepin-qt5integration-5.6.6-3  deepin-qt5platform-plugins-5.6.5-4  diffutils-3.10-1  djvulibre-3.5.28-5  docker-1:24.0.2-1  dtkcommon-5.6.9-1  dtkgui-1:5.6.8-2  dtkwidget-5.6.7-2  dtkwm-2.0.12-19  duktape-2.7.0-6  edk2-aarch64-202305-1  edk2-arm-202305-1  edk2-ovmf-202305-1  electron-1:25-1  electron22-22.3.14-1  electron25-25.2.0-2  eog-44.2-1  epiphany-44.3-1  evince-44.2-1  evolution-data-server-3.48.3-1  exempi-2.6.3-2  exiv2-0.28.0-2  ffmpeg-2:6.0-8  ffmpeg4.4-4.4.4-1  flac-1.4.3-1  fluidsynth-2.3.3-1  freeimage-3.18.0-20  frei0r-plugins-2.3.0-1  fribidi-1.0.13-2  gc-8.2.4-1  gegl-0.4.44-3  ghostscript-10.01.2-1  git-2.41.0-1  gjs-2:1.76.2-1`, "  ")

	p, _ := pterm.DefaultProgressbar.WithTotal(len(fakeInstallList)).WithTitle("Downloading stuff").Start()
	for i := 0; i < p.Total; i++ {
		p.UpdateTitle("Downloading " + fakeInstallList[i])
		pterm.Printfln(p.Title)
		p.Increment()
		time.Sleep(time.Second / 10)
	}
}
