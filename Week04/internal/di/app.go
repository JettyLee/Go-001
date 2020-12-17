package di

import (
	"JettyPayGo/internal/router"
	"JettyPayGo/internal/service"
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type App struct {
	svc  *service.Service
	http *http.Server
}

func ProvideApp(svc *service.Service) (app *App) {
	gin.SetMode(gin.DebugMode)
	router := router.NewRouter(svc)
	s := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	app = &App{
		svc:  svc,
		http: s,
	}

	go func() {
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("s.ListenAndServe err: %v", err)
		}
	}()

	return app
}

func (app *App) CloseFunc() {
	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	if err := app.http.Shutdown(ctx); err != nil {
		log.Println("http.Shutdown error(%v)", err)
	}
	cancel()

	log.Println("Shuting down server...")
}
