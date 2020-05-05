package controllers

import (
	"go-web/models"
	"html/template"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	posts := models.ListPosts()
	t := template.Must(template.ParseFiles("templates/layout.html", "templates/list.html"))
	if err := t.ExecuteTemplate(w, "layout.html", posts); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func CreatePost (w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("templates/layout.html", "templates/create.html"))
	if err := t.ExecuteTemplate(w, "layout.html", nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func StorePost(w http.ResponseWriter, r *http.Request) {
	post := models.Post{}
	post.Title = r.FormValue("title")
	post.Body = r.FormValue("body")

	post.Save()

	http.Redirect(w, r, "/", 301)
}
