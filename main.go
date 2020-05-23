package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"log"
	"net/http"
)

//"github.com/pilu/resh" in used for rebuilding and running
// on file change
//type fresh in cmd to start it
//settings for it are in runner.conf

var homeTemplate *template.Template

func home(
	w http.ResponseWriter,
	r *http.Request,
	_ httprouter.Params) {
	w.Header().Set("Content-Type", "text/html")
	if err := homeTemplate.Execute(w, nil); err != nil {
		panic(err)
	}
}

func login(
	w http.ResponseWriter,
	r *http.Request,
	_ httprouter.Params) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>Login</h1>")
}

func user(
	w http.ResponseWriter,
	r *http.Request,
	u httprouter.Params) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>Login : ", u.ByName("name"), "</h1>")
}

func deck(
	w http.ResponseWriter,
	r *http.Request,
	u httprouter.Params) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>Deck : ", u.ByName("name"), "</h1>")
}

func contact(
	w http.ResponseWriter,
	r *http.Request,
	_ httprouter.Params) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, `to get in touch send an email
		to <a href="mailto:goLearnSupport@gmail.com">
		goLearnSupport@gmail.com</a>`)
}

func faq(
	w http.ResponseWriter,
	r *http.Request,
	_ httprouter.Params) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, `<h1> FAQ </h1>`)
}

func notFound(
	w http.ResponseWriter,
	r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, `<h1> We couldn't find the page
		you're looking for </h1>`)
}

func main() {
	var err error
	homeTemplate, err = template.ParseFiles("views/home.gohtml")
	if err != nil {
		panic(err)
	}
	r := httprouter.New()
	r.NotFound = http.HandlerFunc(notFound)
	r.GET("/", home)
	r.GET("/contact", contact)
	r.GET("/faq", faq)
	r.GET("/login", login)
	r.GET("/users/:name", user)
	r.GET("/decks/:name", deck)
	log.Fatal(http.ListenAndServe("localhost:8080", r))
}
