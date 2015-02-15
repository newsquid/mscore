package workers

import (
	"time"
)

/*
Worker a worker that can performe a routine of work
*/
type Worker interface {
	Routine()
}

/*
Handler handles the worker
Interval time between each routine of work is performed
Worker the worker to handle
*/
type Handler struct {
	Interval time.Duration
	Worker   Worker
}

/*
Start starts the handler, starting the worker in a new goroutine
the routine will be performed on start and then sleep for the defined Interval time
*/
func (h *Handler) Start() {
	go h.run()
}

/*
run runs the worker code
*/
func (h *Handler) run() {
	for {
		//perform the work and then sleep
		h.Worker.Routine()
		time.Sleep(h.Interval)
	}
}
