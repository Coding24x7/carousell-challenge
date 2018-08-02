//go:generate goagen bootstrap -o goa_temp -d github.com/citrusleaf/acc/webapp/design

package main

import (
	"runtime"
	"runtime/debug"

	log "github.com/sirupsen/logrus"

	"github.com/Coding24x7/carousell-challenge/controllers"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			log.Fatal(string(debug.Stack()))
		}
	}()

	runtime.GOMAXPROCS(runtime.NumCPU())

	log.Infof("Trying to start the server...")

	// initialize logger
	logger := log.New()

	controllers.GoaServer(logger)
}
