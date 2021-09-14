package handlers

import (
	"github.com/dimedrolling/bookings/pkg/config"
	"github.com/dimedrolling/bookings/pkg/models"
	"github.com/dimedrolling/bookings/pkg/render"
	"net/http"
)

//Repo is the repository used
var Repo *Repository

//Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

//NewRepo creates a repos
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

//NewHandlers sets the repository for the new handlers
func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	host := m.App.Session.GetString(r.Context(), "host")
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)
	stringMap["host"] = host
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	//perform some logic

	stringMap := make(map[string]string)
	stringMap["test"] = "Hello again "
	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")

	stringMap["remote_ip"] = remoteIP

	host := r.Host
	m.App.Session.Put(r.Context(), "host", host)
	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
