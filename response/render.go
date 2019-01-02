package response

import (
	"html/template"
	"net/http"
	"time"

	"github.com/haskaalo/intribox/config"
)

// RenderData all Variables used in templates
type RenderData struct {
	Title    string
	MetaData MetaData
}

// MetaData tags used in html
type MetaData struct {
}

var (
	templates *template.Template
)

// InitTemplates initialize templates
func InitTemplates() {
	templates = template.Must(template.ParseGlob(config.Client.AssetsPath + "/html/*"))

	// Reload templates every 3 seconds in case of change
	// Must change
	if config.Debug {
		go func() {
			for {
				templates = template.Must(template.ParseGlob(config.Client.AssetsPath + "/html/*"))
				time.Sleep(3 * time.Second)
			}
		}()
	}
}

// Render render template html
func Render(w http.ResponseWriter, data RenderData) {
	templates.ExecuteTemplate(w, "indexPage", data)
}
