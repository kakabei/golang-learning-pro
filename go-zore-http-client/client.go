package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpc"
)

type DataReq struct {
	Offset int64 `json:"offset"`
	Limit  int64 `json:"limit"`
}

type Request struct {
	Header string `json:"_" header:"X-System"`
	DataReq
}

// data from server A
type BaseA struct {
	Offset int64 `json:"offset"`
	Limit  int64 `json:"limit"`
}

// data from server B
type BaseB struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

// send to server C
type BaseC struct {
	BaseA
	BaseB
	Addr string `json:"addr"`
}

func main() {
	flag.Parse()

	req := new(Request)
	req.Header = "test-system"
	req.Offset = 1
	req.Limit = 100

	fmt.Printf("req -----> %+v\n", req)

	baseC := new(BaseC)
	baseC.Offset = 199
	baseC.Id = 1
	baseC.Name = "kane"
	baseC.Addr = "guangdong"
	fmt.Printf("baseC -----> %v\n", baseC)

	// struct to json
	baseCByte, _ := json.Marshal(baseC)
	fmt.Printf("baseCByte -----> %s\n", string(baseCByte))

	data, err := httpc.Do(context.Background(), http.MethodPost, "http://127.0.0.1:21321/v1/pingtest", req)
	if err != nil {
		fmt.Printf("httpc.Do %v...\n", err)
		return
	}

	fmt.Printf("httpc.Do ... data:%+v\n", data)

}
