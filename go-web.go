package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"go-web/controllers"
	"net/http"
)

func main()  {
	r := mux.NewRouter()
	r.PathPrefix("/static").Handler(http.StripPrefix("/static", http.FileServer(http.Dir("static/"))))

	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/post/create", controllers.CreatePost)
	r.HandleFunc("/post/store", controllers.StorePost)


	fmt.Println(http.ListenAndServe(":3000", r))
}
