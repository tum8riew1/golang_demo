package main

import (
	"html/template"
	"log"
	"net/http"
	"reflect"

	"github.com/gorilla/mux"
)

func articlesHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "golang-session")
	//session.Values["token"]
	token := &Token{
		Token:   session.Values["token"].(string),
		App_url: app_url,
		Api_url: api_url,
		User_id: session.Values["id"].(int),
	}
	t := template.Must(template.ParseFiles("views/layouts/layout.html", "views/partials/sidebar.html", "views/layouts/articles.html"))
	t.Execute(w, token)
}

func articlesAddHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "golang-session")
	//session.Values["token"]
	token := &Token{
		Token:   session.Values["token"].(string),
		App_url: app_url,
		Api_url: api_url,
		User_id: session.Values["id"].(int),
	}
	t := template.Must(template.ParseFiles("views/layouts/layout.html", "views/partials/sidebar.html", "views/layouts/articles_add.html"))
	t.Execute(w, token)
}

func articlesEditHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	session, _ := store.Get(r, "golang-session")
	//session.Values["token"]
	token := &Token{
		Token:   session.Values["token"].(string),
		App_url: app_url,
		Api_url: api_url,
		User_id: session.Values["id"].(int),
		ID:      params["id"],
	}
	log.Println(params["id"], reflect.TypeOf(params["id"]))
	t := template.Must(template.ParseFiles("views/layouts/layout.html", "views/partials/sidebar.html", "views/layouts/articles_edit.html"))
	t.Execute(w, token)
}
