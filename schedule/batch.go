package schedule

import "cron-learning/modules"

type BatchManager struct {
	Scheduler *modules.Scheduler
}

func (bm *BatchManager) Init() error {

	bm.Scheduler = &modules.Scheduler{
		Tasks: bm.Build(),
	}

	bm.Scheduler.Start()
	
	return nil
}

func (bm *BatchManager) Build() *[]modules.Task {
	return nil
}

