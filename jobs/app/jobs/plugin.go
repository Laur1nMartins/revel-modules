package jobs

import (
	"github.com/Laur1nMartins/revel"
	"github.com/revel/cron"
)

const DefaultJobPoolSize = 10

var (
	// MainCron is the singleton instance of the underlying job scheduler.
	MainCron *cron.Cron

	// This limits the number of jobs allowed to run concurrently.
	workPermits chan struct{}

	// Is a single job allowed to run concurrently with itself?
	selfConcurrent bool
)

func init() {
	MainCron = cron.New()
	revel.OnAppStart(func() {
		if size := revel.Config.IntDefault("jobs.pool", DefaultJobPoolSize); size > 0 {
			workPermits = make(chan struct{}, size)
		}
		selfConcurrent = revel.Config.BoolDefault("jobs.selfconcurrent", false)
		MainCron.Start()
		jobLog.Info("Go to /@jobs to see job status.")
	})
}
