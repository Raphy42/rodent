package main

import (
	"github.com/raphy42/rodent/core"
	"github.com/raphy42/rodent/core/thread"
)

func main() {
	thread.Run(func() {
		k := core.NewFrom("./app.json")

		go k.Profile(":6060")
		err := k.Start()
		if err != nil {
			panic(err)
		}

		<- k.Wait()
	})
}