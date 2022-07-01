package handlers

import (
	"net/http"

	"github.com/evolutionengine/bookings/cmd/pkg/config"
	"github.com/evolutionengine/bookings/cmd/pkg/models"
	"github.com/evolutionengine/bookings/cmd/pkg/render"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Hello(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remoteIP", remoteIP)

	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	remoteIP := m.App.Session.GetString(r.Context(), "remoteIP")

	stringMap := make(map[string]string)

	stringMap["test"] = "Hello from Zortan"
	stringMap["remoteIP"] = remoteIP

	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{StringMap: stringMap})
}
