package main

import (
	"html/template"
	"net/http"
)

func dashboardHandler(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("views/layouts/layout.html", "views/partials/sidebar.html", "views/layouts/dashboard.html"))
	t.Execute(w, nil)
}
