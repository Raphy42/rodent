package worker

import "runtime"

type OSWorker struct {
	worker *Worker

}

func (o OSWorker) Start(stop chan struct{}) {
	runtime.LockOSThread()
	o.worker.Start(stop)
}

func (o OSWorker) Do(fn func()) {
	o.worker.Do(fn)
}

func (o OSWorker) DoErr(fn func() error) error {
	return o.worker.DoErr(fn)
}

func (o OSWorker) AsyncDo(fn func()) {
	o.worker.AsyncDo(fn)
}

func (o OSWorker) AsyncDoErr(fn func() error) chan error {
	return o.worker.AsyncDoErr(fn)
}

func NewOSWorker() *OSWorker {
	return &OSWorker{
		worker: NewWorker(),
	}
}