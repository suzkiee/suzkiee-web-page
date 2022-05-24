package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

// we will be able to build functions and pass them to the template
var functions = template.FuncMap{}

// RenderTemplate renders a template with html templates
func RenderTemplate(w http.ResponseWriter, tmpl string) {
	tc, err := CreateTemplateCache()
	if err != nil {
		log.Fatal(err)
	}

	t, ok := tc[tmpl]
	if !ok {
		log.Fatal(err)
	}

	// holds bytes of templates
	// It helps to have a buffer so that you can determine where the error occurred more easily
	buf := new(bytes.Buffer)

	// Takes the template, executes it and store in the buf variable
	_ = t.Execute(buf, nil)
	_, err = buf.WriteTo(w)

	if err != nil {
		fmt.Println("error writing template to browser:", err)
	}
}

// CreateTemplateCache creates a map of html templates
func CreateTemplateCache() (map[string]*template.Template, error) {
	// Create a new map of templates
	tmplCache := map[string]*template.Template{}

	// go to the templates directory and get ANYTHING that has page.tmpl
	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return tmplCache, err
	}

	// for each page template...
	for _, page := range pages {
		// get the name of the template
		name := filepath.Base(page)

		// build a new template, pass it our map of custom functions and parse the file
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return tmplCache, err
		}

		// checks in the templates directory for any templates that end in .layout.tmpl
		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return tmplCache, err
		}

		// if there are any matches, parse the glob
		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return tmplCache, err
			}
		}
		// let the name of the template map to the value of the new template we created
		tmplCache[name] = ts
	}

	return tmplCache, nil
}
