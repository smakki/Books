package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var id string
var books Books

func getBookInfoById(id string) {
	var b Book
	db, err := sql.Open("mysql", dbuser+":"+dbpass+"@tcp(localhost:3306)/"+dbname+"?charset=utf8")
	defer db.Close()
	checkErr(err)
	rows, err := db.Query(`select 
	b.id,b.ozon_Id,b.Title,b.year, CAST(b.description AS CHAR(10000) CHARACTER SET utf8) AS Description,
	b.lang_id, b.izd_id, b.series_id,
	IFNULL(i.name, '') AS Publisher,IFNULL(l.name, '') AS LANGUAGE,IFNULL(s.name, '') AS Series
from 
	books b
		left join
			izd i
		on
			b.izd_id=i.id
			
		left join
			languages l
		on
			b.lang_id=l.id
			
		left join
			series s
		on
			b.series_id=s.id
where b.id= ` + id)
	checkErr(err)

	for rows.Next() {

		err = rows.Scan(&b.Id, &b.Ozon_Id, &b.Title, &b.Year, &b.Description,
			&b.Lang_id, b.Izd_Id, b.Series_Id, b.Publisher, b.Language, b.Series)
		checkErr(err)
		books = append(books, b)
	}
}

func getBooksList() {
	var b Book
	db, err := sql.Open("mysql", dbuser+":"+dbpass+"@tcp(localhost:3306)/"+dbname+"?charset=utf8")
	defer db.Close()
	checkErr(err)
	rows, err := db.Query(`select 
	b.id,b.ozon_Id,b.Title,b.year, CAST(b.description AS CHAR(10000) CHARACTER SET utf8) AS Description,
	b.lang_id, b.izd_id, b.series_id,
	IFNULL(i.name, '') AS Publisher,IFNULL(l.name,'') AS LANGUAGE,IFNULL(s.name,'') AS Series
from 
	books b
		left join
			izd i
		on
			b.izd_id=i.id
			
		left join
			languages l
		on
			b.lang_id=l.id
			
		left join
			series s
		on
			b.series_id=s.id`)
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
	query, err := db.Prepare(`INSERT books SET 
		ozon_id=?, description=?, Title=?, year=?, lang_id=?, izd_id=?, series_id=?`)
	checkErr(err)
	res, err := query.Exec(b.Ozon_Id, b.Description, b.Title, b.Year, b.Lang_id, b.Izd_Id, b.Series_Id)
	checkErr(err)
	id, err := res.LastInsertId()
	checkErr(err)
	return id
}
