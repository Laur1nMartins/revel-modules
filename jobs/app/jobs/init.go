package jobs

import (
	"github.com/Laur1nMartins/revel"
)

var jobLog = revel.AppLog

func init() {
	revel.RegisterModuleInit(func(m *revel.Module) {
		jobLog = m.Log
	})
}
