package models

type Post struct {
	Id string
	Title string
	Body string
}

func ListPosts() []Post {
	// open connection to the database
	db := ConnectDB()

	// make query database
	rows, err := db.Query("Select * from posts")
	checkErr(err)

	items := []Post{}

	for rows.Next() {
		post := Post{}
		rows.Scan(&post.Id, &post.Title, &post.Body)
		items = append(items, post)
	}
	// closed connection
	db.Close()

	return items
}

func (post *Post) Save (){
	// open connection to the database
	db := ConnectDB()

	// make query database
	stmt, err := db.Prepare("Insert into posts(title, body) values (?, ?)")
	_, err = stmt.Exec(post.Title, post.Body)
	checkErr(err)

	// closed connection
	db.Close()
}

func checkErr(err error){
	if err != nil {
		panic(err.Error())  // Just for example purpose. You should use proper error handling instead of panic
	}
}

