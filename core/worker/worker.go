package worker

type IWorker interface {
	Start(stop chan struct{})
	Do(fn func())
	DoErr(fn func() error) error
	AsyncDo(fn func())
	AsyncDoErr(fn func() error) chan error
}

type Worker struct {
	queue chan func()
}

func NewWorker() *Worker {
	return &Worker{
		queue: make(chan func(), 16),
	}
}

func (w Worker) Start(stop chan struct{}) {
	for {
		select {
		case task := <- w.queue:
			task()
		case <-stop:
			return
		}
	}
}

func (w Worker) Do(task func()) {
	done := make(chan struct{})
	w.queue <- func() {
		task()
		done <- struct{}{}
	}
	<- done
}

func (w Worker) DoErr(task func() error) error {
	err := make(chan error, 1)
	w.Do(func() {
		err <- task()
	})
	return <- err
}

func (w Worker) AsyncDo(task func()) {
	w.queue <- task
}

func (w Worker) AsyncDoErr(task func() error) chan error {
	err := make(chan error, 1)
	w.AsyncDo(func() {
		err <- task()
	})
	return err
}
