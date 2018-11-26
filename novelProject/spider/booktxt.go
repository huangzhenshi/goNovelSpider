package spider

import (
	"fmt"
	"log"
	"net/http"
	"novelProject/common"
	"novelProject/models"
	"time"
	"strings"
	"github.com/PuerkitoBio/goquery"
)

type BookTextSpider struct{
    
}



func (self *BookTextSpider)SpiderUrl(url string)( error){
	book := models.Book{}
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	if resp.StatusCode != 200 {
		fmt.Println("website connection error")
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil{
		return err
	}
	bookname := common.GbkToUtf8(doc.Find("#info h1").Text())
	fmt.Println("service's 名字是：",bookname)

	b, err := models.GetBookByName(bookname)
	if err != nil || b == nil{
		b := models.Book{Name:bookname, CreatedAt:time.Now(), UpdatedAt:time.Now()}
		models.BookAdd(&b)
	}

	doc.Find("#list dd").Each(func (i int, contentSelection *goquery.Selection){
		if i < 8{
			return
		}
		pre := i - 8
		next := i -6
		title := common.GbkToUtf8(contentSelection.Find("a").Text())
		href, _ := contentSelection.Find("a").Attr("href")
		chapter := models.Chapter{Title:title,Url:url+href, Sort:i - 7, Pre:pre, Next:next}
		book.Chapters = append(book.Chapters, &chapter)
	})

	//开启多线程 爬取各章节的小说内容
	channel := make(chan struct{}, 100)
	for _, chapter := range book.Chapters{
		channel <- struct{}{}
		go SpiderChapter(b.Id, chapter, channel)
	}

	for i := 0; i < 100; i++{
		channel <- struct{}{}
	}
	close(channel)
	fmt.Println("all complate")
	return nil
}

type ChanTag struct{}

//如果失败下载，会重试一次
func SpiderChapter(bookid int, chapter *models.Chapter, c chan struct{}){
	defer func(){<- c}()
	err:= downloadChapter(bookid,chapter,false)
	if err!=nil {
		downloadChapter(bookid,chapter,true)
	}
}

func downloadChapter(bookid int,chapter *models.Chapter,writeLog bool)  error{
	chapterUrl:= chapter.Url

	oldChapter,err := models.GetChapterByUrl(chapterUrl)
	//检查该chapter是否已经下载过了，避免重复下载
	if  oldChapter != nil{
		//fmt.Println("this charpter exists before:" ,chapter.Title)
		return nil
	}else{
		fmt.Println("this charpter require added:" ,chapter.Title)
	}

	//先试探，获取到内容之后，再整理，这样如果失败，也方便排查
	resp, err := http.Get(chapterUrl)
	if err != nil {
		if writeLog{
			undownload := models.Undownload{ErrorBookId:bookid,ChapterTitle:chapter.Title,Info:"downloading fail"}
			logId,errInsert:= models.ErrorLogAdd(&undownload)
			if errInsert!=nil{
				log.Println(errInsert)
				fmt.Println(logId)
			}
			return nil
		}else{
			log.Println(err)
			return err
		}
	}

	if resp.StatusCode != 200 {
		fmt.Println("website connection error and errcode is ",resp.StatusCode)
		return nil
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)

	if err != nil{
		fmt.Println(err)
		fmt.Println(chapter.Title,"downloading fail")
		return nil
	}
	content := doc.Find("#content").Text()
	content = common.GbkToUtf8(content)
	content = strings.Replace(content, "聽", " ", -1)
	ch := models.Chapter{BookId:bookid, Title:chapter.Title, Url:chapter.Url,Content:content,Sort:chapter.Sort, Pre:chapter.Pre, Next:chapter.Next, CreatedAt:time.Now(),UpdatedAt:time.Now()}

	//这里没有添加事务，如果添加成功之后会check是否有undownload记录
	models.ChapterAdd(&ch)
	models.CheckAndDelete(chapter.Title)
	fmt.Println("finish chapter",chapter.Title)
	return nil
}