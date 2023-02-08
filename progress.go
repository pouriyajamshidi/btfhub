package main

import (
	"fmt"
	"time"

	"github.com/dustin/go-humanize"
)

type progressCounter struct {
	Op   string
	Size uint64
	Name string

	written uint64

	lastReport time.Time
}

func (wc *progressCounter) Write(p []byte) (int, error) {
	n := len(p)
	wc.written += uint64(n)

	if wc.written == wc.Size || time.Since(wc.lastReport) > 10*time.Second {
		wc.printProgress()
	}
	return n, nil
}

func (wc *progressCounter) printProgress() {
	// Clear the line by using a character return to go back to the start and remove
	// the remaining characters by filling it with spaces
	//fmt.Printf("\r%s", strings.Repeat(" ", 35))
	//fmt.Printf("%s\n", time.Since(wc.lastReport))

	// Return again and print current status of download
	pct := uint64((float64(wc.written) / float64(wc.Size)) * 100)
	fmt.Printf("%sing %s: %s / %s - %d%% complete\n", wc.Op, wc.Name, humanize.Bytes(wc.written), humanize.Bytes(wc.Size), pct)
	wc.lastReport = time.Now()
}
