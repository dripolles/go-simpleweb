package testapp

import (
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"os"
)

type Page struct {
	Title string
	Body  []byte
}

func loadPage(title string) (*Page, error) {
	filename := "/tmp/" + title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

type View struct{}

func (v *View) Bound(
	f func(*View, http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {

	bound := func(w http.ResponseWriter, r *http.Request) {
		f(v, w, r)
	}
	return bound
}

type TestView struct {
	View
	Value int
}

func (v *TestView) TestHandler(w http.ResponseWriter, r *http.Request) {
	r.Write(os.Stdout)
	vars := mux.Vars(r)
	title := vars["filename"]
	fmt.Println(title)

	p, err := loadPage(title)
	if err == nil {
		v.Value++
		fmt.Fprintf(w, "T: %s -- B: %s (%d)\n", p.Title, p.Body, v.Value)
	} else {
		w.WriteHeader(404)
		w.Write([]byte("404 - Not found :(\n"))

	}
}
