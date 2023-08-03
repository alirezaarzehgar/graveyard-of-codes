package main

import (
	"encoding/json"
	"log"
	"math"
	"math/rand"
	"os"
	"time"

	"github.com/pterm/pterm"
)

const (
	MAX_RESPONSE = 6
	DATA_FILE    = "/etc/rsug.json"
)

func main() {
	colors := []func(...any) string{
		pterm.Green, pterm.Gray,
		pterm.LightCyan, pterm.Blue,
		pterm.Yellow, pterm.Red,
	}

	var mainMenu map[string][]string
	rawData, err := os.ReadFile(DATA_FILE)
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal(rawData, &mainMenu)

	var titles []string
	for k := range mainMenu {
		titles = append(titles, k)
	}
	titles = append(titles, "Exit")

	so := ""
	for {
		so, _ = pterm.DefaultInteractiveSelect.
			WithDefaultOption(so).
			WithOptions(titles).
			WithFilter(false).
			Show()

		if so == "Exit" {
			os.Exit(0)
		}

		r := len(mainMenu[so])
		opts := make([]string, r)
		copy(opts, mainMenu[so])

		rnd := rand.New(rand.NewSource(time.Now().UnixMicro()))
		dur := int(math.Min(MAX_RESPONSE, float64(r)))

		for i := 1; i <= dur; i++ {
			index := rnd.Intn(len(opts))
			f := colors[(i-1)%len(colors)]

			pterm.Println(i, "-", f(opts[index]))
			opts = append(opts[:index], opts[index+1:]...)
		}
	}
}
