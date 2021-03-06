package core

import (
	"net/http"
	"strings"
	"time"

	"agfun/test/rtsp-stream/core/config"
	"agfun/test/rtsp-stream/core/streaming"
	"github.com/julienschmidt/httprouter"
)

// determinesHost is for parsing out the host from the storage path
func determineHost(path string) string {
	parts := strings.Split(path, "/")
	if len(parts) >= 1 {
		return parts[1]
	}
	return ""
}

// GetRouter returns the return for the application
func GetRouter(config *config.Specification) (*httprouter.Router, *Controller) {
	fileServer := http.FileServer(http.Dir(config.StoreDir))
	router := httprouter.New()
	controllers := Controller{config, map[string]*streaming.Stream{}, fileServer, Manager{}, streaming.Processor{}, time.Second * 15}
	if config.ListEndpoint {
		router.GET("/list", controllers.ListStreamHandler)
	}
	router.POST("/start", controllers.StartStreamHandler)
	router.GET("/stream/*filepath", controllers.FileHandler)

	// Start cleaning process in the background
	go func() {
		for {
			<-time.After(config.CleanupTime)
			controllers.cleanUnused()
		}
	}()

	return router, &controllers
}
