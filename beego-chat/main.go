package main

import (
	"beego-chat/models"
	_ "beego-chat/routers"

	"github.com/astaxie/beego/logs"
	beego "github.com/beego/beego/v2/server/web"
)

func main() {

	logs.SetLogger(logs.AdapterFile, `{"filename":"project.log","level":7,"maxlines":0,"maxsize":0,"daily":true,"maxdays":10,"color":true}`)
	// beego.BConfig.Listen.Graceful = true
	models.InitModel()
	beego.Run()
}
