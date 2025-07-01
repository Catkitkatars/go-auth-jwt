package http

import (
	"authjwt/internal/config"
	"authjwt/internal/http/router"
	"net/http"
	"strconv"
)

func ServerStart() error {
	cfg := config.Cfg
	r := router.InitRouter()

	srv := &http.Server{
		Addr:         cfg.AppHost + ":" + strconv.Itoa(cfg.AppPort),
		Handler:      r,
		ReadTimeout:  cfg.AppTimeout,
		WriteTimeout: cfg.AppTimeout,
		IdleTimeout:  cfg.AppIdleTimeout,
	}

	err := srv.ListenAndServe()

	return err
}
