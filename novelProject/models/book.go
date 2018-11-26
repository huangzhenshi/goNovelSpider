package models

import (
	"time"
)



type Book struct{
	Id int
	Name string
	Author string
	Image string
	CreatedAt time.Time
	UpdatedAt time.Time
	Chapters []*Chapter `orm:"-"`
}

func BookAdd(book *Book)(int64, error){
	return O.Insert(book)
}

func GetBookByName(name string)(*Book, error){
	book := new(Book)
	err := O.QueryTable("book").Filter("name", name).One(book)
	if err != nil || book.Id < 1{
		return nil, err
	}
	return book, nil
}