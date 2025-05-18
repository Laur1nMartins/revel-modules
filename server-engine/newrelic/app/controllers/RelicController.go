package controllers

import (
	"github.com/Laur1nMartins/revel"
	revelnewrelic "github.com/Laur1nMartins/revel-modules/server-engine/newrelic"
	newrelic "github.com/newrelic/go-agent"
)

type RelicController struct {
	*revel.Controller
}

func (r *RelicController) GetRelicApplication() newrelic.Application {
	if app, ok := revel.CurrentEngine.(*revelnewrelic.ServerNewRelic); ok {
		return app.NewRelicApp
	}
	return nil
}
