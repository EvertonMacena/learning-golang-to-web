package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Post struct {
	Id int
	Title string
	Body string
}

func main()  {
	post := Post{Id: 1, Title: "First Post", Body: "ours content"}

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		t := template.Must(template.ParseFiles("templates/index.html"))
		if err := t.ExecuteTemplate(writer, "index.html", post); err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
		}
	})

	fmt.Println(http.ListenAndServe(":3000", nil))
}