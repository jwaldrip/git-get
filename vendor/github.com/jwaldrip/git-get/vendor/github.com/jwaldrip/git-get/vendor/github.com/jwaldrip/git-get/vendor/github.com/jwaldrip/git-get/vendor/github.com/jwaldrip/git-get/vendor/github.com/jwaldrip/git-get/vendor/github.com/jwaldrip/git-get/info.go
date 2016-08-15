package main

import (
	"fmt"

	git "github.com/libgit2/git2go"
)

type info struct {
	stats git.TransferProgress
	step  func()
}

func (i *info) update() int {
	if i.step == nil {
		i.setStep(i.showObjectProgress)
	}
	i.step()
	return 0
}

func (i *info) showObjectProgress() {
	if i.stats.ReceivedObjects == i.stats.TotalObjects {
		i.setStep(i.showDeltaProgress)
		return
	}
	percent := calcPercent(i.stats.ReceivedObjects, i.stats.TotalObjects)
	fmt.Printf("\rRecieving objects: %v%% (%v/%v) objects...", percent, i.stats.ReceivedObjects, i.stats.TotalObjects)
}

func (i *info) showDeltaProgress() {
	if i.stats.TotalDeltas != 0 {
		i.setStep(i.showIndexProgress)
		return
	}
	fmt.Printf("\rRecieving deltas: 100%% (%v/%v) objects...", i.stats.TotalDeltas, i.stats.TotalDeltas)
}

func (i *info) showIndexProgress() {
	percent := calcPercent(i.stats.IndexedObjects, i.stats.TotalObjects)
	fmt.Printf("\rIndexing objects: %v%% (%v/%v) objects...", percent, i.stats.ReceivedObjects, i.stats.TotalObjects)
}

func (i *info) setStep(fn func()) {
	if i.step != nil {
		fmt.Println("")
	}
	i.step = fn
}
