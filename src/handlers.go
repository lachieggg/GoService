package main

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
)

// indexHandler
func indexHandler(w http.ResponseWriter, r *http.Request) {

	wd, err := os.Getwd()
	templateFolder := filepath.Join(wd, "src", "templates")
	templatePath := filepath.Join(templateFolder, "*.html")

	// Parse all templates in the specified directory
	t, err := template.ParseGlob(templatePath)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error parsing templates: %v", err), http.StatusInternalServerError)
		return
	}

	err = t.ExecuteTemplate(w, "app.html", nil)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error executing template: %v", err), http.StatusInternalServerError)
	}

	// register the include function
	t.Funcs(template.FuncMap{
		"include": func(name string, data interface{}) (string, error) {
			// load the named template file
			tpl, err := template.ParseFiles(name)
			if err != nil {
				return "", err
			}

			// execute the named template with the provided data
			var buf bytes.Buffer
			if err := tpl.Execute(&buf, data); err != nil {
				return "", err
			}
			return buf.String(), nil
		},
	})
}

func filesHandler(w http.ResponseWriter, r *http.Request) {
	files, err := os.ReadDir("/")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	for _, file := range files {
		fmt.Fprintf(w, "%s\n", filepath.Join("/", file.Name()))
	}

	fmt.Fprint(w, "\n")
	files, err = os.ReadDir("/app")
	for _, file := range files {
		fmt.Fprintf(w, "%s\n", filepath.Join("/", file.Name()))
	}

	fmt.Fprint(w, "\n")
	wd, err := os.Getwd()
	fmt.Fprintf(w, "%s", wd)
}

// Home
func Home(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello, World!")
}

// githubHandler
func githubHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "https://github.com/lachieggg", http.StatusSeeOther)
}
