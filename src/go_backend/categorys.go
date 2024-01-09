package main

import (
	"html/template"
	"log"
	"net/http"
	"reflect"

	"github.com/gorilla/mux"
)

func categoryHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "golang-session")
	//session.Values["token"]
	token := &Token{
		Token:   session.Values["token"].(string),
		App_url: app_url,
		Api_url: api_url,
		User_id: session.Values["id"].(int),
	}
	t := template.Must(template.ParseFiles("views/layouts/layout.html", "views/partials/sidebar.html", "views/layouts/categorys.html"))
	t.Execute(w, token)
}

func categoryAddHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "golang-session")
	//session.Values["token"]
	token := &Token{
		Token:   session.Values["token"].(string),
		App_url: app_url,
		Api_url: api_url,
		User_id: session.Values["id"].(int),
	}
	t := template.Must(template.ParseFiles("views/layouts/layout.html", "views/partials/sidebar.html", "views/layouts/categorys_add.html"))
	t.Execute(w, token)
}

func categoryEditHandler(w http.ResponseWriter, r *http.Request) {
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
	t := template.Must(template.ParseFiles("views/layouts/layout.html", "views/partials/sidebar.html", "views/layouts/categorys_edit.html"))
	t.Execute(w, token)
}
