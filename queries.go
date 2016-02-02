package main

var (
	varGetBookInfoByID = `
	select 
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
	where b.id=`

	varGetBooksList = `
	select 
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
				b.series_id=s.id`

	varInsertBook = `INSERT books SET 
		ozon_id=?, description=?, Title=?, year=?, lang_id=?, izd_id=?, series_id=?`
)
