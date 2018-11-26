package service

import (
	"fmt"
	"github.com/astaxie/beego"
	. "novelProject/spider"
)



func DownloadNovle() error{
	url:=beego.AppConfig.String("downloadUrl")
	s := new(BookTextSpider)
	err := s.SpiderUrl(url)
	if err!=nil {
		fmt.Println("err")
	}
	return err
}