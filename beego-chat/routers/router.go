package routers

import (
	"beego-chat/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/getroomid", &controllers.MainController{}, "Get:GetRoomId")
}
