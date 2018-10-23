package system

import (
	"log"
	"sort"
	"time"
)

type Systems []ISystem

func (s *Systems) StartAll(config map[string]interface{}) error {
	systems := prioritySorter(*s)
	sort.Sort(systems)

	for _, sys := range systems {
		name := sys.Name()
		log.Printf("preInit: %s\n", name)
		if err := sys.PreInit(config[name]); err != nil {
			log.Printf("FAIL preinit: %s\n", name)
			return err
		}
	}

	for _, sys := range systems {
		log.Printf("init: %s\n", sys.Name())
		if err := sys.Init(); err != nil {
			log.Printf("FAIL init: %s\n", sys.Name())
			return err
		}
	}

	for _, sys := range systems {
		log.Printf("postInit: %s\n", sys.Name())
		if err := sys.PostInit(); err != nil {
			log.Printf("FAIL postInit: %s\n", sys.Name())
			return err
		}
	}

	return nil
}

func (s *Systems) Ticker() func(time.Time) {
	tickers := make([]func(time.Time), 0)
	for _, system := range *s {
		tickers = append(tickers, system.Ticker())
	}

	return func(delta time.Time) {
		for _, ticker := range tickers {
			ticker(delta)
		}
	}
}

type prioritySorter Systems

func (p prioritySorter) Len() int {
	return len(p)
}

func (p prioritySorter) Less(i, j int) bool {
	return p[i].Priority() > p[j].Priority()
}

func (p prioritySorter) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
