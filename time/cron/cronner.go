package cron

import (
	"sync"
	"time"

	t "github.com/finnan444/utils/time"
)

// EmptyFunction type
type EmptyFunction func()

type chanFunc struct {
	channel   chan interface{}
	functions []EmptyFunction
}

var (
	cronnes       = map[time.Duration]map[*time.Location]*chanFunc{}
	cronnesLocker sync.RWMutex
)

// Add adds callback to be called every d duration starting from beginning of the day
func Add(d time.Duration, timeZone *time.Location, f EmptyFunction) {
	cronnesLocker.Lock()
	if _, ok := cronnes[d]; !ok {
		cronnes[d] = make(map[*time.Location]*chanFunc)
	}
	if _, ok := cronnes[d][timeZone]; !ok {
		cronnes[d][timeZone] = &chanFunc{
			channel: make(chan interface{}),
		}
		go processCron(d, timeZone, cronnes[d][timeZone].channel)
	}
	cronnes[d][timeZone].functions = append(cronnes[d][timeZone].functions, f)
	cronnesLocker.Unlock()
}

func processCron(d time.Duration, timeZone *time.Location, ch chan interface{}) {
	ticker := updateTicker(d, timeZone)
	for {
		select {
		case <-ticker.C:
			cronnesLocker.RLock()
			if l, ok := cronnes[d][timeZone]; ok {
				for _, f := range l.functions {
					go f()
				}
			}
			cronnesLocker.RUnlock()
			ticker = updateTicker(d, timeZone)
		case <-ch:
			ticker.Stop()
			cronnesLocker.Lock()
			if l, ok := cronnes[d][timeZone]; ok {
				close(l.channel)
			}
			delete(cronnes[d], timeZone)
			if len(cronnes[d]) == 0 {
				delete(cronnes, d)
			}
			cronnesLocker.Unlock()
			return
		}
	}
}

func updateTicker(d time.Duration, timeZone *time.Location) *time.Ticker {
	now := time.Now()
	nextTick := t.StartOfDay(now, timeZone)
	for !nextTick.After(now) {
		nextTick = nextTick.Add(d)
	}
	return time.NewTicker(nextTick.Sub(now))
}
