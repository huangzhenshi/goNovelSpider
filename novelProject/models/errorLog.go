package models

import (
	"fmt"
)

type Undownload struct{
	Id int
	ErrorBookId int
	ChapterTitle string
	Info string
}

func ErrorLogAdd(err *Undownload) (int64,error){
	 return O.Insert(err)
}

func CheckAndDelete(title string){
	var undownload Undownload
	O.QueryTable("undownload").Filter("chapter_title", title).One(&undownload)
	if  undownload.Id >0{
		num,err:=O.QueryTable("undownload").Filter("Id", undownload.Id).Delete()
		if err!=nil{}
		fmt.Println(num)
	}
}