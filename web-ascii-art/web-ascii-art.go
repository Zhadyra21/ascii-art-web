package asciiartweb

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	asciiart "web/ascii-art"
)

type Text struct {
	UserText string
	Result   string
}

type appError struct {
	Error   error
	Message string
	Code    int
}

type appHandler func(http.ResponseWriter, *http.Request) *appError

func (fn appHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer func() {
		err := recover()
		if err != nil {
			http.Error(w, "Internal Server Error", 500)
		}
	}()
	if e := fn(w, r); e != nil {
		w.WriteHeader(e.Code)
		if tmpl, err := template.ParseFiles("templates/err.html"); err == nil {
			tmpl.Execute(w, &e)
		} else {
			http.Error(w, "Internal Server Error", 500)
		}
		if e.Error != nil {
			fmt.Println(e.Error)
		}
	}
}

func Init() {
	http.Handle("/", appHandler(Home_page))
	http.Handle("/ascii-art", appHandler(Art_page))
}

func Home_page(w http.ResponseWriter, r *http.Request) *appError {
	if r.URL.Path != "/" {
		return &appError{nil, "Page Not Found", 404}
	}
	switch r.Method {
	case "GET":
		htmlfile := "templates/home_page.html"
		if !asciiart.FileExists(htmlfile) {
			return &appError{fmt.Errorf("couldn't find an html file"), "Not Found", 404}
		}
		tmpl, err := template.ParseFiles(htmlfile)
		if err == nil {
			err = tmpl.Execute(w, nil)
		}
		if err != nil {
			return &appError{err, "Internal Server Error", 500}
		}
	default:
		return &appError{nil, "Method Not Allowed", 405}
	}
	return nil
}

func Art_page(w http.ResponseWriter, r *http.Request) *appError {
	switch r.Method {
	case "POST":
		htmlfile := "templates/home_page.html"
		if !asciiart.FileExists(htmlfile) {
			return &appError{fmt.Errorf("couldn't find an html file"), "Not Found", 404}
		}
		tmpl, err := template.ParseFiles(htmlfile)
		if err != nil {
			return &appError{err, "Internal Server Error", 500}
		}
		text := Text{}
		text.UserText = r.FormValue("text")
		if !asciiart.IsValidText(text.UserText) {
			return &appError{fmt.Errorf("please, don't use non-ascii character"), "Bad Request", 400}
		}
		banner := r.FormValue("banner")
		if !asciiart.FileExists("ascii-art/" + banner + ".txt") {
			return &appError{fmt.Errorf("couldn't find the %q banner", banner), "Internal Server Error", 500}
		}
		symbols, err := asciiart.PrintSymbols(text.UserText, banner)
		text.Result = symbols
		if err != nil {
			return &appError{err, "Internal Server Error", 500}
		}
		btn := r.FormValue("button")
		if btn == "show" {
			err = tmpl.Execute(w, text)
			if err != nil {
				return &appError{err, "Internal Server Error", 500}
			}
		} else if btn == "download" {
			w.Header().Set("Content-Type", "text/plain")
			w.Header().Set("Content-Disposition", "attachment; filename=ascii-art.txt")
			w.Header().Set("Content-Length", strconv.Itoa(len(symbols)))
			_, err := w.Write([]byte(symbols))
			if err != nil {
				return &appError{err, "Internal Server Error", 500}
			}
		} else {
			return &appError{err, "Internal Server Error", 500}
		}
	default:
		return &appError{nil, "Method not allowed", 405}
	}
	return nil
}
