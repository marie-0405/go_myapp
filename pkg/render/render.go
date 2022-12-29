package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

func RenderTemplate(w http.ResponseWriter, tmpl string) {
	// get the template cache from the app config

	// create cached template
	tc, err := CreateTemplateCache()
	if err != nil {
		log.Fatal(err)
	}
	// get requested template
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("tc has no key of tmpl")
	}

	buf := new(bytes.Buffer)

	err = t.Execute(buf, nil)
	if err != nil {
		log.Println(err)
	}

	// render template
	_, err = buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			// ts, err = template.ParseGlob("./templates/*.layout.tmpl")
		}
		if err != nil {
			return myCache, err
		}

		myCache[name] = ts

	}
	return myCache, nil
}
