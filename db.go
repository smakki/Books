package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var id string
var books Books

//func findBook(bookId int64) int64 {
//	bookId int64
//	db, err := sql.Open("mysql", dbuser+":"+dbpass+"@tcp(localhost:3306)/"+dbname+"?charset=utf8")
//	defer db.Close()
//	checkErr(err)
//	rows, err := db.Query(`SELECT id from books where id= ` + id)
//	checkErr(err)

//	for rows.Next() {
//		err = rows.Scan(&bookId)
//		checkErr(err)
//	}
//	return bookId
//}

func getBookInfoById(id string) Book {
	var b Book
	db, err := sql.Open("mysql", dbuser+":"+dbpass+"@tcp(localhost:3306)/"+dbname+"?charset=utf8")
	defer db.Close()
	checkErr(err)
	rows, err := db.Query(varGetBookInfoByID + id)
	checkErr(err)

	for rows.Next() {
		err = rows.Scan(&b.Id, &b.Ozon_Id, &b.Title, &b.Year, &b.Description,
			&b.Lang_id, &b.Izd_Id, &b.Series_Id, &b.Publisher, &b.Language, &b.Series)
		checkErr(err)
	}
	return b
}

func getBooksList() {
	var b Book
	db, err := sql.Open("mysql", dbuser+":"+dbpass+"@tcp(localhost:3306)/"+dbname+"?charset=utf8")
	defer db.Close()
	checkErr(err)
	rows, err := db.Query(varGetBooksList)
	checkErr(err)

	for rows.Next() {
		err = rows.Scan(&b.Id, &b.Ozon_Id, &b.Title, &b.Year, &b.Description, &b.Lang_id, &b.Izd_Id, &b.Series_Id, &b.Publisher, &b.Language, &b.Series)
		checkErr(err)
		books = append(books, b)
	}
}

func insertBook(b Book) int64 {
	db, err := sql.Open("mysql", dbuser+":"+dbpass+"@tcp(localhost:3306)/"+dbname+"?charset=utf8")
	defer db.Close()
	checkErr(err)
	query, err := db.Prepare(varInsertBook)
	checkErr(err)
	res, err := query.Exec(b.Ozon_Id, b.Description, b.Title, b.Year, b.Lang_id, b.Izd_Id, b.Series_Id)
	checkErr(err)
	id, err := res.LastInsertId()
	checkErr(err)
	return id
}
