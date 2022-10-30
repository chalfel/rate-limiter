package network

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/sirupsen/logrus"
)

type Server struct {
	Client *http.Server
	router *Router
}

func NewServer(router *Router, port string) *Server {
	return &Server{
		router: router,
		Client: &http.Server{
			Addr:    port,
			Handler: router.Engine,
		},
	}
}

func (s *Server) Init() error {
	go func() {
		if err := s.Client.ListenAndServe(); err != nil {
			logrus.WithError(err).Info("something went wrong")
		}

		logrus.Infof("Server is running and listening on port %s\n", s.Client.Addr)
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	logrus.Info("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	if err := s.Client.Shutdown(ctx); err != nil {
		logrus.Errorf("Server forced to shutdown: %v", err)
		return err
	}

	logrus.Info("Server exiting")

	return nil
}
