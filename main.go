package main

import (
	"net/http"
	"time"

	. "github.com/YAWAL/ETLUNL/logging"
	"github.com/YAWAL/ETLUNL/router"

)

func main() {

	// init router
	r := router.New()

	srv := &http.Server{
		Handler:      r,
		Addr:         ":8080",
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}
	Log.Infof("Application is running on %s ", srv.Addr)

	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		Log.Errorf("Could not listen on %s: %v\n", srv.Addr, err)
	}

}
