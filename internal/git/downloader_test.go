package git_test

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go-geoffrey/internal/config"
	"go-geoffrey/internal/git"
	"go-geoffrey/internal/pkg/command/mock_command"
	"testing"
)

func TestDownloader_Download(t *testing.T) {
	ctrl := gomock.NewController(t)
	commanderMock := mock_command.NewMockCommander(ctrl)
	downloader := git.NewDownloader(commanderMock)
	repo := config.NewRepository("https://github.com/anyowner/anyrepo", "johndoerepo", "johndoe", "johndoepassword")

	commanderMock.EXPECT().Execute("git", "clone", "https://johndoe:johndoepassword@github.com/anyowner/anyrepo.git", "/tmp/go-geoffrey/johndoerepo").
		Return("return message", nil)

	message, err := downloader.Download(repo)

	assert.NoError(t, err)
	assert.Equal(t, "return message", message)
}
