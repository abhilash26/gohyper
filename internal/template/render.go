package view

import (
	"html/template"
	"log"
	"net/http"
	"strings"

	"github.com/abhilash26/gohyper/internal/option"
)

func RenderTemplate(w http.ResponseWriter, templatePaths []string, templateName string, data interface{}) {
	viewPath := option.GetStringEnv("VIEW_PATH", "./internal/view/")
	viewExtension := option.GetStringEnv("VIEW_EXTENSION", ".tmpl")
	// Append ".html" to the template name if it doesn't have it already.
	if !strings.HasSuffix(templateName, viewExtension) {
		templateName += viewExtension
	}

	// Add the prefix to each template path.
	for i, path := range templatePaths {
		templatePaths[i] = viewPath + path + viewExtension
	}

	// Parse the templates.
	templates, err := template.ParseFiles(templatePaths...)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Println("Error parsing templates:", err)
		return
	}

	// Execute the templates and pass the data.
	err = templates.ExecuteTemplate(w, templateName, data)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Println("Error executing template:", err)
		return
	}
}
