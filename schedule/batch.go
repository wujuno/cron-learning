package schedule

import "cron-learning/modules"

type BatchManager struct {
	Scheduler *modules.Scheduler
	Channel	chan interface{}
}

func (bm *BatchManager) Init() error {
	return nil
}

func (bm *BatchManager) Build() *[]modules.Task {
	return nil
}