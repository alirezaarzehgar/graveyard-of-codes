package main

import (
	"github.com/pterm/pterm"
)

func main() {
	bars := pterm.Bars{
		pterm.Bar{
			Label: "Money",
			Value: 5,
		},
		pterm.Bar{
			Label: "Emotion",
			Value: 8,
		},
		pterm.Bar{
			Label: "Knowledge",
			Value: 3,
		},
	}

	pterm.DefaultBarChart.WithBars(bars).Render()
	pterm.DefaultBarChart.WithHorizontal().WithBars(bars).Render()

	pterm.DefaultBigText.WithLetters(
		pterm.NewLettersFromStringWithStyle("H", pterm.NewStyle(pterm.BgRed)),
		pterm.NewLettersFromStringWithStyle("E", pterm.NewStyle(pterm.BgGreen)),
		pterm.NewLettersFromStringWithStyle("L", pterm.NewStyle(pterm.BgBlue)),
		pterm.NewLettersFromStringWithStyle("L", pterm.NewStyle(pterm.FgCyan)),
		pterm.NewLettersFromStringWithStyle("O", pterm.NewStyle(pterm.FgYellow)),
	).Render()

	box := pterm.DefaultBox.WithTitle("sub1").Sprintln("how are you")
	pterm.DefaultPanel.WithPanels([][]pterm.Panel{
		{{box}, {box}},
	}).Render()

	var options []string
	v, _ := pterm.DefaultBigText.WithLetters(pterm.NewLettersFromString("Kill")).Srender()
	options = append(options, v)

	v, _ = pterm.DefaultBigText.WithLetters(pterm.NewLettersFromString("Kiss")).Srender()
	options = append(options, v)

	v, _ = pterm.DefaultBigText.WithLetters(pterm.NewLettersFromString("Fuck")).Srender()
	options = append(options, v)

	pterm.DefaultInteractiveMultiselect.WithOptions(options).Show()

	pterm.DefaultBulletList.WithItems([]pterm.BulletListItem{
		{Level: 0, Text: "programming-language"},
		{Level: 1, Text: "C"},
		{Level: 2, Text: "Kore.io"},
		{Level: 2, Text: "FRR"},
		{Level: 2, Text: "TRex"},
		{Level: 1, Text: "Python"},
		{Level: 2, Text: "Django"},
		{Level: 2, Text: "Flask"},
		{Level: 2, Text: "FastAPI"},
		{Level: 1, Text: "Go"},
		{Level: 2, Text: "Gin"},
		{Level: 2, Text: "Echo"},
		{Level: 2, Text: "PTerm"},
	}).Render()

	pterm.DefaultPanel.WithPanels(pterm.Panels{
		{{Data: "hello"}, {Data: pterm.DefaultHeader.Sprint("header")}, {Data: pterm.Red("hoyya")}},
		{{Data: "hello"}, {Data: "Ohh"}, {Data: pterm.Red("hoyya")}},
	}).WithPadding(4).Render()

	pterm.DefaultTable.WithHasHeader().WithData(pterm.TableData{
		{"firstname", "lastname", "email"},
		{"ali", "arzehgar", "aliarzehgar@gmail.com"},
		{"mohammad", "zaker", "mohammadzaker@gmail.com"},
		{"reza", "shoghi", "rezashoghi@gmail.com"},
	}).Render()

	pterm.DefaultTree.WithRoot(pterm.TreeNode{
		Text: "programming languages",
		Children: []pterm.TreeNode{
			pterm.TreeNode{
				Text: "C",
				Children: []pterm.TreeNode{
					pterm.TreeNode{Text: "Kore.io"},
					pterm.TreeNode{Text: "FRR"},
					pterm.TreeNode{Text: "TRex"},
				},
			},
			pterm.TreeNode{
				Text: "Python",
				Children: []pterm.TreeNode{
					pterm.TreeNode{Text: "Django"},
					pterm.TreeNode{Text: "Flask"},
					pterm.TreeNode{Text: "FastAPI"},
				}},
			pterm.TreeNode{
				Text: "Go",
				Children: []pterm.TreeNode{
					pterm.TreeNode{Text: "Gin"},
					pterm.TreeNode{Text: "Echo"},
					pterm.TreeNode{Text: "PTrem"},
				}},
		},
	}).Render()

	pterm.DefaultSection.Println("first")
	pterm.DefaultParagraph.Println("NewLettersFromStringWithStyle _ ")
	pterm.DefaultSection.Println("second")
}
