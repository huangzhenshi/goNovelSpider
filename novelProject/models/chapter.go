package models

import (
	"fmt"
	"strings"
	"time"
)


type Chapter struct{
	Id int
	BookId int
	Title string
	Content string
	Url string
	Sort int
	Pre int
	Next int
	CreatedAt time.Time
	UpdatedAt time.Time
}

func ChapterAdd(chapter *Chapter)(int64, error){
	return O.Insert(chapter)
}

func GetChapterById(id int)(*Chapter, error){
	chapter := new(Chapter)
	err := O.QueryTable("chapter").Filter("id", id).One(chapter)
	if err != nil{
		return nil, err
	}

	content:= chapter.Content
	chapter.Content=strings.Replace(content,"。","。<br>",-1)

	return chapter, nil
}

func GetChapterByUrl(url string)(*Chapter, error){
	var chapter Chapter
	err :=  O.QueryTable("chapter").Filter("url", url).One(&chapter)

	if err != nil{
		fmt.Println(" search error",url)
		return nil, err
	}
	return &chapter, nil
}

func GetAllChapter() ([]*Chapter, int64){
	query := O.QueryTable("chapter")
	total, _ := query.Count()
	list := make([]*Chapter, 0)
	query.OrderBy("sort").All(&list)

	fmt.Println("the size of chapter",len(list))
	return list, total
}