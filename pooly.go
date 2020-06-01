package pooly

import (
	"context"
	"sync"
)

// Pooly is the workers pool
type Pooly struct {
	jobs      chan Job
	ctx       context.Context
	workersWG *sync.WaitGroup
}

//Job is a job that perform the task within the worker
type Job func()

// New creates a new Pooly
func New(ctx context.Context, workers int) *Pooly {
	p := &Pooly{
		jobs:      make(chan Job),
		ctx:       ctx,
		workersWG: &sync.WaitGroup{},
	}

	for workers > 0 {
		p.startNewWorker(ctx)
		workers--
	}

	return p
}

func (p *Pooly) startNewWorker(ctx context.Context) {
	w := Worker{
		jobs: p.jobs,
		ctx:  ctx,
		shutdown: func() {
			p.workersWG.Done()
		},
	}
	p.workersWG.Add(1)
	go w.Start()
}

// RunFunc pushes the f func into the first available worker,
// if all the workers are busy the RunFunc caller will remain
// blocked until the first worker is ready to accept the job
func (p *Pooly) RunFunc(f func()) {
	p.jobs <- f
}

// Wait for the closing context to shutdown all the workers
// and then the Pooly. This func will drain the execution for
// all the long running jobs
func (p *Pooly) Wait() {
	select {
	case <-p.ctx.Done():
		close(p.jobs)
		p.workersWG.Wait()
		return
	}
}
