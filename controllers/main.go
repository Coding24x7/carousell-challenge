//go:generate goagen bootstrap -d github.com/Coding24x7/twitter-clone/webapp/design

package controllers

import (
	"net/http"
	"time"

	"github.com/Coding24x7/twitter-clone/app"
	"github.com/goadesign/goa"
	goalogrus "github.com/goadesign/goa/logging/logrus"
	"github.com/goadesign/goa/middleware"
	log "github.com/sirupsen/logrus"
	"github.com/tylerb/graceful"
)

// GoaServer starts server
func GoaServer(logger *log.Logger, address string) {
	// Create service
	service := goa.New("twitter-clone")
	logAdapter := goalogrus.New(logger)
	// set up service to use logger
	service.WithLogger(logAdapter)

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// Mount "public" controller
	app.MountPublicController(service, NewPublicController(service))
	// Mount "swagger" controller
	app.MountSwaggerController(service, NewSwaggerController(service))
	// Mount "topics" controller
	app.MountTopicController(service, NewTopicController(service))
	// Mount "user" controller
	c4 := NewUserController(service)
	app.MountUserController(service, c4)

	// Setup graceful server
	server := &graceful.Server{
		Timeout: time.Duration(15) * time.Second,
		Server: &http.Server{
			Addr:    address,
			Handler: service.Mux,
		},
	}

	// Start service
	if err := server.ListenAndServe(); err != nil {
		service.LogError("startup", "err", err)
	}

}
