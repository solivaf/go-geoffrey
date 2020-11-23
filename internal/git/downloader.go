package git

import (
	"github.com/solivaf/go-geoffrey/internal/config"
	"github.com/solivaf/go-geoffrey/internal/pkg/command"
	_repository "github.com/solivaf/go-geoffrey/internal/repository"
	"log"
	_url "net/url"
	"strings"
)

const httpsPrefix = "https://"
const httpPrefix = "http://"

type Downloader interface {
	Download(repository config.Repository) (string, error)
}

type downloader struct {
	commander command.Commander
}

func NewDownloader(commander command.Commander) Downloader {
	return &downloader{commander: commander}
}

func (d *downloader) Download(repository config.Repository) (string, error) {
	log.Println("cloning repository " + repository.Url())
	d.commander.Execute("rm", "-rf", _repository.CheckoutPath(repository.Name()))
	if repository.IsSsh() {
		return d.commander.Execute("git", "clone", repository.Url(), _repository.CheckoutPath(repository.Name()))
	}
	formattedUrl := d.formatUrl(repository.Url(), repository.Username(), repository.Password())
	return d.commander.Execute("git", "clone", formattedUrl, _repository.CheckoutPath(repository.Name()))
}

func (d *downloader) formatUrl(url, username, password string) string {
	if isHttps := d.isHttps(url); isHttps {
		urlParts := strings.Split(url, httpsPrefix)
		return d.getFormattedUrl(httpsPrefix, username, password, urlParts)
	}

    urlParts := strings.Split(url, httpPrefix)
    return d.getFormattedUrl(httpPrefix, username, password, urlParts)
}

func (d *downloader) getFormattedUrl(prefix string, username string, password string, urlParts []string) string {
	return prefix + _url.QueryEscape(username) + ":" + _url.QueryEscape(password) + "@" + urlParts[1] + ".git"
}

func (d *downloader) isHttps(url string) bool {
	return strings.HasPrefix(url, httpsPrefix)
}
