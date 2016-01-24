package main

type Book struct {
	Id          int    `json:"id"`
	Ozon_Id     int    `json:"Ozon_Id"`
	Description string `json:"Description"`
	Title       string `json:"Title"`
	Year        int    `json:"Year"`
	Lang_id     int    `json:"Lang_id"`
	Language    string `json:"Language"`
	Izd_Id      int    `json:"Izd_Id"`
	Publisher   string `json:"Publisher"`
	Series      string `json:"Series"`
	Series_Id   int    `json:"Series_Id"`
}

type Books []Book
