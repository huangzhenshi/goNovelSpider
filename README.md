# goNovelSpider


# 功能描述
1. 通过Goquery爬虫爬取https://www.booktxt.net/2_2219/ 这个网址的《圣墟》这篇连载的小说
2. 多线程（100个Goruntie）并发下载相应的章节
3. 把Book和Chapter分开存储在本地，实现离线、无广告观看
4. 支持/download的功能，手动同步下载该小说最新更新的章节
5. 支持断点下载（对篇章的URL做了校验，如果已经存在的话，不会重新下载）
6. 对异常做了日志处理（经常会有超时的下载)，下载失败会重试一次，仍然失败会记录在日志里面（如果重新下载成功的话，这个日志也会删除掉,日志中仅保留未下载成功的chapter)

# 框架
- Beego框架（路由、配置文件、ORM、日志）
- Goquery来爬取相应的内容
- Goruntie实现并发爬取（开了100个线程并发的去下载）

# 支持的功能
- 通过 localhost:8090/ 进入主页面，书名，已经成功下载的章节,点击已经下载的章节，会跳转到对应的内容页
- 通过 localhost:8090/read?id=5366，也可以直接查看相关的章节
- 支持 localhost:8090/download 重新下载未下载成功的章节

# 并发控制
1. 通过指定大小为100的channel来实现并发控制：
2. 每次开启一个线程的时候，往channel里面put一个空的对象，如果当前并发数已经达到100了，则主线程阻塞
3. 每次执行完一个下载任务的时候，pull出这个空的对象
4. 最后往channel里面连续put 100个空对象，从而确保所有的下载任务都已经结束了，避免下载未结束主线程就结束了，阻塞主线程等待下载结束
```
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

//如果失败下载，会重试一次
func SpiderChapter(bookid int, chapter *models.Chapter, c chan struct{}){
	defer func(){<- c}()
	err:= downloadChapter(bookid,chapter,false)
	if err!=nil {
		downloadChapter(bookid,chapter,true)
	}
}
```

# 解决下载失败的问题
因为网络问题无法保证每次下载都成功，而且经常因为超时而导致部分章节下载失败，所以
1. 引入重试，每次下载失败会重试一下
2. 重试仍然失败，则会记录在undownload表里面
3. 重新下载的时候，会跳过已经下载过了的章节并且重新下载未下载的章节，下载成功后会删除掉在undownload表里的记录 
```
	chapterUrl:= chapter.Url
	oldChapter,err := models.GetChapterByUrl(chapterUrl)
	//检查该chapter是否已经下载过了，避免重复下载
	if  oldChapter != nil{
		fmt.Println("this charpter exists before:" ,chapter.Title)
		return nil
	}else{
		fmt.Println("this charpter require added:" ,chapter.Title)
	}
	
	...

	//这里没有添加事务，如果添加成功之后会check是否有undownload记录
	models.ChapterAdd(&ch)
	models.CheckAndDelete(chapter.Title)


func CheckAndDelete(title string){
	var undownload Undownload
	O.QueryTable("undownload").Filter("chapter_title", title).One(&undownload)
	if  undownload.Id >0{
		O.QueryTable("undownload").Filter("Id", undownload.Id).Delete()
	}
}
```
