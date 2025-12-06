package server

import (
	"context"
	"errors"
	"fmt"
	"github.com/saleh-ghazimoradi/StandardMicroEcoBay/internal/helper"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Server struct {
	Host         string
	Port         string
	Handler      http.Handler
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	IdleTimeout  time.Duration
	ErrorLog     *log.Logger
}

type Options func(*Server)

func WithHost(host string) Options {
	return func(s *Server) {
		s.Host = host
	}
}

func WithPort(port string) Options {
	return func(s *Server) {
		s.Port = port
	}
}

func WithReadTimeout(timeout time.Duration) Options {
	return func(s *Server) {
		s.ReadTimeout = timeout
	}
}

func WithWriteTimeout(timeout time.Duration) Options {
	return func(s *Server) {
		s.WriteTimeout = timeout
	}
}

func WithIdleTimeout(timeout time.Duration) Options {
	return func(s *Server) {
		s.IdleTimeout = timeout
	}
}

func WithErrorLog(logger *log.Logger) Options {
	return func(s *Server) {
		s.ErrorLog = logger
	}
}

func (s *Server) Connect() error {
	addr := fmt.Sprintf("%s:%s", s.Host, s.Port)
	server := &http.Server{
		Addr:         addr,
		Handler:      s.Handler,
		ReadTimeout:  s.ReadTimeout,
		WriteTimeout: s.WriteTimeout,
		IdleTimeout:  s.IdleTimeout,
		ErrorLog:     s.ErrorLog,
	}

	shutdownError := make(chan error)

	go func() {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		se := <-quit
		log.Println("caught signal", "signal", se.String())

		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		err := server.Shutdown(ctx)
		if err != nil {
			shutdownError <- err
		}

		log.Println("completing background tasks", "addr", server.Addr)
		helper.WG.Wait()
		shutdownError <- nil
	}()

	if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	if err := <-shutdownError; err != nil {
		return err
	}

	log.Println("Stopped server", "addr", server.Addr)

	return nil
}

func NewServer(opts ...Options) *Server {
	s := &Server{}
	for _, o := range opts {
		o(s)
	}
	return s
}
