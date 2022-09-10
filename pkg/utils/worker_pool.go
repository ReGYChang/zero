package utils

type Job interface {
	Execute()
}

type Worker struct {
	workerPool chan chan Job
	jobChannel chan Job
	done       chan bool
}

func NewWorker(workerPool chan chan Job) *Worker {
	return &Worker{
		workerPool: workerPool,
		jobChannel: make(chan Job),
		done:       make(chan bool),
	}
}

// Start method starts the run loop for the worker, listening for a done channel in
// case we need to stop it
func (w *Worker) Start() {
	go func() {
		for {
			// register the current worker into the worker pool.
			w.workerPool <- w.jobChannel

			select {
			case job := <-w.jobChannel:
				// we have received a work request.
				job.Execute()

			case <-w.done:
				// we have received a signal to stop
				return
			}
		}
	}()
}

func (w *Worker) Stop() {
	go func() {
		w.done <- true
	}()
}

type Dispatcher struct {
	// A pool of jobs channel that are waiting to be executed
	JobQueue chan Job
	// A pool of workers channels that are registered with the dispatcher
	workerPool chan chan Job
	maxWorkers int
	maxJobs    int
}

func NewDispatcher(maxWorkers int, maxJobs int) *Dispatcher {
	jobQueue := make(chan Job, maxJobs)
	workerPool := make(chan chan Job, maxWorkers)

	return &Dispatcher{
		JobQueue:   jobQueue,
		workerPool: workerPool,
		maxJobs:    maxJobs,
		maxWorkers: maxWorkers,
	}
}

func (d *Dispatcher) Run() {
	// starting n number of workers
	for i := 0; i < d.maxWorkers; i++ {
		worker := NewWorker(d.workerPool)
		worker.Start()
	}

	go d.dispatch()
}

func (d *Dispatcher) dispatch() {
	for {
		job := <-d.JobQueue

		// a job request has been received
		go func(job Job) {
			// try to obtain a worker job channel that is available.
			// this will block until a worker is idle
			jobChannel := <-d.workerPool

			// dispatch the job to the worker job channel
			jobChannel <- job
		}(job)
	}
}
