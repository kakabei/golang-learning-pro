package controllers

import (
	"beego-chat/models"
	"fmt"

	"github.com/beego/beego/v2/core/config"
)

func (c *MainController) GetRoomId() {

	uid, _ := c.GetInt64("uid")
	name := c.GetString("name")

	req := models.GetRoomIdRequest{
		Uid:  uid,
		Name: name,
	}

	resp := models.GetRoomIdRespones{
		Uid:    1024,
		RoomId: 12345,
	}

	dbMap, _ := config.GetSection("mysql1")
	Port := dbMap["port"]
	fmt.Printf("Port %v \n", Port)
	fmt.Printf("req :%v, resp:%v \n", req, resp)

	c.SendResp(resp)
}
