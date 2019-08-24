package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/solivaf/go-geoffrey/internal/api/handler"
	"github.com/solivaf/go-geoffrey/internal/app"
	"github.com/solivaf/go-geoffrey/internal/git"
	"github.com/solivaf/go-geoffrey/internal/pkg/command"
	"github.com/solivaf/go-geoffrey/internal/repository"
	"log"
	"net/http"
	"time"
)

func main() {
	config := app.LoadConfig()
	commander := command.NewCommander()
	downloader := git.NewDownloader(commander)
	fetcher := git.NewFetcher(config.GitConfig(), downloader)

	go func() {
		fetchRepositories(fetcher)
		repository.LoadFiles()
	}()

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(5 * time.Second))
	h := handler.NewHandler(fetcher)
	r.Get("/{app}/{env}", h.RetrieveConfig)
	r.Get("/health", h.Health)

	log.Fatal(http.ListenAndServe(":3000", r))
}

func fetchRepositories(fetcher git.Fetcher) {
	log.Println("loading files from repositories")
	if err := fetcher.Fetch(); err != nil {
		log.Fatal("error while downloading files", err)
	}
	log.Println("repository cloned successfully")
}
