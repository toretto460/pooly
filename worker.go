package pooly

import "context"

// Worker is the worker struct
type Worker struct {
	jobs     chan Job
	shutdown func()
	ctx      context.Context
}

// Start will start the job listening
func (w *Worker) Start() {
	defer w.shutdown()

	for {
		select {
		case <-w.ctx.Done():
			return
		case j, ok := <-w.jobs:
			if ok {
				j()
			} else {
				return
			}
		}
	}
}
