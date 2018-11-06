package logic

import "time"

type Logic interface {
	Tick(delta time.Time)

}