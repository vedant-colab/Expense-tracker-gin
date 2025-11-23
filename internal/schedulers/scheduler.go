package scheduler

import (
	"time"

	"github.com/go-co-op/gocron"
)

type Scheduler struct {
	s *gocron.Scheduler
}

func New() *Scheduler {
	return &Scheduler{
		s: gocron.NewScheduler(time.Local),
	}
}

// Register a recurring job
func (sc *Scheduler) EveryMinute(job func()) {
	sc.s.Every(1).Minute().Do(job)
}

func (sc *Scheduler) Start() {
	sc.s.StartAsync()
}
