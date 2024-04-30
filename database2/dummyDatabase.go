package database

import (
	"errors"
	"fmt"
	"manav402/graphql/graph/model"
)

type bookData map[string]interface{}

var DummyData = bookData{
	"data": []model.Book{},
}


func (b *bookData) GetBooks()([]model.Book,error){
	if v,ok := DummyData["data"];ok{
		if data,ok:=v.([]model.Book); ok {
			return data,nil
		}else{
			return []model.Book{},errors.New("can't parse the data from database")
		}
	}else{
		return []model.Book{},errors.New("the database seems empty")
	}
}

func (b *bookData)GetBook(bookId int)(model.Book,error){
	if v,ok := DummyData["data"]; ok{
		if allData,ok := v.([]model.Book); ok{
			for _,v := range allData {
				if v.ID == fmt.Sprintf("%s",bookId) {
					return v,nil
				}
			}
			return model.Book{},errors.New("can't find passed id on database")
		}else{
			return model.Book{},errors.New("error parsing database from data")
		}
	}else{
		return model.Book{},errors.New("the database seems empty or has no data fields")
	}
}


func (b *bookData)CreateBook(input model.Book)(model.Book,error){
	if data,ok := DummyData["data"];ok{
		if allData,ok := data.([]model.Book);ok {
			allData = append(allData, input)
			DummyData["data"] = allData
			return input,nil
		}else{
			return model.Book{},errors.New("can't parse the data from database")
		}
	}else{
		return model.Book{},errors.New("the database seems empty")
	}
}
