package render

import (
	"fmt"
	"net/http"
	"text/template"
)

func RenderTemplate(w http.ResponseWriter, templ string) {
	parsedTemplate, err := template.ParseFiles("./templates/" + templ)
	if err != nil {
		http.Error(w, "Error parsing template", http.StatusInternalServerError)
		fmt.Printf("Error parsing template: %s\n", err)
		return
	}

	err = parsedTemplate.Execute(w, nil)
	if err != nil {
		http.Error(w, "Error executing template", http.StatusInternalServerError)
		fmt.Printf("Error executing template: %s\n", err)
		return
	}
}