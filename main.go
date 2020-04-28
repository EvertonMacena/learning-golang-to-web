package main

import (
	"fmt"
	"html/template"
	"net/http"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type Post struct {
	Id int
	Title string
	Body string
}

var db, err = sql.Open("mysql", "root:root@tcp(192.168.99.100:3306)/go_db")

func main()  {
	//stmt, err := db.Prepare("Insert into posts(title, body) values (?, ?)")
	//checkErr(err)
	//
	//_, err = stmt.Exec("My first post", "My first content")
	//checkErr(err)
	//
	rows, err := db.Query("Select * from posts")
	checkErr(err)

	items := []Post{}

	for rows.Next() {
		post := Post{}

		rows.Scan(&post.Id, &post.Title, &post.Body)
		items = append(items, post)
	}

	db.Close()

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		post := Post{Id: 1, Title: "Unamed Post", Body: "No content"}

		if title := request.FormValue("title"); title != "" {
			post.Title = title
		}

		t := template.Must(template.ParseFiles("templates/index.html"))
		if err := t.ExecuteTemplate(writer, "index.html", items); err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
		}
	})

	fmt.Println(http.ListenAndServe(":3000", nil))
}

func checkErr(err error){
	if err != nil {
		panic(err.Error())  // Just for example purpose. You should use proper error handling instead of panic
	}
}
