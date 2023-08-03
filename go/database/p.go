package main

import "github.com/pterm/pterm"

func main() {
	opt, _ := pterm.DefaultInteractiveSelect.
		WithOptions([]string{"opt", "opt", "opt"}).
		Show()
	pterm.Info.Println(opt)
}
