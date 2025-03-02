package main

import (
	"backend/config"
	"backend/internal/router"
	"backend/pkg/dbutil"
	"backend/pkg/logutil"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

func main() {
	if err := config.InitConfig(); err != nil {
		logutil.Error("Config initialization error: %v", err)
	}

	logutil.InitLoggers()

	dbs, err := dbutil.OpenDatabases()
	if err != nil {
		logutil.Error("Database connection error: %v", err)
	}
	defer dbs.MyDB.Close()

	e := initializeServer(dbs.MyDB)

	go func() {
		port := config.ConfigData.ServerConfig.Port
		address := fmt.Sprintf(":%s", port)
		logutil.Info("Server starting on %s", address)

		if err := e.Start(address); err != nil && err != http.ErrServerClosed {
			logutil.Error("Server error: %v", err)
		}
	}()

	gracefulShutdown(e)
}

func initializeServer(db *sqlx.DB) *echo.Echo {
	e := echo.New()
	router := router.NewRouter(e, db)
	router.SetupRoutes()
	return e
}

func gracefulShutdown(e *echo.Echo) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit
	logutil.Info("Shutting down server...")

	if err := e.Shutdown(nil); err != nil {
		logutil.Error("Server forced to shutdown: %v", err)
	}
}
