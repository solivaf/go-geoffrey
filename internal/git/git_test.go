package git_test

import (
	"github.com/golang/mock/gomock"
	"github.com/solivaf/go-geoffrey/internal/config"
	"github.com/solivaf/go-geoffrey/internal/git"
	"github.com/solivaf/go-geoffrey/internal/git/mock_git"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestGitFetcher_Fetch(t *testing.T) {
	c := &config.GitConfig{}
	ctrl := gomock.NewController(t)
	mockDownloader := mock_git.NewMockDownloader(ctrl)

	fetcher := git.NewFetcher(c, mockDownloader)

	path, err := filepath.Abs(filepath.Dir("./../../testdata/"))
	f, err := os.Open(path + "/" + "application.yml")
	content, err := ioutil.ReadAll(f)

	repository := config.NewRepositoryPassword(c.Url(), "default", c.Username(), c.Password())
	mockDownloader.EXPECT().Download(repository).Return(string(content), err)

	assert.NoError(t, fetcher.Fetch())
}
