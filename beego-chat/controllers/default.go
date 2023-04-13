package controllers

import (
	"time"

	beego "github.com/beego/beego/v2/server/web"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

// HTTPCommonHead 响应消息头
type HTTPCommonHead struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Time int64  `json:"time,omitempty"`
}

// HTTPResponse 响应消息
type HTTPResponse struct {
	Head HTTPCommonHead `json:"ret"`
	Body interface{}    `json:"body,omitempty"`
}

func GenerateRspHead(Code int, Msg string) HTTPCommonHead {
	return HTTPCommonHead{Code: Code, Msg: Msg, Time: time.Now().Unix()}
}

func (c *MainController) Output(rsp interface{}) {
	c.Data["json"] = rsp
	c.ServeJSON()
	return
}

func (c *MainController) SendResp(Ret interface{}) {
	rsp := HTTPResponse{}
	rsp.Head = GenerateRspHead(0, "ok")
	if Ret != nil {
		rsp.Body = Ret
	}
	c.Output(rsp)
}

func (c *MainController) SendRespMsg(Ret interface{}, msg string) {
	rsp := HTTPResponse{}
	rsp.Head = GenerateRspHead(0, msg)
	if Ret != nil {
		rsp.Body = Ret
	}
	c.Output(rsp)
}

func (c *MainController) SendRespErr(msg string) {
	rsp := HTTPResponse{}
	rsp.Head = GenerateRspHead(1, msg)
	c.Output(rsp)
}

func (c *MainController) SendRespErrCode(errcode int, msg string) {
	rsp := HTTPResponse{}
	rsp.Head = GenerateRspHead(errcode, msg)
	c.Output(rsp)
}
