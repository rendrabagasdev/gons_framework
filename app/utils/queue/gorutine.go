package queue

import (
	"gons/app/contracts"
	"log/slog"
)

type GoroutineDriver struct {
	jobChanel chan contracts.Job
}

func NewGoroutineDriver(bufferSize int) *GoroutineDriver {
	driver := &GoroutineDriver{
		jobChanel: make(chan contracts.Job, bufferSize),
	}

	go driver.worker()

	return driver
}

func (g *GoroutineDriver) Push(job contracts.Job) {
	g.jobChanel <- job
}

func (g *GoroutineDriver) worker() {
	slog.Info("Queue Worker is running... ")
	for job := range g.jobChanel {
		job()
	}
}
