package main

import (
	"github.com/pterm/pterm"
)

func main() {
	r, _ := pterm.DefaultInteractiveConfirm.
		WithDefaultValue(true).
		WithDefaultText("Fuck you").
		Show("prompt")

	pterm.Printfln("You answered: %s", r)

	sr, _ := pterm.DefaultInteractiveTextInput.
		WithMultiLine().
		Show("Enter your fucking message")
	pterm.Println(sr)

	ss := pterm.DefaultInteractiveSelect.
		WithOptions([]string{"Login", "Register", "Logout"}).
		WithDefaultOption("Register").
		WithMaxHeight(3)
	ss.Selector = "->"
	ssr, _ := ss.Show("ohh")
	pterm.Println(ssr)

	pt, _ := pterm.DefaultInteractiveMultiselect.
		WithOptions([]string{"a", "b", "c", "d"}).
		Show()
	for _, v := range pt {
		for i := 0; i < 3; i++ {
			pterm.Print(v)
		}
		pterm.Println()
	}
}
