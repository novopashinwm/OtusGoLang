package main

import "fmt"

type Bar struct {
	percent int64  // progress percentage
	cur     int64  // current progress
	total   int64  // total value for progress
	rate    string // the actual progress bar to be printed
	graph   string // the fill value for progress bar
}

func (progress *Bar) NewOption(start, total int64) {
	progress.cur = start
	progress.total = total
	if progress.graph == "" {
		progress.graph = "█"
	}
	progress.percent = progress.getPercent()
	for i := 0; i < int(progress.percent); i += 2 {
		progress.rate += progress.graph // initial progress position
	}
}

func (progress *Bar) getPercent() int64 {
	return int64(float32(progress.cur) / float32(progress.total) * 100)
}

func (progress *Bar) NewOptionWithGraph(start, total int64, graph string) {
	progress.graph = graph
	progress.NewOption(start, total)
}

func (progress *Bar) Play(cur int64) {
	progress.cur = cur
	last := progress.percent
	progress.percent = progress.getPercent()
	if progress.percent != last && progress.percent%2 == 0 {
		progress.rate += progress.graph
	}
	fmt.Printf("\r[%-50s]%3d%% %8d/%d", progress.rate, progress.percent, progress.cur, progress.total)
}

func (progress *Bar) Finish() {
	fmt.Println()
}
