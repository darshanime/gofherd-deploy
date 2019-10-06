package main

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"os"
	"time"

	gf "github.com/darshanime/gofherd"
)

var herd *gf.Gofherd

func LoadWork(herd *gf.Gofherd) {
	sites, err := os.Open("sites.csv")
	if err != nil {
		panic("could not find sites.csv. Please run from inside examples dir")
	}
	sitesReader := csv.NewReader(sites)
	site, err := sitesReader.ReadAll()
	counter := 0
	for {
		w := gf.Work{ID: fmt.Sprintf("%d", counter), Body: site[counter%len(site)][0]}
		herd.SendWork(w)
		counter++
	}
	// herd.CloseInputChan()
}

func ProcessWork(w *gf.Work) gf.Status {
	start := time.Now()
	http.Get(w.Body.(string))
	w.SetResult(time.Since(start))
	return gf.Success
}

func ReviewOutput(outputChan <-chan gf.Work) {
	for range outputChan {
	}
}

func main() {
	herd = gf.New(ProcessWork)
	herd.SetHerdSize(5)
	herd.SetMaxRetries(0)
	herd.SetAddr("0.0.0.0:5555")
	go LoadWork(herd)
	herd.Start()
	ReviewOutput(herd.OutputChan())
}
