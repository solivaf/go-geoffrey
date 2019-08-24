package handler

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/solivaf/go-geoffrey/internal/git"
	"github.com/solivaf/go-geoffrey/internal/repository"
	"gopkg.in/yaml.v2"
	"net/http"
)

type Handler struct {
	fetcher git.Fetcher
}

func NewHandler(fetcher git.Fetcher) *Handler {
	return &Handler{fetcher}
}

func (h *Handler) RetrieveConfig(w http.ResponseWriter, r *http.Request) {
	app := chi.URLParam(r, "app")
	env := chi.URLParam(r, "env")
	response := h.createContent(app, env)

	m := make(map[string]interface{})
	if err := yaml.Unmarshal([]byte(response), &m); err != nil {
		w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
	}
	b, _ := yaml.Marshal(&m)

	w.Write(b)
	w.Header().Set("Content-Type", "application/json")
}

func (h *Handler) Health(w http.ResponseWriter, r *http.Request) {
	files := repository.Files()
	fileNames := make([]string, 0)
	for fileName, _ := range files {
		fileNames = append(fileNames, fileName)
	}

	resp := make(map[string][]string)
	resp["files"] = fileNames
	b, err := json.Marshal(&resp)
	if err != nil {
		w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
	w.WriteHeader(http.StatusOK)

}

func (h *Handler) createContent(app string, env string) string {
	files := repository.Files()
	var response string
	if content, exists := files[repository.GetFormattedFileNameWithEnv(app, env)]; exists {
		response += content + "\n"
	}
	if content, exists := files[repository.GetFormattedFileName(app)]; exists {
		response += content + "\n"
	}
	if content, exists := files[repository.GetDefaultFileNameWithEnv(env)]; exists {
		response += content + "\n"
	}
	if content, exists := files[repository.GetDefaultFileName()]; exists {
		response += content
	}
	return response
}

func (h *Handler) isValidFileName(fileName string, app string, env string) bool {
	return fileName == repository.GetFormattedFileNameWithEnv(app, env) ||
		fileName == repository.GetDefaultFileName() ||
		fileName == repository.GetDefaultFileNameWithEnv(env)
}
