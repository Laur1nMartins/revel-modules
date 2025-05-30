package controllers

import (
	"net/http"
	"net/http/pprof"

	"github.com/Laur1nMartins/revel"
)

type Pprof struct {
	*revel.Controller
}

// The PprofHandler type makes it easy to call the net/http/pprof handler methods
// since they all have the same method signature.
type PprofHandler func(http.ResponseWriter, *http.Request)

func (r PprofHandler) Apply(req *revel.Request, resp *revel.Response) {
	r(resp.Out.Server.GetRaw().(http.ResponseWriter), req.In.GetRaw().(*http.Request))
}

func (c Pprof) Profile() revel.Result {
	return PprofHandler(pprof.Profile)
}

func (c Pprof) Symbol() revel.Result {
	return PprofHandler(pprof.Symbol)
}

func (c Pprof) Cmdline() revel.Result {
	return PprofHandler(pprof.Cmdline)
}

func (c Pprof) Trace() revel.Result {
	return PprofHandler(pprof.Trace)
}

func (c Pprof) Index() revel.Result {
	return PprofHandler(pprof.Index)
}
