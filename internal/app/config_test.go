package app_test

import (
	"github.com/solivaf/go-geoffrey/internal/app"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	assert.NoError(t, os.Setenv("GEOFFREY_CONFIG", "./../../testdata/application.yml"), "Error to set envVar")

	config := app.LoadConfig()
	serverConfig := config.ServerConfig()
	gitConfig := config.GitConfig()

	assert.Equal(t, "9090", serverConfig.Port())
	assert.Equal(t, "https://github.com/solivaf/geoffrey", gitConfig.Url())
	assert.Equal(t, "johndoe", gitConfig.Username())
	assert.Equal(t, "johndoepassword", gitConfig.Password())

	repositories := gitConfig.Repositories()
	assert.Equal(t, 1, len(repositories))

	repository := repositories[0]
	assert.Equal(t, "go-maria", repository.Name())
	assert.Equal(t, "https://github.com/solivaf/go-maria", repository.Url())
	assert.Equal(t, "janedoe", repository.Username())
	assert.Equal(t, "janedoepassword", repository.Password())
}
