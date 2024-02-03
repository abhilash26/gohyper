package view

import (
	"html/template"
	"log"
	"net/http"
	"strings"
)

func RenderTemplate(w http.ResponseWriter, templatePaths []string, templateName string, data interface{}) {
	// Append ".html" to the template name if it doesn't have it already.
	if !strings.HasSuffix(templateName, ".html") {
		templateName += ".html"
	}

	// Add the prefix to each template path.
	for i, path := range templatePaths {
		templatePaths[i] = "./internal/view/" + path + ".html"
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
