package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/dimedrolling/bookings/pkg/config"
	"github.com/dimedrolling/bookings/pkg/models"
	"github.com/dimedrolling/bookings/pkg/render"
	"log"
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
	render.RenderTemplate(w, r, "home.page.tmpl", &models.TemplateData{
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
	render.RenderTemplate(w, r, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

// Reservation renders the make a reservation page
func (m *Repository) Reservation(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "make-reservation.page.tmpl", &models.TemplateData{})
}

// Attic renders the make a room page
func (m *Repository) Attic(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "attic.page.tmpl", &models.TemplateData{})
}

// Magic renders the make a room page
func (m *Repository) Magic(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "magic.page.tmpl", &models.TemplateData{})
}

// Availability renders the make a check-availability page
func (m *Repository) Availability(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "search-availability.page.tmpl", &models.TemplateData{})
}

// PostAvailability post the form from  the make a check-availability page
func (m *Repository) PostAvailability(w http.ResponseWriter, r *http.Request) {
	start := r.Form.Get("start")
	end := r.Form.Get("end")
	w.Write([]byte(fmt.Sprintf("start date is %s and end date is %s", start, end)))

}

type jsonResponse struct {
	OK      bool   `json:"ok"`
	Message string `json:"message"`
}

// AvailabilityJSON handle request for availability from each room page
func (m *Repository) AvailabilityJSON(w http.ResponseWriter, r *http.Request) {
	resp := jsonResponse{
		OK:      true,
		Message: "Available!",
	}

	out, err := json.MarshalIndent(resp, "", "     ")
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}

// Contact renders the make a check-availability page
func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "contact.page.tmpl", &models.TemplateData{})
}
