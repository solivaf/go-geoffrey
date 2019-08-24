package git

import (
	"github.com/solivaf/go-geoffrey/internal/config"
	"log"
)

type GitFetcher struct {
	downloader Downloader
	config     *config.GitConfig
}

func NewFetcher(config *config.GitConfig, downloader Downloader) Fetcher {
	return &GitFetcher{downloader: downloader, config: config}
}

func (f *GitFetcher) Fetch() error {
	url := f.config.Url()
	user := f.config.Username()
	password := f.config.Password()

	for _, repository := range f.config.Repositories() {
		go func() {
			f.downloadRepository(repository)
		}()
	}
	defaultRepository := config.NewRepository(url, "default", user, password)
	f.downloadRepository(defaultRepository)
	return nil
}

func (f *GitFetcher) downloadRepository(repository config.Repository) {
	log.Println("cloning repository " + repository.Name())
	if message, err := f.downloader.Download(repository); err != nil {
		log.Fatal("error cloning repository "+repository.Name()+" message "+message, err)
	}
	log.Println("repository " + repository.Name() + " cloned successfully ")
}
