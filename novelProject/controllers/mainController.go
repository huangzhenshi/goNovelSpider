package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"novelProject/models"
	"novelProject/service"
)

type MainController struct {
	beego.Controller
}

// @router / [get]
func (this *MainController) Index() {
	list,count:=models.GetAllChapter()
	fmt.Println(count)
	this.Data["list"]=list
	this.TplName = "index.html"
}

// @router /read [get]
func (this *MainController) ReadNovel() {
	id,err:=this.GetInt("id")
	chapter,err:=models.GetChapterById(id)
	if err!=nil{}
	this.Data["chapter"]=chapter
	this.TplName = "novel.html"
}

// @router /download [get]
func (this *MainController) Download() {
	err:= service.DownloadNovle()
	if err!=nil{}
	this.Ctx.WriteString("download finished")
}