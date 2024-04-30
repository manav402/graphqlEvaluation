package controller

import (
	"errors"
	"io"
	"log"
	"manav402/graphql/database"
	"manav402/graphql/graph/model"
)

func GetBook(bookId int) (model.Book, error) {
	query, err := database.DB.Prepare(`SELECT * FROM books WHERE id = $1 LIMIT 1`)
	if err != nil {
		return model.Book{}, err
	}

	rows, err := query.Query(bookId)
	var output model.Book

	for rows.Next() {
		err := rows.Scan(&output.ID, &output.Name, &output.Price)
		if err != nil && err != io.EOF {
			log.Println("err parsin db data ", err)
			return model.Book{},err
		}
	}

	return output,nil
}

func CreateBook(input model.Book)(model.Book,error){
	query,err := database.DB.Prepare(`INSERT INTO books(id,name,price) VALUES ($1,$2,$3)`)
	if err != nil {
		return model.Book{}, err
	}

	rows, err := query.Exec(input.ID,input.Name,input.Price)
	affected,err := rows.RowsAffected()
	if err != nil {
		return model.Book{},err
	}
	if affected != 1 {
		return model.Book{},errors.New("there are some error inserting data in database")
	}
	return input,nil
}
