package modules

import (
	cron "github.com/robfig/cron/v3"
)

type Task struct {
	Name     string
	Schedule string
	Handler  func()
	Job      cron.FuncJob
	EntryId  cron.EntryID
	Error    error
}

type Scheduler struct {
	Tasks *[]Task
}