package thread

import (
	"runtime"

	"github.com/pkg/errors"
)

var (
	queue chan func()
)

func init() {
	runtime.LockOSThread()
}

func status() {
	if queue == nil {
		panic(errors.New("queue not started"))
	}
}

func Run(main func()) {
	queue = make(chan func())

	done := make(chan struct{})
	go func() {
		main()
		done <- struct{}{}
	}()

	for {
		select {
		case f := <-queue:
			f()
		case <-done:
			return
		}
	}
}

func Do(f func()) {
	status()
	done := make(chan struct{})
	queue <- func() {
		f()
		done <- struct{}{}
	}
	<-done
}

func DoErr(f func() error) error {
	status()
	err := make(chan error)
	queue <- func() {
		err <- f()
	}
	return <-err
}

func DoAsync(f func()) {
	status()
	queue <- f
}
