package git

import (
	"github.com/solivaf/go-geoffrey/internal/config"
	"log"
	"strings"
)

type GitFetcher struct {
	downloader Downloader
	config     *config.GitConfig
}

func NewFetcher(config *config.GitConfig, downloader Downloader) Fetcher {
	return &GitFetcher{downloader: downloader, config: config}
}

func (f *GitFetcher) Fetch() error {
	f.fetchDefaultRepository()

	for _, repository := range f.config.Repositories() {
		go func() {
			f.downloadRepository(repository)
		}()
	}

	return nil
}

func (f *GitFetcher) fetchDefaultRepository() {
	var defaultRepository config.Repository
	url := f.config.Url()
	if strings.HasPrefix(url, "http")  {
		user := f.config.Username()
		password := f.config.Password()
		defaultRepository = config.NewRepositoryPassword(url, "default", user, password)
	}
	if strings.HasPrefix(url, "git") {
		defaultRepository = config.NewRepositorySsh(url, "default")
	}
	f.downloadRepository(defaultRepository)
}

func (f *GitFetcher) downloadRepository(repository config.Repository) {
	log.Println("cloning repository " + repository.Name())
	if message, err := f.downloader.Download(repository); err != nil {
		log.Fatal("error cloning repository "+repository.Name()+" message "+message, err)
	}
	log.Println("repository " + repository.Name() + " cloned successfully ")
}
