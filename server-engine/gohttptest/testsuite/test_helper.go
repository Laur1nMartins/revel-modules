package testsuite

import (
	"flag"
	"os"
	"sync"
	"testing"

	"github.com/Laur1nMartins/revel"
)

var importPath *string = flag.String("revel.importPath", "", "Go Import Path for the app.")

// This function is a helper to allow a test to wrap the Revel server using this
// GoHttpTest server. Which simply transfers the request / response calls using a channel.
func RevelTestHelper(m *testing.M, mode string, runner func(port int)) {
	flag.Parse()
	// call flag.Parse() here if TestMain uses flags
	locker := sync.Mutex{}
	revel.AddInitEventHandler(func(event revel.Event, value interface{}) (returnType revel.EventResponse) {
		switch event {
		case revel.REVEL_BEFORE_MODULES_LOADED:
			revel.Config.SetOption("server.engine", "go-test")
			revel.Config.SetOption("module.go-test", "github.com/revel/modules/server-engine/gohttptest")
		case revel.ENGINE_STARTED:
			go func() {
				// Wait for the server to send back a start response
				<-revel.CurrentEngine.(*GoHttpServer).StartedChan
				locker.Unlock()
			}()
		case revel.REVEL_FAILURE:
			locker.Unlock()
		default:
		}

		return 0
	})

	locker.Lock()

	revel.RevelLog.Info("Initializing the engine")
	// go test -coverprofile=coverage.out github.com/revel/examples/booking/app/controllers/  -args -revel.importPath=github.com/revel/examples/booking
	if len(*importPath) == 0 {
		// TODO add possible detection of import path from executable
		for x := 0; x < len(os.Args); x++ {
			println("App path ", os.Args[x])
		}
		serverLog.Fatal("No import path specified, aborting. Start test by using -args -revel.importPath=<your app import path>")
	}

	// Initialize revel, using the test server engine regardless of what is specified in the config.
	revel.Init(mode, *importPath, "")
	go func() {
		runner(-1)
	}()
	locker.Lock()
	result := m.Run()
	revel.StopServer(0)
	os.Exit(result)
}
