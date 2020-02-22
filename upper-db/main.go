package main

import (
	"fmt"
	"log"

	"upper.io/db.v3"
	"upper.io/db.v3/sqlite"
)

type Book struct {
	ID        *uint  `db:"id"`
	Title     string `db:"title"`
	AuthorID  uint   `db:"author_id"`
	SubjectID uint   `db:"subject_id"`
}

// https://upper.io/db.v3
func main() {
	settings := sqlite.ConnectionURL{
		Database: "./testdb",
		Options:  nil,
	}

	sess, err := sqlite.Open(settings)
	if err != nil {
		log.Fatal(err)
	}
	defer sess.Close()

	sess.SetLogging(true)

	_, err = sess.Exec("create table if not exists books ( id int auto_increment primary key, title varchar(255), author_id int, subject_id int)")
	if err != nil {
		log.Fatal(err)
	}

	b := Book{
		ID:        nil,
		Title:     "My sample book",
		AuthorID:  2,
		SubjectID: 3,
	}

	id, err := sess.Collection("books").Insert(b)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted id", id)

	res := sess.Collection("books").Find("id", 1)
	cnt, err := res.Count()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Count", cnt)

	var b2 Book
	res = sess.Collection("books").Find(db.Cond{"title =": "My sample book"})
	err = res.One(&b2)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", b2)



	_, _ = sess.Collection("books").Insert(Book{
		Title:     "Javascript, the good parts",
		AuthorID:  1,
		SubjectID: 2,
	})

	_, _ = sess.Collection("books").Insert(Book{
		Title:     "Javascript, the bad parts",
		AuthorID:  1,
		SubjectID: 2,
	})

	_, _ = sess.Collection("books").Insert(Book{
		Title:     "Javascript, the so-so parts",
		AuthorID:  1,
		SubjectID: 2,
	})

	res = sess.Collection("books").Find(db.Cond{"title LIKE": "%javascript%"})
	var books []Book
	err = res.All(&books)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(books, len(books))
}
