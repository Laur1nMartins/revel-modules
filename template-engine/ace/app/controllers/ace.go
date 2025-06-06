package controllers

import "github.com/Laur1nMartins/revel"

type AceController struct {
	*revel.Controller
}

// Called to render the ace template inner.
func (c *AceController) RenderAceTemplate(base, inner string) revel.Result {
	c.ViewArgs["ace-inner"] = inner
	return c.RenderTemplate(base)
}
