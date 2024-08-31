package handlers

import (
	"net/http"

	"github.com/rubutar/goProject/pkg/config"
	"github.com/rubutar/goProject/pkg/models"
	"github.com/rubutar/goProject/pkg/render"
)

// Repository used by the handlers
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

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	intMap := make(map[string]int)
	intMap["numb"] = 5

	render.RenderTemplate(w, "home.page.html", &models.TemplateData{
		IntMap: intMap,
	})
}
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	// perform logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again."
	stringMap["two"] = "Here, I am."

	//send the data to the template
	render.RenderTemplate(w, "about.page.html", &models.TemplateData{
		StringMap: stringMap,
	})
}
