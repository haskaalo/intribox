package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/haskaalo/intribox/config"
	"github.com/haskaalo/intribox/response"
	"github.com/haskaalo/intribox/routers"
	"github.com/rs/zerolog/log"
)

func main() {
	response.InitTemplates()
	router := routers.BootstrapRouters()
	log.Info().Int("port", config.Server.Port).Msg("Webserver starting")

	srv := &http.Server{
		Handler:      router,
		Addr:         fmt.Sprintf(":%v", config.Server.Port),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	srv.ListenAndServe()
}
