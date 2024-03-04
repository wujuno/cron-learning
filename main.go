package main

import (
	"cron-learning/schedule"
	"fmt"
)

func main() {
	// batch 실행
	batch := schedule.BatchManager{}
	
	if err := batch.Init(); err != nil {
		fmt.Print("batch Init failed:", err)
	} else {
		fmt.Print("batch Init Success")
	}
}