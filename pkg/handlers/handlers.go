package handlers

import (
	"net/http"

	"github.com/leonard-ladva/bookings/pkg/config"
	"github.com/leonard-ladva/bookings/pkg/models"
	"github.com/leonard-ladva/bookings/pkg/render"
)

// Repo the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the homepage handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, "home.page.html", &models.TemplateData{})
}

// About is the about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {

	stringmap := make(map[string]string)
	stringmap["test"] = "hello, again"

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringmap["remote_ip"] = remoteIP

	// Send the data to the template
	render.RenderTemplate(w, "about.page.html", &models.TemplateData{
		StringMap: stringmap,
	})
}