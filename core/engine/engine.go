package engine

import (
	app "github.com/raphy42/rodent/core/application"
	"github.com/raphy42/rodent/core/worker"
)

type Engine struct {
	worker   *worker.Worker
	osWorker *worker.OSWorker
	app      *app.Application
}

func New() *Engine {
	return &Engine{
		worker:   worker.NewWorker(),
		osWorker: worker.NewOSWorker(),
	}
}

func (e *Engine) Start() error {
	k := app.New(app.Resizable(true), app.WindowSize(1200, 800), app.GLVersion(4, 1))
	if err := k.Init(); err != nil {
		return err
	}
	e.app = k
	return nil
}
