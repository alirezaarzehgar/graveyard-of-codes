package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/Arafatk/glot"
)

func processCSV(csvfile string) error {
	f, err := os.Open(csvfile)
	if err != nil {
		return err
	}
	defer f.Close()

	reader := csv.NewReader(f)
	reader.FieldsPerRecord = -1
	allRecords, err := reader.ReadAll()
	if err != nil {
		return err
	}

	xP := []float64{}
	yP := []float64{}
	for _, rec := range allRecords {
		x, _ := strconv.ParseFloat(rec[0], 64)
		y, _ := strconv.ParseFloat(rec[1], 64)
		xP = append(xP, x)
		yP = append(yP, y)
	}
	points := [][]float64{xP, yP}
	fmt.Println(points)

	plot, err := glot.NewPlot(2, true, false)
	if err != nil {
		return err
	}

	plot.SetTitle("Using glot with CSV file")
	plot.SetXLabel("X-AXIS")
	plot.SetYLabel("Y_AXIS")
	allowed := []string{
		"lines", "points", "linepoints",
		"impulses", "dots", "bar",
		"steps", "fill solid", "histogram", "circle",
		"errorbars", "boxerrorbars",
		"boxes", "lp"}
	for _, style := range allowed {
		plot.AddPointGroup("Circle:", style, points)
		time.Sleep(time.Second * 3)
		plot.ResetPlot()
	}
	// plot.SavePlot("output.png")

	return nil
}

func main() {
	flag.Parse()
	for _, csvfile := range flag.Args() {
		if err := processCSV(csvfile); err != nil {
			log.Println(err)
		}
	}
}
